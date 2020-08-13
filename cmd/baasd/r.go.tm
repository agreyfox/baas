package baasd

/*
ddds// #cgo CFLAGS:  -I /home/lq/block/baas/src/Fast-Fourier-Transform/include -I /usr/local/include/opencv4 -I /usr/local/include -g -Wall
asdfasfd// #cgo LDFLAGS:  -L /usr/local/lib -L /home/lq/libs -lwrapperwm
asdfsafd// #include <wrapper.h>
import "C"
*/
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"flag"

	"github.com/davidbyttow/govips/pkg/vips"
	"gopkg.in/h2non/bimg.v1"
)

var (
	TempDir = "/home/lq/data/temp"
)
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
}

/*
var flagInputImage = flag.String("in", "in.png", "Input image")
var flagOutputImage = flag.String("out", "out.jpg", "Output image")
var flagMkImage = flag.String("mk", "mk.png", "Watermask image")
*/
func CreateWaterMarkImageFile(text string, wm string) {
	rec := image.Rect(0, 0, 896, 262) //此数值为经验数值，
	background := color.RGBA{255, 255, 255, 0}

	myImage := image.NewRGBA(rec)
	var buff bytes.Buffer

	draw.Draw(myImage, rec, &image.Uniform{background}, image.ZP, draw.Src)

	jpeg.Encode(&buff, myImage, &jpeg.Options{})
	watermark := bimg.Watermark{
		Text:        text,
		Opacity:     0.8,
		Width:       240,
		DPI:         200,
		Margin:      0,
		NoReplicate: true,
		Font:        "sans bold 24",
		Background:  bimg.Color{0, 0, 0},
	}
	/*
		o := bimg.Options{
			Width:   512,
			Height:  512,
			Gravity: bimg.GravitySmart, // smart crop
			Type:    bimg.JPEG,
			Crop:    true,
			Rotate:  bimg.Angle(0),
			Flip:    false,
			Flop:    false,
			Quality: 100,
			Zoom:    1,
		}
	*/

	nimg := bimg.NewImage(buff.Bytes())
	m, err := nimg.Watermark(watermark)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	err = bimg.Write(wm, m)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DoWater(in, wm string) (string, error) {
	_, err := os.Stat(in)
	_, errr := os.Stat(wm)
	extension := filepath.Ext(in)
	storagedir, _ := filepath.Split(in)

	outtemp := storagedir + "/out." + RandString(4) + extension
	if err == nil && errr == nil {
		// path/to/whatever exists
		rr := C.makeWaterMark(C.CString(in), C.CString(wm), C.CString(outtemp))

		fmt.Printf("C function called with result %d\n", rr)
		//fmt.Println("mv  output")
		//destFile := TempDir + "/" + RandString(6) + extension
		//MoveFile(outtemp, destFile)
		return outtemp, nil
		//MoveFile(destFile, out)

		/* 	rrr := os.Rename(outtemp, out)
		if err != nil {
			fmt.Printf("Rename file error %s\n", rrr)
		}
		*/
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist

		fmt.Fprintln(os.Stderr, "File Not Found!")
		return "", err

	} else if os.IsNotExist(errr) {
		fmt.Fprintln(os.Stderr, "File not found!")
		return "", errr
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	return "", err
}

// out is extract watermark image
func ExtractWater(imgneedcheck, origin string) (string, error) {
	_, err := os.Stat(imgneedcheck)
	_, errr := os.Stat(origin)
	extension := filepath.Ext(origin)
	storagedir, _ := filepath.Split(origin)

	outtemp := storagedir + "/wm." + RandString(4) + extension
	if err == nil && errr == nil {
		// path/to/whatever exists
		ret := C.extractWaterMarker(C.CString(imgneedcheck), C.CString(origin), C.CString(outtemp))
		if ret == 0 {
			return outtemp, nil
		}
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist

		fmt.Fprintln(os.Stderr, "File Not Found!")
		return "", err

	} else if os.IsNotExist(errr) {
		fmt.Fprintln(os.Stderr, "File not found!")
		return "", errr
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	return "", err
}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

func CreateWaterMarkImageFile2(text string, wm string) error {

	rec := image.Rect(0, 0, 900, 300) //此数值为经验数值，
	background := color.RGBA{255, 255, 255, 3}

	myImage := image.NewRGBA(rec)
	var buff bytes.Buffer

	draw.Draw(myImage, rec, &image.Uniform{background}, image.ZP, draw.Src)

	jpeg.Encode(&buff, myImage, &jpeg.Options{})

	textimage, err := vips.Text(text, vips.InputInt("width", 64*40), vips.InputInt("height", 64))
	w := int(textimage.Xsize)
	h := int(textimage.Ysize)
	fmt.Printf("text image size is %dX%d\n", w, h)
	tRef := vips.NewImageRef(textimage, vips.ImageTypePNG)

	vips.NewTransform().LoadBuffer(buff.Bytes()).Resize(64*2, 32*2).Image(tRef).Invert().OutputFile(wm).Apply()
	return err
}

func Tryvips() {
	flag.Parse()

	vips.Startup(nil)
	defer vips.Shutdown()
	/* inFile, err := os.Open(*flagInputImage)

	if err != nil {
		log.Fatal(err)
	}

	defer inFile.Close()

	outFile, err := os.Create(*flagOutputImage)

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	processImage(inFile, outFile) */
	//ffttest(*flagInputImage, *flagMkImage, *flagOutputImage)
}

// Fftwtest try open input ,watermark to create out image with watermark
func Fftwtest(in, mk, out string) error {
	//Aphla := 35.0

	vips.Startup(nil)

	defer vips.Shutdown()
	fmt.Printf("Open Input File %s\n", in)
	abc, err := vips.NewImageFromFile(in)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//vips.Jpegsave(abc.Image(), "q.jpg", vips.InputInt("compress", 1))
	w := abc.Width()
	h := abc.Height()
	band := int(abc.Image().Bands)
	itype := abc.Image().Type

	//fmt.Println(abc)
	fmt.Printf("Original picture size %d*%d, type is %v band is %d,format %v,bits is %d,Coding is %v\n", w, h, itype, band, abc.Image().BandFmt, abc.Image().Bbits, abc.Image().Coding)
	mt, _ := abc.ToBytes()
	fmt.Println("Origin Buffer size is ", len(mt))
	fmt.Println("transfer to space input")
	space, err := vips.Fwfft(abc.Image())
	if err != nil {
		fmt.Println(err)
		return err
	}

	sw := int(space.Xsize)
	sh := int(space.Ysize)

	targeImgWrap := vips.NewImageRef(space, vips.ImageTypeJPEG)
	targetBuf, err := targeImgWrap.ToBytes() //abc.Export(vips.ExportParams{})

	//fmt.Println(targeImgWrap)
	//vips.Jpegsave(space, "ff1.jpg")
	fmt.Printf("input image  after fftw: bits:%v,width %d,height %d, band %d,format %v,type is %d,Coding is %v\nReady buf length is %d\n", space.Bbits, sw, sh, space.Bands, space.BandFmt, space.Type, space.Coding, len(targetBuf))

	fmt.Println("Read mask and resize file ")
	buf, rmkImgtype, err := vips.NewTransform().LoadFile(mk).Resize(w, h).Apply()
	rmkImg, err := vips.NewImageFromBuffer(buf)
	bw := rmkImg.Width()
	bh := rmkImg.Height()

	fmt.Printf("Transfor watermark file to half size of origin file,file type is %s, 长宽%dx%d\n", rmkImgtype, bw, bh)
	//encode(rmkImg, rmkSpaceImg, 35.0)
	/*
		tBuf, err := rmkImg.ToBytes()

		fmt.Printf("After Transform the buffer is %d ,band is %d\n", len(tBuf), rmkImg.Image().Bands)
		tmpbuf := make([]byte, w*h*3)
		wideLenArray := rand.Perm(bw)
		highLenArray := rand.Perm(bh)
		sizeOfPixel := 3
		alpha := byte(35)
		for ii := 0; ii < h/2; ii++ {
			for jj := 0; jj < w; jj++ {
				if wideLenArray[ii] < bh && highLenArray[jj] < bw {
					tmpbuf[(ii*w+jj)*sizeOfPixel] = tBuf[(wideLenArray[ii]*w+highLenArray[jj])*sizeOfPixel] * alpha
					tmpbuf[(ii*w+jj)*sizeOfPixel+1] = tBuf[(wideLenArray[ii]*w+highLenArray[jj])*sizeOfPixel+1] * alpha
					tmpbuf[(ii*w+jj)*sizeOfPixel+2] = tBuf[(wideLenArray[ii]*w+highLenArray[jj])*sizeOfPixel+2] * alpha

					tmpbuf[((h-1-ii)*w+(w-1-jj))*sizeOfPixel] = tBuf[(ii*w+jj)*sizeOfPixel] * alpha
					tmpbuf[((h-1-ii)*w+(w-1-jj))*sizeOfPixel+1] = tBuf[(ii*w+jj)*sizeOfPixel+1] * alpha
					tmpbuf[((h-1-ii)*w+(w-1-jj))*sizeOfPixel+2] = tBuf[(ii*w+jj)*sizeOfPixel+2] * alpha
				}
			}
		} */

	//tmpImg, err := vips.NewImageFromBuffer(buf)
	//result, err := vips.Add(targeImgWrap.Image(), tmpImg.Image())
	result := targeImgWrap.Image()
	outImage, err := vips.Invfft(result) //, vips.InputBool("real", true))
	if err != nil {
		fmt.Println(err)
		return err
	}
	vips.Jpegsave(outImage, out)
	out1, err := vips.Invfft(space)
	vips.Jpegsave(out1, "1-"+out)
	//fmt.Print(len(wideLenArray), len(highLenArray), len(tempbuf))
	//bxc, _ := abc.ToBytes()
	//fmt.Printf("\rWaterMark Memory format\n%3x\n", tBuf)

	//vips.NewTransform().Image(rmkImg).OutputFile("tt.jpg").Apply()  // can output

	/* 	rmkSpaceImg := rmkImg

	   	fmt.Println("变换加法", rmkSpaceImg)
	   	fmt.Println("transfer mark to fft space mark input")
	   	mkSpace, err := vips.Fwfft(rmkImg.Image())
	   	fmt.Println("FFt  to space input")
	   	outSpace, err := vips.Freqmult(space, mkSpace)
	   	fmt.Println("FFt  space  to groud")
	   	outImage, err := vips.Invfft(outSpace)
	   	if err != nil {
	   		fmt.Println(err)
	   		return err
	   	}
	   	fmt.Println("spectrum to out")
	   	oo, err := vips.Spectrum(outImage)
	   	vips.OutputImage(out, &oo)
	   	fmt.Println("save to out")
	   	vips.Jpegsave(oo, out, &vips.Option{Name: "mark"}) */
	return nil
}

// 将水印内容相加
func encode(origin *vips.ImageRef, mask *vips.ImageRef, factor float32) bool {
	width := origin.Width()
	height := origin.Height()
	buf, err := origin.ToBytes()
	if err != nil {
		fmt.Println("Get Image buffer error ", err)
		return false
	}
	fmt.Printf("the image buffer length is %d \n", len(buf))
	return true
	for k := 0; k < 3; k++ {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {

				//result_real[k][j][i] += factor * img.at<Vec3f>(i, j)[k];
			}
		}
	}
	return false
}

/*
func maintt() {
	rec := image.Rect(0, 0, 512, 512)
	background := color.RGBA{255, 255, 255, 0}

	myImage := image.NewRGBA(rec)
	var buff bytes.Buffer
	draw.Draw(myImage, rec, &image.Uniform{background}, image.ZP, draw.Src)

		o := bimg.Options{
			Width:      512,
			Height:     512,
			Gravity:    bimg.GravitySmart, // smart crop
			Type:       bimg.JPEG,
			Crop:       true,
			Rotate:     bimg.Angle(90),
			Flip:       false,
			Flop:       false,
			Quality:    100,
			Zoom:       0,
			Background: bimg.Color{244, 244, 244},
		}

	nimg := bimg.NewImage(buff.Bytes())
	 	buffer, err := bimg.Read("image.jpg")
	   	if err != nil {
	   		fmt.Fprintln(os.Stderr, err)
	   	}

	watermark := bimg.Watermark{
		Text:        "中文sdsd sdfsafafsdsafsaf s",
		Opacity:     0.5,
		Width:       200,
		DPI:         150,
		Margin:      0,
		NoReplicate: true,
		Font:        "sans bold 12",
		Background:  bimg.Color{0, 0, 0},
	}

	newImage, err := nimg.Watermark(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("new.jpg", newImage)
}
*/

/*
func processImage(inFile, outFile *os.File) {
	inImage, _, err := image.Decode(inFile)

	if err != nil {
		log.Fatal(err)
	}

	if _, ok := inImage.(*image.Gray); !ok {
		log.Fatal("Only works on gray scale images")
	}

	bounds := inImage.Bounds()

	realPixels := make([][]float64, bounds.Dy())

	for y := 0; y < bounds.Dy(); y++ {
		realPixels[y] = make([]float64, bounds.Dx())
		for x := 0; x < bounds.Dx(); x++ {
			r, _, _, _ := inImage.At(x, y).RGBA()
			realPixels[y][x] = float64(r)
		}
	}

	coeffs := fft.FFT2Real(realPixels)

	outImage := image.NewGray(bounds)

	coeffs = fft.IFFT2(coeffs)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			px := uint8(cmplx.Abs(coeffs[y][x]))
			outImage.SetGray(x, y, color.Gray{px})
		}
	}

	err = png.Encode(outFile, outImage)

	if err != nil {
		log.Fatal(err)
	}
} */

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}
func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func ByteToInt32(bytes []byte) int32 {
	//	bits := binary.LittleEndian.Uint32(bytes)
	bits := binary.BigEndian.Uint32(bytes)
	return int32(bits)
}

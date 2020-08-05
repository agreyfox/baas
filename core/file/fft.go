package file

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/agreyfox/baas/cmd/baas"
	"github.com/agreyfox/baas"
)

func (s *TransformService) MakeWaterInImage(ctx context.Context, ff *baas.File, blob *baas.FileBlob, text string) (*baas.FileBlob, error) {
	file, err := ioutil.TempFile("/tmp", "wm.*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	name := file.Name()
	file.Close()
	ok := s.ConvertTxtToImage(text, name)
	if !ok {
		s.Log.Error().Err(err).Msg("failed to create watermark image file")
		return &baas.FileBlob{}, errors.New("failed to create watermark image file")
	}
	_, err = os.Stat(name)
	//fd, err := os.Open(name)
	//fi, err := fd.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		s.Log.Error().Err(err).Msg("Error in generate watermark file ")
		return &baas.FileBlob{}, errors.New("Error in generate watermark file")
	}

	sourcefile := blob.TempFileName + "." + ff.Extension
	err = os.Rename(blob.TempFileName, sourcefile)
	if err != nil {
		s.Log.Error().Err(err).Msg("Error in generate source  file ")
		return &baas.FileBlob{}, errors.New("Error in generate source file")
	}
	blob.TempFileName = sourcefile // make sure the temporary file be deleted
	//	resultfile := blob.TempFileName + ".wm." + ff.Extension
	//resultfile := "1.png"
	/*resultfile, err := ioutil.TempFile("/tmp", "result.*.jpg")
	if err != nil {
		s.Log.Error().Err(err).Msg("Error in generate result file ")
	}
	resultname := resultfile.Name() */

	s.Log.Info().Msg("Now call lib with:" + sourcefile + " " + name)
	resultfile, err := baas..DoWater(sourcefile, name)
	if err == nil {
		fmt.Printf("result filename is %s\n", resultfile)
		fddd, _ := os.Open(resultfile)
		fii, err := fddd.Stat()
		fmt.Println("access temp watermark image file with error:", err)
		fmt.Println("watermark output file size is ", fii.Size())
		f := &baas.FileBlob{
			ID:           "watermark",
			Data:         ioutil.NopCloser(fddd),
			MIMEValue:    ff.MIMEValue,
			Size:         fii.Size(),
			TempFileName: resultfile,
			Hash:         "",
		}
		//fmt.Println(f)
		s.Log.Info().Msg("Water compute!")
		return f, nil
	} else {
		fmt.Printf("Open result filename  has error: %s\n", err)
		return &baas.FileBlob{}, errors.New("internal error")
	}

}

func (s *TransformService) VerifyWaterInImage(ctx context.Context, f *baas.File, blob *baas.FileBlob, origf *baas.File, origblob *baas.FileBlob) (*baas.FileBlob, error) {
	origsourcefile := origblob.TempFileName + "." + origf.Extension
	err := os.Rename(origblob.TempFileName, origsourcefile)
	if err != nil {
		s.Log.Error().Err(err).Msg("Error in generate source  file ")
		return &baas.FileBlob{}, errors.New("Error in generate source file")
	}
	origblob.TempFileName = origsourcefile // make sure the temporary file be deleted
	sourcefile := blob.TempFileName + "." + f.Extension
	err = os.Rename(blob.TempFileName, sourcefile)
	if err != nil {
		s.Log.Error().Err(err).Msg("Error in generate source  file ")
		return &baas.FileBlob{}, errors.New("Error in generate source file")
	}
	blob.TempFileName = sourcefile // make sure the temporary file be deleted

	s.Log.Info().Msg("Now call lib with:" + sourcefile + "," + origsourcefile)
	resultfile, err := baas..ExtractWater(sourcefile, origsourcefile)
	if err == nil {
		fmt.Printf("result filename is %s\n", resultfile)
		fddd, _ := os.Open(resultfile)
		fii, err := fddd.Stat()
		fmt.Println("access extract temp  watermark image file with error:", err)
		fmt.Println("watermark file size is ", fii.Size())
		f := &baas.FileBlob{
			ID:           "watermark",
			Data:         ioutil.NopCloser(fddd),
			MIMEValue:    origf.MIMEValue,
			Size:         fii.Size(),
			TempFileName: resultfile,
			Hash:         "",
		}

		s.Log.Info().Msg("Water extract  compute!")
		return f, nil
	}
	return &baas.FileBlob{}, nil
}

func (s *TransformService) ConvertTxtToImage(txt string, outputfile string) bool {
	ret := baas..CreateWaterMarkImageFile2(txt, outputfile)
	s.Log.Info().Msgf("result is %d", ret)
	return true
}

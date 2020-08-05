package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	_ "net/http/pprof"
	"os"

	"github.com/agreyfox/baas"
	"github.com/agreyfox/baas/adapter/bolt"
	"github.com/fatih/color"

	"github.com/agreyfox/baas/cmd/baasd"

	"github.com/agreyfox/baas/core/usage"

	"github.com/agreyfox/baas/core/file"

	"github.com/agreyfox/baas/core/upload"

	"github.com/agreyfox/baas/transport/http"

	"github.com/agreyfox/baas/core/application"

	"github.com/rs/cors"
	"github.com/rs/zerolog"

	"github.com/agreyfox/baas/adapter/aes"

	"github.com/gorilla/schema"

	"github.com/BurntSushi/toml"
)

var (
	tomlConfigPathPtr = flag.String("config", "./baas.toml", "the file path for the toml configuration file")
	inputImage        = flag.String("in", "test.jpg", "输入图片")
	outImage          = flag.String("out", "out.jpg", "输出图片")
	mkImage           = flag.String("wm", "wm.jpg", "水印图片")
	origImage         = flag.String("orig", "orig.jpg", "水印图片")
	textInput         = flag.String("text", "watermark", "输入文字")
	createwm          = flag.Bool("c", false, "创建watermask图像文件")
	convertwm         = flag.Bool("w", false, "手工转换带有watermask图片")
	extractwm         = flag.Bool("e", false, "获取wm图像")
)

func main() {

	flag.Parse()

	if err := run(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

func run() error {
	if *createwm {
		fmt.Println("创建水印...")
		baasd.CreateWaterMarkImageFile(*textInput, *mkImage)
		baasd.CreateWaterMarkImageFile2(*textInput, "wm.jpg")
		return nil
	}
	if *convertwm {
		fmt.Println("现在给图片打水印")
		//baas..Fftwtest(*inputImage, *mkImage, *outImage)
		name, err := baasd.DoWater(*inputImage, *mkImage)
		fmt.Printf("生成文件:%s,error:%s\n", name, err)
		return nil
	}

	if *extractwm {
		fmt.Println("现在尝试获取图片隐藏水印")
		//baas..Fftwtest(*inputImage, *mkImage, *outImage)
		name, err := baasd.ExtractWater(*inputImage, *origImage)
		fmt.Printf("生成文件:%s,error:%v\n", name, err)
		return nil
	}

	color.Blue("Start baas...")
	conf, err := parseConfig(*tomlConfigPathPtr)
	if err != nil {
		return fmt.Errorf("failed parsing config file %w", err)
	}

	if err := conf.InitConfig(); err != nil {
		return fmt.Errorf("failed initializing config %w", err)
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "baas").Logger()

	if err := os.MkdirAll(conf.Upload.ChunkUploadDir, 02750); err != nil {
		return err
	}

	if err := os.MkdirAll(conf.Upload.FullFileDir, 02750); err != nil {
		return err
	}

	var fileStorage baas.FileStorage
	var applicationStorage baas.ApplicationStorage
	var usageStorage baas.UsageStorage
	var transformStorage baas.TransformStorage

	/* if conf.DbEngine == DBEnginePostgreSQL {
		pgConf := conf.PostgreSQL
		pgxConf, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", pgConf.User, pgConf.Password, pgConf.Host, pgConf.Port, pgConf.Database))
		if err != nil {
			return fmt.Errorf("failed to parse db url %w", err)
		}

		db, err := pgxpool.ConnectConfig(context.Background(), pgxConf)
		if err != nil {
			return err
		}
		defer db.Close()

		fileStorage = &postgre.FileStorage{
			DB: db,
		}

		applicationStorage = &postgre.ApplicationStorage{
			DB: db,
		}

		usageStorage = &postgre.UsageStorage{
			DB: db,
		}

		transformStorage = &postgre.TransformStorage{
			DB: db,
		}
	} else { */
	// bolt db
	db, err := bolt.Open(conf.Bolt.Dir)
	if err != nil {
		return err
	}
	defer db.Close()

	fileStorage = &bolt.FileStorage{
		DB: db,
	}

	applicationStorage = &bolt.ApplicationStorage{
		DB: db,
	}

	usageStorage = &bolt.UsageStorage{
		DB: db,
	}

	transformStorage = &bolt.TransformStorage{
		DB: db,
	}
	//

	// schema decoder
	schemaDecoder := schema.NewDecoder()

	//////////////////////////////
	// SERVICES //
	/////////////////////////////
	encryptionService := &aes.EncryptionService{
		Key: []byte(conf.Security.AESKey),
	}

	applicationService := &application.Service{
		EncryptionService:  encryptionService,
		ApplicationStorage: applicationStorage,
	}

	fileService := &file.Service{
		FileStorage: fileStorage,
	}

	usageService := &usage.Service{
		UsageStorage: usageStorage,
	}

	transformService := &file.TransformService{
		UsageService:         usageService,
		TransformStorage:     transformStorage,
		Log:                  logger,
		MaxTransformFileSize: conf.Upload.MaxTransformFileSize,
	}

	fileServeService := &file.ServeService{
		FileStorage:  fileStorage,
		UsageService: usageService,

		ApplicationService: applicationService,
		TransformService:   transformService,

		FullFileDir: conf.Upload.FullFileDir,

		Log: logger,
	}

	uploadService := &upload.Service{
		ApplicationService: applicationService,
		FileService:        fileService,
		UsageService:       usageService,

		MaxChunkSize:      conf.Upload.MaxChunkSize,
		MaxUploadFileSize: conf.Upload.MaxUploadFileSize,
		ChunkUploadDir:    conf.Upload.ChunkUploadDir,
		FullFileDir:       conf.Upload.FullFileDir,

		Log: logger,
	}

	//////////////////////////////
	// HTTP //
	/////////////////////////////

	h := http.NewServer(&http.ServerConfig{
		ApplicationService: applicationService,
		UploadService:      uploadService,
		FileServeService:   fileServeService,
		UsageService:       usageService,

		QueryDecoder: schemaDecoder,

		AuthToken: conf.Security.AuthToken,

		Log: logger,
	})

	opts := cors.Options{
		AllowedMethods: []string{"PUT", "POST", "GET", "DELETE", "PATCH"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Referer"},
		AllowedOrigins: []string{"*"},
	}
	corsOpts := cors.New(opts)
	corsHandler := corsOpts.Handler(h)
	requestIDMiddleware := http.RequestIDMiddleware(corsHandler)

	return http.Serve(requestIDMiddleware, conf.HTTP.Port, conf.HTTP.HTTPS, conf.HTTP.SSLCertPath, conf.HTTP.SSLKeyPath)
}

func parseConfig(path string) (*baasd.Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	if _, err := io.Copy(buf, configFile); err != nil {
		return nil, err
	}

	//var conf baas..Config
	conf := baasd.GetSystemConfig()
	if _, err := toml.Decode(buf.String(), conf); err != nil {
		return nil, err
	}
	conf.Init = true // indicated the config is initialized
	return conf, nil
}

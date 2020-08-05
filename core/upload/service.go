package upload

import (
	"github.com/rs/zerolog"
)

type Service struct {
	FileService        baas.FileService
	UsageService       baas.UsageService
	ApplicationService baas.ApplicationService

	ChunkUploadDir    string
	FullFileDir       string
	MaxChunkSize      int64
	MaxUploadFileSize int64

	Log zerolog.Logger
}

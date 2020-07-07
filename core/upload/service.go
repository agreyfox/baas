package upload

import (
	"github.com/agreyfox/gisvs"
	"github.com/rs/zerolog"
)

type Service struct {
	FileService        gisvs.FileService
	UsageService       gisvs.UsageService
	ApplicationService gisvs.ApplicationService

	ChunkUploadDir    string
	FullFileDir       string
	MaxChunkSize      int64
	MaxUploadFileSize int64

	Log zerolog.Logger
}

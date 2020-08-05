package mock

import (
	"context"
	"mime/multipart"

	"github.com/agreyfox/baas"
)

type UploadService struct {
}

func (s *UploadService) Upload(ctx context.Context, r *multipart.Reader) (*baas.File, error) {
	return nil, nil
}

func (s *UploadService) ChunkUpload(ctx context.Context, r *multipart.Reader) error {
	return nil
}

func (s *UploadService) CompleteChunkUpload(ctx context.Context, appID, uploadID, filename string) (*baas.File, error) {
	return nil, nil
}

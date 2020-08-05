package mock

import (
	"context"
)

type FileBlobStorage struct {
}

func (s *FileBlobStorage) Upload(ctx context.Context, bucket string, blob *baas.FileBlob) error {
	return nil
}

func (s *FileBlobStorage) CreateBucket(ctx context.Context, bucket string) error {
	return nil
}

func (s *FileBlobStorage) FileBlob(ctx context.Context, bucket, id, tempDir string) (*baas.FileBlob, error) {
	return &baas.FileBlob{}, nil
}

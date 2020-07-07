package mock

import (
	"context"

	"github.com/agreyfox/gisvs"
)

type FileBlobStorage struct {
}

func (s *FileBlobStorage) Upload(ctx context.Context, bucket string, blob *gisvs.FileBlob) error {
	return nil
}

func (s *FileBlobStorage) CreateBucket(ctx context.Context, bucket string) error {
	return nil
}

func (s *FileBlobStorage) FileBlob(ctx context.Context, bucket, id, tempDir string) (*gisvs.FileBlob, error) {
	return &gisvs.FileBlob{}, nil
}

package mock

import (
	"context"

	"github.com/agreyfox/gisvs"
)

type FileService struct {
}

func (s *FileService) Create(ctx context.Context, n *gisvs.NewFile) (*gisvs.File, error) {
	return &gisvs.File{}, nil
}

func (s *FileService) File(ctx context.Context, id string) (*gisvs.File, error) {
	return &gisvs.File{}, nil
}

func (s *FileService) ApplicationFiles(ctx context.Context, appID, sinceID string, limit int) ([]*gisvs.File, error) {
	return []*gisvs.File{}, nil
}

func (s *FileService) Delete(ctx context.Context, id string) error {
	return nil
}

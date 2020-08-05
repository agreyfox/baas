package mock

import (
	"context"
)

type FileService struct {
}

func (s *FileService) Create(ctx context.Context, n *baas.NewFile) (*baas.File, error) {
	return &baas.File{}, nil
}

func (s *FileService) File(ctx context.Context, id string) (*baas.File, error) {
	return &baas.File{}, nil
}

func (s *FileService) ApplicationFiles(ctx context.Context, appID, sinceID string, limit int) ([]*baas.File, error) {
	return []*baas.File{}, nil
}

func (s *FileService) Delete(ctx context.Context, id string) error {
	return nil
}

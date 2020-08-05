package file

import (
	"context"
)

// Service is the implementation of the file service interface.
// It will handle the apps business logic concerning files.
type Service struct {
	FileStorage baas.FileStorage
}

func (s *Service) Create(ctx context.Context, n *baas.NewFile) (*baas.File, error) {
	return s.FileStorage.Store(ctx, n)
}

func (s *Service) File(ctx context.Context, id string) (*baas.File, error) {
	return s.FileStorage.File(ctx, id)
}

func (s *Service) ApplicationFiles(ctx context.Context, applicationID, sinceID string, limit int) ([]*baas.File, error) {
	if limit == 0 {
		limit = baas.DefaultFilePaginationLimit
	}

	return s.FileStorage.ApplicationFiles(ctx, applicationID, sinceID, limit)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.FileStorage.Delete(ctx, id)
}

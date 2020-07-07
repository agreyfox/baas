package file

import (
	"context"

	"github.com/agreyfox/gisvs"
)

// Service is the implementation of the file service interface.
// It will handle the apps business logic concerning files.
type Service struct {
	FileStorage gisvs.FileStorage
}

func (s *Service) Create(ctx context.Context, n *gisvs.NewFile) (*gisvs.File, error) {
	return s.FileStorage.Store(ctx, n)
}

func (s *Service) File(ctx context.Context, id string) (*gisvs.File, error) {
	return s.FileStorage.File(ctx, id)
}

func (s *Service) ApplicationFiles(ctx context.Context, applicationID, sinceID string, limit int) ([]*gisvs.File, error) {
	if limit == 0 {
		limit = gisvs.DefaultFilePaginationLimit
	}

	return s.FileStorage.ApplicationFiles(ctx, applicationID, sinceID, limit)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.FileStorage.Delete(ctx, id)
}

package application

import (
	"context"
	"fmt"

	"github.com/agreyfox/baas"
	"github.com/agreyfox/baas/cmd/baasd"
)

type Service struct {
	ApplicationStorage baas.ApplicationStorage

	EncryptionService baas.EncryptionService
}

func (s *Service) Create(ctx context.Context, n *baas.NewApplication) (*baas.Application, error) {
	cipherStorageSecretKey, err := s.EncryptionService.EncryptToString([]byte(n.StorageSecretKey))
	if err != nil {
		return nil, fmt.Errorf("failed enrcypting secret key %w", err)
	}
	conf := baasd.GetSystemConfig()
	if !conf.Security.Application {
		return nil, fmt.Errorf("Not allowed!")
	}
	if n.StorageEngine != "bolt" {

		n.StorageSecretKey = cipherStorageSecretKey

		if n.StorageEndpoint == "" {
			n.StorageEndpoint = makeStorageEndpoint(n.StorageEngine, n.StorageRegion)
		}

		if n.StorageBucket == "" {
			n.StorageBucket = makeStorageBucketName(n.Name)
		}

		if err := s.createStorageBucket(ctx, n); err != nil {
			return nil, fmt.Errorf("failed creating bucket %w", err)
		}
	}
	return s.ApplicationStorage.Store(ctx, n)
}

func (s *Service) Application(ctx context.Context, id string) (*baas.Application, error) {
	return s.ApplicationStorage.Application(ctx, id)
}

func (s *Service) Applications(ctx context.Context, sinceID string, limit int) ([]*baas.Application, error) {
	if limit == 0 {
		limit = baas.DefaultFilePaginationLimit
	}

	return s.ApplicationStorage.Applications(ctx, sinceID, limit)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.ApplicationStorage.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, u *baas.UpdateApplication) (*baas.Application, error) {
	return s.ApplicationStorage.Update(ctx, u)
}

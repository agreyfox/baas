package mock

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ApplicationService struct {
}

func (s *ApplicationService) Create(ctx context.Context, n *baas.NewApplication) (*baas.Application, error) {
	return &baas.Application{
		ID:               uuid.NewV4().String(),
		Name:             n.Name,
		Description:      n.Description,
		StorageRegion:    n.StorageRegion,
		StorageEndpoint:  "mock endpoint",
		StorageEngine:    n.StorageEngine,
		StorageBucket:    "mock bucket",
		StorageAccessKey: n.StorageAccessKey,
		DeliveryURL:      n.DeliveryURL,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, nil
}

func (s *ApplicationService) Application(ctx context.Context, id string) (*baas.Application, error) {
	if id != TestID {
		return nil, baas.ErrApplicationNotFound
	}
	return &baas.Application{}, nil
}

func (s *ApplicationService) Applications(ctx context.Context, sinceID string, limit int) ([]*baas.Application, error) {
	return []*baas.Application{}, nil
}

func (s *ApplicationService) Delete(ctx context.Context, id string) error {
	if id != TestID {
		return baas.ErrApplicationNotFound
	}
	return nil
}

func (s *ApplicationService) Update(ctx context.Context, u *baas.UpdateApplication) (*baas.Application, error) {
	if u.ID != TestID {
		return nil, baas.ErrApplicationNotFound
	}
	return &baas.Application{}, nil
}

func (s *ApplicationService) FileBlobStorage(engine, accessKey, secretKey, region, endpoint string) (baas.FileBlobStorage, error) {
	return &FileBlobStorage{}, nil
}

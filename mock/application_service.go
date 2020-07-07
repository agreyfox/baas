package mock

import (
	"context"
	"time"

	"github.com/agreyfox/gisvs"
	uuid "github.com/satori/go.uuid"
)

type ApplicationService struct {
}

func (s *ApplicationService) Create(ctx context.Context, n *gisvs.NewApplication) (*gisvs.Application, error) {
	return &gisvs.Application{
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

func (s *ApplicationService) Application(ctx context.Context, id string) (*gisvs.Application, error) {
	if id != TestID {
		return nil, gisvs.ErrApplicationNotFound
	}
	return &gisvs.Application{}, nil
}

func (s *ApplicationService) Applications(ctx context.Context, sinceID string, limit int) ([]*gisvs.Application, error) {
	return []*gisvs.Application{}, nil
}

func (s *ApplicationService) Delete(ctx context.Context, id string) error {
	if id != TestID {
		return gisvs.ErrApplicationNotFound
	}
	return nil
}

func (s *ApplicationService) Update(ctx context.Context, u *gisvs.UpdateApplication) (*gisvs.Application, error) {
	if u.ID != TestID {
		return nil, gisvs.ErrApplicationNotFound
	}
	return &gisvs.Application{}, nil
}

func (s *ApplicationService) FileBlobStorage(engine, accessKey, secretKey, region, endpoint string) (gisvs.FileBlobStorage, error) {
	return &FileBlobStorage{}, nil
}

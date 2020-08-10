package block

import (
	"context"
	"time"

	"github.com/agreyfox/baas"
	"github.com/rs/zerolog"
)

type Service struct {
	BlockStorage       baas.BlockStorage
	ApplicationService baas.ApplicationService
	EncryptionService  baas.EncryptionService
	UsageService       baas.UsageService
	Log                zerolog.Logger
}

func (s *Service) Create(ctx context.Context, n *baas.NewBAASUser) (*baas.BAASUser, error) {
	return &baas.BAASUser{
		Name:             "",
		Description:      "",
		StorageAccessKey: "",
		StorageSecretKey: "",
		StorageBucket:    "",
		ApplicationID:    "",
		PrivateKey:       "",
		Password:         "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}, nil
}

func (s *Service) DeleteBAASUser(ctx context.Context, id string) error {
	return nil
}
func (s *Service) UpdateBAASUser(ctx context.Context, u *baas.UpdateBAASUser) (*baas.BAASUser, error) {
	return &baas.BAASUser{
		Name:             "",
		Description:      "",
		StorageAccessKey: "",
		StorageSecretKey: "",
		StorageBucket:    "",
		ApplicationID:    "",
		PrivateKey:       "",
		Password:         "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}, nil
}
func (s *Service) GetKey(ctx context.Context, id string) (string, error) {
	return "", nil
}
func (s *Service) RecoverKey(ctx context.Context, id string) (string, error)        { return "", nil }
func (s *Service) GetBalance(ctx context.Context, addr string) (string, error)      { return "", nil }
func (s *Service) SendToken(ctx context.Context, addr string) (string, error)       { return "", nil }
func (s *Service) WriteMsg(ctx context.Context, addr string) (string, error)        { return "", nil }
func (s *Service) ReadMsg(ctx context.Context, addr string) (string, error)         { return "", nil }
func (s *Service) GetTxByHash(ctx context.Context, addr string) (string, error)     { return "", nil }
func (s *Service) GetErc20Balance(ctx context.Context, addr string) (string, error) { return "", nil }

func (s *Service) SendErc20Token(ctx context.Context, addr string) (string, error)   { return "", nil }
func (s *Service) CreateErc20Token(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) GetErc721Balance(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) GetErc721TotalSupply(ctx context.Context, addr string) (string, error) {
	return "", nil
}
func (s *Service) CreateErc721Token(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) SendErc721Token(ctx context.Context, addr string) (string, error)   { return "", nil }
func (s *Service) GetErc721MetaData(ctx context.Context, addr string) (string, error) { return "", nil }
func (s *Service) GetTx(ctx context.Context, addr string) (*baas.BlockTx, error) {
	return &baas.BlockTx{
		Type:   "",
		Filter: "",
	}, nil
}

package bolt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/agreyfox/baas"
	"github.com/asdine/storm/v3/q"

	uuid "github.com/satori/go.uuid"

	"github.com/asdine/storm/v3"
)

type BlockStorage struct {
	DB *storm.DB
}

type blockUser struct {
	Pk            int    `storm:"id,increment"`
	ID            string `storm:"index"`
	ApplicationID string
	FileID        string
	PrivateKey    string
	UserName      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t blockUser) validate(db *storm.DB) error {
	if t.ApplicationID == "" {
		return errors.New("applicationID is required")
	}

	if t.FileID == "" {
		return errors.New("fileID is required")
	}

	var tran transformation
	if err := db.Select(q.Eq("FileID", t.FileID)).First(&tran); err != nil {
		if err == storm.ErrNotFound {
			return nil
		}
		return err
	}

	return nil
}

func (s *BlockStorage) Store(ctx context.Context, n *baas.NewBAASUser) (*baas.BAASUser, error) {
	t := blockUser{
		ID:            uuid.NewV4().String(),
		ApplicationID: n.ApplicationID,
		PrivateKey:    n.PrivateKey,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := t.validate(s.DB); err != nil {
		return nil, err
	}

	if err := s.DB.Save(&t); err != nil {
		if err == storm.ErrAlreadyExists {
			return nil, baas.ErrApplicationNameTaken
		}
		return nil, err
	}

	mahiTran := sanitizeblockUser(t)

	return &mahiTran, nil
}

func (s *BlockStorage) GetKey(ctx context.Context, id string) (string, error) {
	fmt.Printf("User try to get private key")
	return "mock-key"
}

func (s *BlockStorage) Update(ctx context.Context, u *baas.UpdateBAASUser) (*baas.BAASUser, error) {
	fmt.Printf("User try to get private key")
	return &baas.BAASUser{
		Name:             "",
		Description:      "",
		StorageAccessKey: "",
		StorageSecretKey: "",
		StorageBucket:    "",
		ApplicationID:    "",
		PrivateKey:       "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
}
func sanitizeblockUser(t blockUser) baas.BAASUser {
	return baas.BAASUser{
		Name:          t.UserName,
		ApplicationID: t.ApplicationID,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

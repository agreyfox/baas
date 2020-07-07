package bolt

import (
	"context"
	"errors"
	"time"

	"github.com/agreyfox/gisvs"
	"github.com/asdine/storm/v3/q"

	uuid "github.com/satori/go.uuid"

	"github.com/asdine/storm/v3"
)

type TransformStorage struct {
	DB *storm.DB
}

type transformation struct {
	Pk            int    `storm:"id,increment"`
	ID            string `storm:"index"`
	ApplicationID string
	FileID        string
	Actions       gisvs.TransformationOption
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t transformation) validate(db *storm.DB) error {
	if t.ApplicationID == "" {
		return errors.New("applicationID is required")
	}

	if t.FileID == "" {
		return errors.New("fileID is required")
	}

	if t.Actions == (gisvs.TransformationOption{}) {
		return errors.New("actions is required")
	}

	var tran transformation
	if err := db.Select(q.Eq("FileID", t.FileID), q.Eq("Actions", t.Actions)).First(&tran); err != nil {
		if err == storm.ErrNotFound {
			return nil
		}
		return err
	}

	if tran.ID != "" {
		return gisvs.ErrTransformationNotUnique
	}

	return nil
}

func (s *TransformStorage) Store(ctx context.Context, n *gisvs.NewTransformation) (*gisvs.Transformation, error) {
	t := transformation{
		ID:            uuid.NewV4().String(),
		ApplicationID: n.ApplicationID,
		FileID:        n.FileID,
		Actions:       n.Actions,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := t.validate(s.DB); err != nil {
		return nil, err
	}

	if err := s.DB.Save(&t); err != nil {
		if err == storm.ErrAlreadyExists {
			return nil, gisvs.ErrApplicationNameTaken
		}
		return nil, err
	}

	mahiTran := sanitizeTransformation(t)

	return &mahiTran, nil
}

func sanitizeTransformation(t transformation) gisvs.Transformation {
	return gisvs.Transformation{
		ID:            t.ID,
		ApplicationID: t.ApplicationID,
		FileID:        t.FileID,
		Actions:       t.Actions,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

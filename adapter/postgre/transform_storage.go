package postgre

import (
	"context"
	"errors"
	"time"

	"github.com/agreyfox/gisvs"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	uniqueFileTransformationConstraint = "mahi_transformations_file_id_actions_key"
)

type TransformStorage struct {
	DB *pgxpool.Pool
}

type transformation struct {
	ID            string
	ApplicationID string
	FileID        string
	Actions       gisvs.TransformationOption
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (s *TransformStorage) Store(ctx context.Context, n *gisvs.NewTransformation) (*gisvs.Transformation, error) {
	var t transformation

	query := `
		INSERT INTO mahi_transformations (application_id, file_id, actions)
		VALUES ($1, $2, $3)
		RETURNING id, application_id, file_id, actions, created_at, updated_at
 `

	if err := s.DB.QueryRow(
		ctx,
		query,
		NewNullString(n.ApplicationID),
		NewNullString(n.FileID),
		n.Actions,
	).Scan(
		&t.ID,
		&t.ApplicationID,
		&t.FileID,
		&t.Actions,
		&t.CreatedAt,
		&t.UpdatedAt,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == uniqueFileTransformationConstraint {
				return nil, gisvs.ErrTransformationNotUnique
			}
		}
		return nil, err
	}

	mahiTransformation := sanitizeTransformation(t)

	return &mahiTransformation, nil
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

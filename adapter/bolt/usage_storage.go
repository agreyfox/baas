package bolt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/agreyfox/gisvs"
	"github.com/asdine/storm/v3/q"

	"github.com/jinzhu/now"

	uuid "github.com/satori/go.uuid"

	"github.com/asdine/storm/v3"
)

type usage struct {
	Pk                    int    `storm:"id,increment"`
	ID                    string `storm:"index"`
	ApplicationID         string `storm:"index"`
	Transformations       int64
	UniqueTransformations int64
	Bandwidth             int64
	Storage               int64
	FileCount             int64
	StartDate             string
	EndDate               string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type UsageStorage struct {
	DB *storm.DB
}

func (s *UsageStorage) Store(ctx context.Context, n *gisvs.NewUsage) (*gisvs.Usage, error) {
	start := n.StartDate.Format(gisvs.DateLayout)
	end := n.EndDate.Format(gisvs.DateLayout)

	if start == end {
		return nil, errors.New("start and end times cannot be the same")
	}

	u := usage{
		ID:                    uuid.NewV4().String(),
		ApplicationID:         n.ApplicationID,
		Transformations:       n.Transformations,
		UniqueTransformations: n.UniqueTransformations,
		Bandwidth:             n.Bandwidth,
		Storage:               n.Storage,
		FileCount:             n.FileCount,
		StartDate:             start,
		EndDate:               end,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	if err := s.DB.Save(&u); err != nil {
		return nil, err
	}

	mahiUsage, err := sanitizeUsage(u)
	if err != nil {
		return nil, err
	}

	return &mahiUsage, nil
}

func (s *UsageStorage) Update(ctx context.Context, u *gisvs.UpdateUsage) (*gisvs.Usage, error) {
	if u.ApplicationID == "" {
		return nil, errors.New("application id is required")
	}

	start := now.BeginningOfDay()
	if u.StartDate != (time.Time{}) {
		start = now.New(u.StartDate).BeginningOfDay()
	}

	end := now.BeginningOfDay().Add(25 * time.Hour)
	if u.EndDate != (time.Time{}) {
		end = now.New(u.EndDate).EndOfDay()
	}

	if start.Format(gisvs.DateLayout) == end.Format(gisvs.DateLayout) {
		return nil, errors.New("start and end times cannot be the same")
	}

	usage, err := s.Usage(ctx, u.ApplicationID, start, end)
	if err != nil && err != gisvs.ErrUsageNotFound {
		return nil, fmt.Errorf("failed getting usage %w", err)
	}

	// first entry of the day
	if err != nil && err == gisvs.ErrUsageNotFound {
		latestUsage, err := s.lastApplicationUsage(ctx, u.ApplicationID)
		if err != nil && err != gisvs.ErrUsageNotFound {
			return nil, err
		}

		// these items get rolled over they don't reset per day
		storage := latestUsage.Storage + u.Storage
		fileCount := latestUsage.FileCount + u.FileCount

		newUsage := &gisvs.NewUsage{
			ApplicationID:         u.ApplicationID,
			Transformations:       u.Transformations,
			UniqueTransformations: u.UniqueTransformations,
			Bandwidth:             u.Bandwidth,
			Storage:               storage,
			FileCount:             fileCount,
			StartDate:             start,
			EndDate:               end,
		}

		return s.Store(ctx, newUsage)
	}

	updatedUsage := &gisvs.UpdateUsage{
		ApplicationID:         u.ApplicationID,
		Transformations:       usage.Transformations + u.Transformations,
		UniqueTransformations: usage.UniqueTransformations + u.UniqueTransformations,
		Bandwidth:             usage.Bandwidth + u.Bandwidth,
		Storage:               usage.Storage + u.Storage,
		FileCount:             usage.FileCount + u.FileCount,
		StartDate:             start,
		EndDate:               end,
	}

	return s.update(ctx, usage.ID, updatedUsage)
}

func (s *UsageStorage) Usage(ctx context.Context, applicationID string, startDate, endDate time.Time) (*gisvs.Usage, error) {
	start := startDate.Format(gisvs.DateLayout)
	end := endDate.Format(gisvs.DateLayout)

	var u usage
	if err := s.DB.Select(q.Eq("ApplicationID", applicationID), q.Gte("StartDate", start), q.Lte("EndDate", end)).First(&u); err != nil {
		if err == storm.ErrNotFound {
			return nil, gisvs.ErrUsageNotFound
		}
		return nil, err
	}

	mahiUsage, err := sanitizeUsage(u)
	if err != nil {
		return nil, err
	}

	return &mahiUsage, nil
}

func (s *UsageStorage) boltUsage(ctx context.Context, id string) (*usage, error) {
	var u usage
	if err := s.DB.Select(q.Eq("ID", id)).First(&u); err != nil {
		if err == storm.ErrNotFound {
			return nil, gisvs.ErrUsageNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (s *UsageStorage) ApplicationUsages(ctx context.Context, applicationID string, startDate, endDate time.Time) ([]*gisvs.Usage, error) {
	start := startDate.Format(gisvs.DateLayout)
	end := endDate.Format(gisvs.DateLayout)

	var usages []*usage
	if err := s.DB.Select(q.Eq("ApplicationID", applicationID), q.Gte("StartDate", start), q.Lte("EndDate", end)).OrderBy("StartDate").Find(&usages); err != nil {
		if err == storm.ErrNotFound {
			return []*gisvs.Usage{}, nil
		}
		return nil, err
	}

	var mahiUsages []*gisvs.Usage

	for _, u := range usages {
		mahiUsage, err := sanitizeUsage(*u)
		if err != nil {
			return nil, err
		}

		mahiUsages = append(mahiUsages, &mahiUsage)
	}

	return mahiUsages, nil
}

func (s *UsageStorage) Usages(ctx context.Context, startDate, endDate time.Time) ([]*gisvs.TotalUsage, error) {
	start := startDate.Format(gisvs.DateLayout)
	end := endDate.Format(gisvs.DateLayout)

	var usages []*usage
	if err := s.DB.Select(q.Gte("StartDate", start), q.Lte("EndDate", end)).OrderBy("StartDate").Find(&usages); err != nil {
		if err == storm.ErrNotFound {
			return []*gisvs.TotalUsage{}, nil
		}
		return nil, err
	}

	var mahiTotalUsages []*gisvs.TotalUsage

	for _, u := range usages {
		mahiUsage, err := sanitizeTotalUsage(*u)
		if err != nil {
			return nil, err
		}

		mahiTotalUsages = append(mahiTotalUsages, &mahiUsage)
	}

	return mahiTotalUsages, nil
}

func (s *UsageStorage) update(ctx context.Context, id string, updatedUsage *gisvs.UpdateUsage) (*gisvs.Usage, error) {
	u, err := s.boltUsage(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting bolt usage %w", err)
	}

	u.Transformations = updatedUsage.Transformations
	u.UniqueTransformations = updatedUsage.UniqueTransformations
	u.Bandwidth = updatedUsage.Bandwidth
	u.Storage = updatedUsage.Storage
	u.FileCount = updatedUsage.FileCount
	u.UpdatedAt = time.Now()

	if err := s.DB.Save(u); err != nil {
		return nil, err
	}

	mahiUsage, err := sanitizeUsage(*u)
	if err != nil {
		return nil, err
	}

	return &mahiUsage, nil
}

func (s *UsageStorage) lastApplicationUsage(ctx context.Context, applicationID string) (gisvs.Usage, error) {
	var u usage
	if err := s.DB.Select(q.Eq("ApplicationID", applicationID)).OrderBy("StartDate").Reverse().First(&u); err != nil {
		if err == storm.ErrNotFound {
			return gisvs.Usage{}, gisvs.ErrUsageNotFound
		}
		return gisvs.Usage{}, err
	}

	mahiUsage, err := sanitizeUsage(u)
	if err != nil {
		return gisvs.Usage{}, err
	}

	return mahiUsage, nil
}

func sanitizeUsage(u usage) (gisvs.Usage, error) {
	start, err := time.Parse(gisvs.DateLayout, u.StartDate)
	if err != nil {
		return gisvs.Usage{}, err
	}

	end, err := time.Parse(gisvs.DateLayout, u.EndDate)
	if err != nil {
		return gisvs.Usage{}, err
	}

	return gisvs.Usage{
		ID:                    u.ID,
		ApplicationID:         u.ApplicationID,
		Transformations:       u.Transformations,
		UniqueTransformations: u.UniqueTransformations,
		Bandwidth:             u.Bandwidth,
		Storage:               u.Storage,
		FileCount:             u.FileCount,
		StartDate:             start,
		EndDate:               end,
		CreatedAt:             u.CreatedAt,
		UpdatedAt:             u.UpdatedAt,
	}, nil
}

func sanitizeTotalUsage(u usage) (gisvs.TotalUsage, error) {
	start, err := time.Parse(gisvs.DateLayout, u.StartDate)
	if err != nil {
		return gisvs.TotalUsage{}, err
	}

	end, err := time.Parse(gisvs.DateLayout, u.EndDate)
	if err != nil {
		return gisvs.TotalUsage{}, err
	}

	return gisvs.TotalUsage{
		Transformations:       u.Transformations,
		UniqueTransformations: u.UniqueTransformations,
		Bandwidth:             u.Bandwidth,
		Storage:               u.Storage,
		FileCount:             u.FileCount,
		StartDate:             start,
		EndDate:               end,
	}, nil
}

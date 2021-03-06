package bolt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/agreyfox/baas"
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
	TxCount               int64
	UserCount             int64
	StartDate             string
	EndDate               string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type UsageStorage struct {
	DB *storm.DB
}

func (s *UsageStorage) Store(ctx context.Context, n *baas.NewUsage) (*baas.Usage, error) {
	start := n.StartDate.Format(baas.DateLayout)
	end := n.EndDate.Format(baas.DateLayout)

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
		TxCount:               n.TxCount,
		UserCount:             n.UserCount,
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

func (s *UsageStorage) Update(ctx context.Context, u *baas.UpdateUsage) (*baas.Usage, error) {
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

	if start.Format(baas.DateLayout) == end.Format(baas.DateLayout) {
		return nil, errors.New("start and end times cannot be the same")
	}

	usage, err := s.Usage(ctx, u.ApplicationID, start, end)
	if err != nil && err != baas.ErrUsageNotFound {
		return nil, fmt.Errorf("failed getting usage %w", err)
	}

	// first entry of the day
	if err != nil && err == baas.ErrUsageNotFound {
		latestUsage, err := s.lastApplicationUsage(ctx, u.ApplicationID)
		if err != nil && err != baas.ErrUsageNotFound {
			return nil, err
		}

		// these items get rolled over they don't reset per day
		storage := latestUsage.Storage + u.Storage
		fileCount := latestUsage.FileCount + u.FileCount
		tx := latestUsage.TxCount + u.TxCount
		user := latestUsage.UserCount + u.UserCount
		newUsage := &baas.NewUsage{
			ApplicationID:         u.ApplicationID,
			Transformations:       u.Transformations,
			UniqueTransformations: u.UniqueTransformations,
			Bandwidth:             u.Bandwidth,
			Storage:               storage,
			FileCount:             fileCount,
			TxCount:               tx,
			UserCount:             user,
			StartDate:             start,
			EndDate:               end,
		}

		return s.Store(ctx, newUsage)
	}

	updatedUsage := &baas.UpdateUsage{
		ApplicationID:         u.ApplicationID,
		Transformations:       usage.Transformations + u.Transformations,
		UniqueTransformations: usage.UniqueTransformations + u.UniqueTransformations,
		Bandwidth:             usage.Bandwidth + u.Bandwidth,
		Storage:               usage.Storage + u.Storage,
		FileCount:             usage.FileCount + u.FileCount,
		TxCount:               usage.TxCount + u.TxCount,
		UserCount:             usage.UserCount + u.UserCount,
		StartDate:             start,
		EndDate:               end,
	}

	return s.update(ctx, usage.ID, updatedUsage)
}

func (s *UsageStorage) Usage(ctx context.Context, applicationID string, startDate, endDate time.Time) (*baas.Usage, error) {
	start := startDate.Format(baas.DateLayout)
	end := endDate.Format(baas.DateLayout)

	var u usage
	if err := s.DB.Select(q.Eq("ApplicationID", applicationID), q.Gte("StartDate", start), q.Lte("EndDate", end)).First(&u); err != nil {
		if err == storm.ErrNotFound {
			return nil, baas.ErrUsageNotFound
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
			return nil, baas.ErrUsageNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (s *UsageStorage) ApplicationUsages(ctx context.Context, applicationID string, startDate, endDate time.Time) ([]*baas.Usage, error) {
	start := startDate.Format(baas.DateLayout)
	end := endDate.Format(baas.DateLayout)

	var usages []*usage
	if err := s.DB.Select(q.Eq("ApplicationID", applicationID), q.Gte("StartDate", start), q.Lte("EndDate", end)).OrderBy("StartDate").Find(&usages); err != nil {
		if err == storm.ErrNotFound {
			return []*baas.Usage{}, nil
		}
		return nil, err
	}

	var mahiUsages []*baas.Usage

	for _, u := range usages {
		mahiUsage, err := sanitizeUsage(*u)
		if err != nil {
			return nil, err
		}

		mahiUsages = append(mahiUsages, &mahiUsage)
	}

	return mahiUsages, nil
}

func (s *UsageStorage) Usages(ctx context.Context, startDate, endDate time.Time) ([]*baas.TotalUsage, error) {
	start := startDate.Format(baas.DateLayout)
	end := endDate.Format(baas.DateLayout)

	var usages []*usage
	if err := s.DB.Select(q.Gte("StartDate", start), q.Lte("EndDate", end)).OrderBy("StartDate").Find(&usages); err != nil {
		if err == storm.ErrNotFound {
			return []*baas.TotalUsage{}, nil
		}
		return nil, err
	}

	var mahiTotalUsages []*baas.TotalUsage

	for _, u := range usages {
		mahiUsage, err := sanitizeTotalUsage(*u)
		if err != nil {
			return nil, err
		}

		mahiTotalUsages = append(mahiTotalUsages, &mahiUsage)
	}

	return mahiTotalUsages, nil
}

func (s *UsageStorage) update(ctx context.Context, id string, updatedUsage *baas.UpdateUsage) (*baas.Usage, error) {
	u, err := s.boltUsage(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting bolt usage %w", err)
	}

	u.Transformations = updatedUsage.Transformations
	u.UniqueTransformations = updatedUsage.UniqueTransformations
	u.Bandwidth = updatedUsage.Bandwidth
	u.Storage = updatedUsage.Storage
	u.FileCount = updatedUsage.FileCount
	u.TxCount = updatedUsage.TxCount
	u.UserCount = updatedUsage.UserCount
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

func (s *UsageStorage) lastApplicationUsage(ctx context.Context, applicationID string) (baas.Usage, error) {
	var u usage
	if err := s.DB.Select(q.Eq("ApplicationID", applicationID)).OrderBy("StartDate").Reverse().First(&u); err != nil {
		if err == storm.ErrNotFound {
			return baas.Usage{}, baas.ErrUsageNotFound
		}
		return baas.Usage{}, err
	}

	mahiUsage, err := sanitizeUsage(u)
	if err != nil {
		return baas.Usage{}, err
	}

	return mahiUsage, nil
}

func sanitizeUsage(u usage) (baas.Usage, error) {
	start, err := time.Parse(baas.DateLayout, u.StartDate)
	if err != nil {
		return baas.Usage{}, err
	}

	end, err := time.Parse(baas.DateLayout, u.EndDate)
	if err != nil {
		return baas.Usage{}, err
	}

	return baas.Usage{
		ID:                    u.ID,
		ApplicationID:         u.ApplicationID,
		Transformations:       u.Transformations,
		UniqueTransformations: u.UniqueTransformations,
		Bandwidth:             u.Bandwidth,
		Storage:               u.Storage,
		FileCount:             u.FileCount,
		TxCount:               u.TxCount,
		UserCount:             u.UserCount,
		StartDate:             start,
		EndDate:               end,
		CreatedAt:             u.CreatedAt,
		UpdatedAt:             u.UpdatedAt,
	}, nil
}

func sanitizeTotalUsage(u usage) (baas.TotalUsage, error) {
	start, err := time.Parse(baas.DateLayout, u.StartDate)
	if err != nil {
		return baas.TotalUsage{}, err
	}

	end, err := time.Parse(baas.DateLayout, u.EndDate)
	if err != nil {
		return baas.TotalUsage{}, err
	}

	return baas.TotalUsage{
		Transformations:       u.Transformations,
		UniqueTransformations: u.UniqueTransformations,
		Bandwidth:             u.Bandwidth,
		Storage:               u.Storage,
		FileCount:             u.FileCount,
		UserCount:             u.UserCount,
		TxCount:               u.TxCount,
		StartDate:             start,
		EndDate:               end,
	}, nil
}

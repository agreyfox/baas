package usage

import (
	"context"
	"time"

	"github.com/jinzhu/now"
)

type Service struct {
	UsageStorage baas.UsageStorage
}

func (s *Service) Update(ctx context.Context, u *baas.UpdateUsage) error {
	if _, err := s.UsageStorage.Update(ctx, u); err != nil {
		return err
	}

	return nil
}

func (s *Service) Usages(ctx context.Context, startDate, endDate time.Time) ([]*baas.TotalUsage, error) {
	start := startDate
	if start == (time.Time{}) {
		start = now.BeginningOfMonth()
	}

	end := endDate
	if end == (time.Time{}) {
		end = now.EndOfMonth()
	}

	return s.UsageStorage.Usages(ctx, start, end)
}

func (s *Service) ApplicationUsages(ctx context.Context, applicationID string, startDate, endDate time.Time) ([]*baas.Usage, error) {
	start := startDate
	if start == (time.Time{}) {
		start = now.BeginningOfMonth()
	}

	end := endDate
	if end == (time.Time{}) {
		end = now.EndOfMonth()
	}

	return s.UsageStorage.ApplicationUsages(ctx, applicationID, start, end)
}

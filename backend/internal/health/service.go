package health

import (
	"context"
)

type CheckResult struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type Service struct {
	checkers []Checker
}

func NewService(checkers ...Checker) *Service {
	return &Service{
		checkers: checkers,
	}
}

func (s *Service) CheckReadiness(ctx context.Context) ([]CheckResult, bool) {
	results := make([]CheckResult, 0, len(s.checkers))
	allHealthy := true

	for _, checker := range s.checkers {
		err := checker.Check(ctx)

		if err != nil {
			allHealthy = false

			results = append(results, CheckResult{
				Name:   checker.Name(),
				Status: "error",
				Error:  err.Error(),
			})

			continue
		}

		results = append(results, CheckResult{
			Name:   checker.Name(),
			Status: "ok",
		})
	}

	return results, allHealthy
}

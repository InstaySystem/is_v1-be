package service

import (
	"context"

	"github.com/InstaySystem/is_v1-be/internal/types"
)

type DashboardService interface {
	Overview(ctx context.Context) (*types.DashboardResponse, error)
}

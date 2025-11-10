package service

import (
	"context"

	"github.com/InstaySystem/is-be/internal/types"
)

type DepartmentService interface {
	CreateDepartment(ctx context.Context, userID int64, req types.CreateDepartmentRequest) (int64, error)
}
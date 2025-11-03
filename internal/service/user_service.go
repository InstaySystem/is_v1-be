package service

import (
	"context"

	"github.com/InstaySystem/is-be/internal/types"
)

type UserService interface {
	CreateUser(ctx context.Context, req types.CreateUserRequest) (int64, error)
}
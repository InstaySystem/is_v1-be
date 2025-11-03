package repository

import (
	"context"

	"github.com/InstaySystem/is-be/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}
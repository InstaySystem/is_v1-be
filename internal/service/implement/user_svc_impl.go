package implement

import (
	"context"

	"github.com/InstaySystem/is-be/internal/common"
	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/repository"
	"github.com/InstaySystem/is-be/internal/service"
	"github.com/InstaySystem/is-be/internal/types"
	"github.com/InstaySystem/is-be/pkg/bcrypt"
	"github.com/InstaySystem/is-be/pkg/snowflake"
	"go.uber.org/zap"
)

type userSvcImpl struct {
	userRepo repository.UserRepository
	sfGen    snowflake.Generator
	logger   *zap.Logger
	bHash    bcrypt.Hasher
}

func NewUserService(userRepo repository.UserRepository, sfGen snowflake.Generator, logger *zap.Logger, bHash bcrypt.Hasher) service.UserService {
	return &userSvcImpl{
		userRepo,
		sfGen,
		logger,
		bHash,
	}
}

func (s *userSvcImpl) CreateUser(ctx context.Context, req types.CreateUserRequest) (int64, error) {
	hashedPass, err := s.bHash.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("hash password failed", zap.Error(err))
		return 0, err
	}

	id, err := s.sfGen.NextID()
	if err != nil {
		s.logger.Error("generate ID failed", zap.Error(err))
		return 0, err
	}

	user := &model.User{
		ID:        id,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPass,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	}

	if err = s.userRepo.Create(ctx, user); err != nil {
		ok, constraint := common.IsUniqueViolation(err)
		if ok {
			switch constraint {
			case "users_email_key":
				return 0, common.ErrEmailAlreadyExists
			case "users_username_key":
				return 0, common.ErrUsernameAlreadyExists
			}
		}
		s.logger.Error("create user failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}

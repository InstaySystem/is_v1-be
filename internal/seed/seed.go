package seed

import (
	"context"
	"time"

	"github.com/InstaySystem/is_v1-be/internal/config"
	"github.com/InstaySystem/is_v1-be/internal/model"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	"github.com/InstaySystem/is_v1-be/pkg/bcrypt"
	"github.com/InstaySystem/is_v1-be/pkg/snowflake"
	"go.uber.org/zap"
)

type Seed struct {
	cfg      *config.Config
	userRepo repository.UserRepository
	logger   *zap.Logger
	bHash    bcrypt.Hasher
	sfGen    snowflake.Generator
}

func NewSeed(
	cfg *config.Config,
	userRepo repository.UserRepository,
	logger *zap.Logger,
	bHash bcrypt.Hasher,
	sfGen snowflake.Generator,
) *Seed {
	return &Seed{
		cfg,
		userRepo,
		logger,
		bHash,
		sfGen,
	}
}

func (s *Seed) AdminSeed() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := s.userRepo.ExistsActiveAdmin(ctx)
	if err != nil {
		s.logger.Error("check active admin failed", zap.Error(err))
	}
	if exists {
		s.logger.Info("Admin already exists, skipping")
		return nil
	}

	hashedPass, err := s.bHash.HashPassword(s.cfg.Admin.Password)
	if err != nil {
		s.logger.Error("hash password failed", zap.Error(err))
		return err
	}

	id, err := s.sfGen.NextID()
	if err != nil {
		s.logger.Error("generate admin id failed", zap.Error(err))
		return err
	}

	admin := &model.User{
		ID:        id,
		Username:  s.cfg.Admin.Username,
		Email:     s.cfg.Admin.Email,
		Role:      "admin",
		FirstName: "Main",
		LastName:  "Administrator",
		Phone:     "0123456789",
		Password:  hashedPass,
		IsActive:  true,
	}
	if err = s.userRepo.Create(ctx, admin); err != nil {
		s.logger.Error("create admin failed", zap.Error(err))
		return err
	}

	s.logger.Info("Admin created successfully")
	return nil
}

package implement

import (
	"github.com/InstaySystem/is-be/internal/repository"
	"github.com/InstaySystem/is-be/internal/service"
)

type authSvcImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) service.AuthService {
	return &authSvcImpl{userRepo}
}
package container

import (
	"github.com/InstaySystem/is-be/internal/handler"
	repoImpl "github.com/InstaySystem/is-be/internal/repository/implement"
	svcImpl "github.com/InstaySystem/is-be/internal/service/implement"
	"gorm.io/gorm"
)

type AuthContainer struct {
	Hdl *handler.AuthHandler
}

func NewAuthContainer(db *gorm.DB) *AuthContainer {
	userRepo := repoImpl.NewUserRepository(db)
	svc := svcImpl.NewAuthService(userRepo)
	hdl := handler.NewAuthHandler(svc)

	return &AuthContainer{hdl}
}

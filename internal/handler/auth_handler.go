package handler

import "github.com/InstaySystem/is-be/internal/service"

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc}
}

package service

import (
	"context"

	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/types"
)

type ChatService interface {
	CreateMessage(ctx context.Context, clientID int64, departmentID *int64, senderType string, req types.CreateMessageRequest) (*model.Message, error)
}
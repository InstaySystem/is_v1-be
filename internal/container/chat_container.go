package container

import (
	"github.com/InstaySystem/is_v1-be/internal/handler"
	"github.com/InstaySystem/is_v1-be/internal/repository"
	"github.com/InstaySystem/is_v1-be/internal/service"
	svcImpl "github.com/InstaySystem/is_v1-be/internal/service/implement"
	"github.com/InstaySystem/is_v1-be/pkg/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ChatContainer struct {
	Hdl *handler.ChatHandler
	Svc service.ChatService
}

func NewChatContainer(
	db *gorm.DB,
	chatRepo repository.ChatRepository,
	orderRepo repository.OrderRepository,
	userRepo repository.UserRepository,
	sfGen snowflake.Generator,
	logger *zap.Logger,
) *ChatContainer {
	svc := svcImpl.NewChatService(db, chatRepo, orderRepo, userRepo, sfGen, logger)
	hdl := handler.NewChatHandler(svc)

	return &ChatContainer{
		hdl,
		svc,
	}
}

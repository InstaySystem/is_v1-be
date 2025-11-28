package implement

import (
	"context"
	"errors"

	"github.com/InstaySystem/is-be/internal/common"
	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/repository"
	"github.com/InstaySystem/is-be/internal/service"
	"github.com/InstaySystem/is-be/internal/types"
	"github.com/InstaySystem/is-be/pkg/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type chatSvcImpl struct {
	db        *gorm.DB
	chatRepo  repository.ChatRepository
	orderRepo repository.OrderRepository
	sfGen     snowflake.Generator
	logger    *zap.Logger
}

func NewChatService(
	db *gorm.DB,
	chatRepo repository.ChatRepository,
	orderRepo repository.OrderRepository,
	sfGen snowflake.Generator,
	logger *zap.Logger,
) service.ChatService {
	return &chatSvcImpl{
		db,
		chatRepo,
		orderRepo,
		sfGen,
		logger,
	}
}

func (s *chatSvcImpl) CreateMessage(ctx context.Context, clientID int64, departmentID *int64, senderType string, req types.CreateMessageRequest) (*model.Message, error) {
	var message *model.Message

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if req.ChatID != nil {
			chat, err := s.getOrCreateChat(tx, req, clientID, departmentID, senderType)
			if err != nil {
				return err
			}

			messageID, err := s.sfGen.NextID()
			if err != nil {
				s.logger.Error("generate message id failed", zap.Error(err))
				return err
			}

			var senderID *int64
			if senderType == "staff" {
				senderID = &clientID
			}

			message = &model.Message{
				ID:         messageID,
				ChatID:     chat.ID,
				SenderType: senderType,
				SenderID:   senderID,
				ImageKey:   req.ImageKey,
				Content:    req.Content,
				Chat:       chat,
			}

			if err = s.chatRepo.CreateMessageTx(tx, message); err != nil {
				s.logger.Error("create message failed", zap.Error(err))
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *chatSvcImpl) getOrCreateChat(tx *gorm.DB, req types.CreateMessageRequest, clientID int64, departmentID *int64, senderType string) (*model.Chat, error) {
	if req.ChatID != nil {
		chat, err := s.chatRepo.FindChatByIDTx(tx, *req.ChatID)
		if err != nil {
			s.logger.Error("find chat by id failed", zap.Error(err))
			return nil, err
		}
		return chat, nil
	}

	if req.ReceiverID == nil {
		return nil, errors.New("receiverid is required for new chat")
	}

	orderRoomID := clientID
	if senderType == "staff" {
		orderRoomID = *req.ReceiverID
	}

	orderRoom, err := s.orderRepo.FindOrderRoomByIDWithBookingTx(tx, orderRoomID)
	if err != nil {
		return nil, common.ErrOrderRoomNotFound
	}

	chatID, err := s.sfGen.NextID()
	if err != nil {
		s.logger.Error("generate chat id failed", zap.Error(err))
		return nil, err
	}

	if departmentID == nil {
		return nil, errors.New("department_id is required")
	}

	chat := &model.Chat{
		ID:           chatID,
		OrderRoomID:  orderRoomID,
		DepartmentID: *departmentID,
		ExpiredAt:    orderRoom.Booking.CheckOut,
	}

	if err = s.chatRepo.CreateChatTx(tx, chat); err != nil {
		s.logger.Error("create chat failed", zap.Error(err))
		return nil, err
	}

	return chat, nil
}

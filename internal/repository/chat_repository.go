package repository

import (
	"github.com/InstaySystem/is-be/internal/model"
	"gorm.io/gorm"
)

type ChatRepository interface {
	CreateChatTx(tx *gorm.DB, chat *model.Chat) error

	CreateMessageTx(tx *gorm.DB, message *model.Message) error

	FindChatByIDTx(tx *gorm.DB, chatID int64) (*model.Chat, error)
}

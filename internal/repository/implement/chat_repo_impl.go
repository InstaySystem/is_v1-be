package implement

import (
	"errors"

	"github.com/InstaySystem/is-be/internal/model"
	"github.com/InstaySystem/is-be/internal/repository"
	"gorm.io/gorm"
)

type chatRepoImpl struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) repository.ChatRepository {
	return &chatRepoImpl{db}
}

func (r *chatRepoImpl) CreateChatTx(tx *gorm.DB, chat *model.Chat) error {
	return tx.Create(chat).Error
}

func (r *chatRepoImpl) CreateMessageTx(tx *gorm.DB, message *model.Message) error {
	return tx.Create(message).Error
}

func (r *chatRepoImpl) FindChatByIDTx(tx *gorm.DB, chatID int64) (*model.Chat, error) {
	var chat model.Chat
	if err := tx.Where("id = ?", chatID).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &chat, nil
}

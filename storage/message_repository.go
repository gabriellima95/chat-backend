package storage

import (
	"msn/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return MessageRepository{
		DB: db,
	}
}

func (r MessageRepository) Create(message *models.Message) error {
	message.ID = uuid.New()
	return r.DB.Create(message).Error
}

func (r MessageRepository) SaveAttachment(attachment *models.Attachment) error {
	attachment.ID = uuid.New()
	return r.DB.Create(attachment).Error
}

func (r MessageRepository) ListByChatID(id uuid.UUID) ([]models.Message, error) {
	var messages []models.Message
	err := r.DB.Preload("Sender").Preload("Attachments").Where("chat_id = ?", id.String()).Order("created_at desc").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

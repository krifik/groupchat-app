package repository

import (
	"mangojek-backend/entity"
)

type MessageRepository interface {
	SendMessage(message *entity.Message)
	GetMessages() (messages []entity.Message)
}

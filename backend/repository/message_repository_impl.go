package repository

import (
	"mangojek-backend/config"
	"mangojek-backend/entity"
	"mangojek-backend/exception"

	"gorm.io/gorm"
)

type MessageRepositoryImpl struct {
	DB *gorm.DB
}

func NewMessageRepositoryImpl(db *gorm.DB) MessageRepository {
	return &MessageRepositoryImpl{DB: db}
}

func (repository *MessageRepositoryImpl) GetMessages() (messages []entity.Message) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	err := repository.DB.WithContext(ctx).Find(&messages)
	exception.PanicIfNeeded(err.Error)
	return messages
}
func (repository *MessageRepositoryImpl) SendMessage(message *entity.Message) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	result := repository.DB.WithContext(ctx).Create(&message)
	exception.PanicIfNeeded(result.Error)
}

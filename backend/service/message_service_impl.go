package service

import (
	"fmt"
	"mangojek-backend/entity"
	"mangojek-backend/model"
	"mangojek-backend/repository"
)

type MessageServiceImpl struct {
	MessageRepository repository.MessageRepository
}

func NewMessageRepositoryImpl(messageRepository repository.MessageRepository) MessageService {
	return &MessageServiceImpl{MessageRepository: messageRepository}
}

func (service *MessageServiceImpl) SendMessage(request model.CreateMessageRequest) (response model.CreateMessageResponse) {
	// validation here
	message := entity.Message{
		UserID:  request.UserID,
		Content: request.Content,
	}
	service.MessageRepository.SendMessage(&message)
	fmt.Println(message)
	response = model.CreateMessageResponse{
		Id:        int(message.ID),
		UserID:    message.UserID,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
		DeletedAt: message.DeletedAt,
	}
	return response
}

func (service *MessageServiceImpl) GetMessages() (responses []model.GetMessageResponse) {
	messages := service.MessageRepository.GetMessages()
	for _, message := range messages {
		responses = append(responses, model.GetMessageResponse{
			Id:        int(message.ID),
			UserID:    message.UserID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
			UpdatedAt: message.UpdatedAt,
			DeletedAt: message.DeletedAt,
		})
	}
	return responses
}

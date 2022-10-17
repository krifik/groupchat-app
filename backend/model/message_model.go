package model

import (
	"time"

	"gorm.io/gorm"
)

type CreateMessageRequest struct {
	Id        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CreateMessageResponse struct {
	Id        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type GetMessageResponse struct {
	Id        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain"
	"github.com/jinzhu/gorm"
)

type MessageRepository interface {
	Create(data *domain.Message) (int64, error)
	GetAll() ([]*domain.Message, error)
	GetByID(id int64) (*domain.Message, error)
	Delete(id int64) (int64, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) Create(data *domain.Message) (int64, error) {
	result := r.db.Create(data)
	return data.ID, result.Error
}

func (r *messageRepository) GetAll() ([]*domain.Message, error) {
	var messages []*domain.Message
	result := r.db.Find(&messages)
	return messages, result.Error
}

func (r *messageRepository) GetByID(id int64) (*domain.Message, error) {
	var message domain.Message
	result := r.db.First(&message, id)
	return &message, result.Error
}

func (r *messageRepository) Delete(id int64) (int64, error) {
	result := r.db.Delete(&domain.Message{}, id)
	return id, result.Error
}

package repository

type MessageRepository interface {
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

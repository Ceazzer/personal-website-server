package repository

import "github.com/Ceazzer/personal-website-server/core/domain"

type MessageRepoResponse struct {
	FindByID func(id int) (*domain.Message, error)
	Create   func(entity *domain.Message) (string, error)
	Delete   func(id int) (string, error)
}

type MessageRepoFunc func() (*MessageRepoResponse, error)

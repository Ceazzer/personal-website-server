package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
)

type MessageRepoResponse struct {
	FindByID func(id int) (*entity.Message, error)
	Create   func(entity *entity.Message) (string, error)
}

type MessageRepoFunc func() (*MessageRepoResponse, error)

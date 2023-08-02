package repository

import "github.com/Ceazzer/personal-website-server/core/domain/dto"

type MessageRepoResponse struct {
	FindByID func(id int) (*dto.MessageDTO, error)
	Create   func(data *dto.MessageDTO) (string, error)
}

type MessageRepoFunc func() (*MessageRepoResponse, error)

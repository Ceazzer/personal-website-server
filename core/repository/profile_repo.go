package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
)

type ProfileRepoResponse struct {
	Create func(data *entity.Profile) (string, error)
}

type ProfileRepoFunc func() (*ProfileRepoResponse, error)

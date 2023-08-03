package repository

import (
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
)

type ProfileRepo struct {
	Create func(data *entity.Profile) (int64, error)
	Delete func(id int64) (int64, error)
}

type ProfileRepoFunc func() (*ProfileRepo, error)

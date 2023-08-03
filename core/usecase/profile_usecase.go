package usecase

import (
	"context"

	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/repository"
)

type ProfileOpts struct {
	Repo repository.ProfileRepoFunc
}

type CreateProfileFunc func(ctx context.Context, data *entity.Profile, opts ProfileOpts) (string, error)

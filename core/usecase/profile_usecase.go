package usecase

import (
	"context"

	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/repository"
)

type ProfileUsecase struct {
	CreateProfile func(ctx context.Context, data *entity.Profile) (int64, error)
	DeleteProfile func(ctx context.Context, id int64) (int64, error)
}

type ProfileUsecaseOpts struct {
	Repo repository.ProfileRepo
}

type ProfileUseCaseFunc func(opts *ProfileUsecaseOpts) (*ProfileUsecase, error)

package usecaseimpl

import (
	"context"

	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/usecase"
)

func ProfileUsecase(opts *usecase.ProfileUsecaseOpts) (*usecase.ProfileUsecase, error) {

	createProfile := func(ctx context.Context, data *entity.Profile) (int64, error) {
		return opts.Repo.Create(data)
	}

	deleteProfile := func(ctx context.Context, id int64) (int64, error) {
		return opts.Repo.Delete(id)
	}

	return &usecase.ProfileUsecase{
		CreateProfile: createProfile,
		DeleteProfile: deleteProfile,
	}, nil
}

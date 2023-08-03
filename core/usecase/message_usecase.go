package usecase

import (
	"context"

	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/repository"
)

type Opts struct {
	Repo repository.MessageRepoFunc
}

type CreateMessageFunc func(ctx context.Context, entity *entity.Message, opts Opts) (string, error)
type FindMessageByIDFunc func(ctx context.Context, id string, opts Opts) (*entity.Message, error)

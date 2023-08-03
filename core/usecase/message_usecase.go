package usecase

import (
	"context"

	"github.com/Ceazzer/personal-website-server/core/domain/dto"
	"github.com/Ceazzer/personal-website-server/core/domain/entity"
	"github.com/Ceazzer/personal-website-server/core/repository"
	"github.com/Ceazzer/personal-website-server/core/service"
)

type MessageOpts struct {
	ResendService service.ResendServiceFunc
	MessageRepo   repository.MessageRepoFunc
}

type CreateMessageFunc func(ctx context.Context, entity *entity.Message, opts MessageOpts) (string, error)
type FindMessageByIDFunc func(ctx context.Context, id string, opts MessageOpts) (*dto.MessageDTO, error)
type FindMessageByProfileIDFunc func(ctx context.Context, id string, opts MessageOpts) ([]*dto.MessageDTO, error)
type SendMessageToEmailFunc func(ctx context.Context, entity *dto.MessageDTO, opts MessageOpts) error

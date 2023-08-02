package dto

import "github.com/Ceazzer/personal-website-server/core/domain/entity"

type MessageDTO struct {
	*entity.Message `json:",inline"`
	Profile         *entity.Profile `json:"profile"`
}

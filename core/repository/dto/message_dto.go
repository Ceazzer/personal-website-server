package repositorydto

import "github.com/Ceazzer/personal-website-server/core/domain"

type MessageDTO struct {
	*domain.Message `json:",inline"`
	Profile         *domain.Profile `json:"profile"`
}

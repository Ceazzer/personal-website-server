package service

import (
	"github.com/Ceazzer/personal-website-server/core/domain/dto"
	"github.com/resendlabs/resend-go"
)

type ResendServiceResponse struct {
	SendEmail func(data *dto.MessageDTO) (string, error)
}

type ResendServiceFunc func(client *resend.Client) (*ResendServiceResponse, error)

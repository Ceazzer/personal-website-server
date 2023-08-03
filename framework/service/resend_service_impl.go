package serviceimpl

import (
	"os"

	"github.com/Ceazzer/personal-website-server/core/domain/dto"
	"github.com/Ceazzer/personal-website-server/core/service"
	"github.com/resendlabs/resend-go"
)

func ResendService(client *resend.Client) (*service.ResendServiceResponse, error) {

	sendEmail := func(data *dto.MessageDTO) (string, error) {
		message := &resend.SendEmailRequest{
			From: data.Profile.Email,
			To:   []string{os.Getenv("RESEND_EMAIL")},
		}

		sent, err := client.Emails.Send(message)
		if err != nil {
			return "", err
		}

		return sent.Id, nil
	}

	return &service.ResendServiceResponse{
		SendEmail: sendEmail,
	}, nil
}

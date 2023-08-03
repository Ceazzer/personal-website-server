package config

import (
	"os"

	"github.com/resendlabs/resend-go"
)

var ResendClient *resend.Client

func ResendSetup() {
	ResendClient = resend.NewClient(os.Getenv("RESEND_API_KEY"))
}

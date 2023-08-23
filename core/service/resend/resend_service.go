package service

type ResendService interface {
	SendMail()
}

type resendService struct{}

func NewResendService() ResendService {
	return &resendService{}
}

func (s *resendService) SendMail() {}

package email_service

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"os"
)

type emailService struct {
	SMTPHost    string
	SMTPPort    string
	SenderEmail string
	SenderPass  string
}

func NewEmailnService() *emailService {
	return &emailService{
		SMTPHost:    "smtp.gmail.com",
		SMTPPort:    "587",
		SenderEmail: os.Getenv("SMTP_EMAIL"),
		SenderPass:  os.Getenv("SMTP_PASSWORD"),
	}
}

type EmailService interface {
	SendEmail(to, subject, body string) *rest_err.RestErr
}

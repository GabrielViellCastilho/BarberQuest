package email_service

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"go.uber.org/zap"
	"net/smtp"
)

func (m *emailService) SendEmail(to, subject, body string) *rest_err.RestErr {
	logger.Info("Init email service",
		zap.String("journey", "email"))

	auth := smtp.PlainAuth("", m.SenderEmail, m.SenderPass, m.SMTPHost)

	msg := []byte("Subject: " + subject + "\r\n" +
		"From: " + m.SenderEmail + "\r\n" +
		"To: " + to + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
		body)

	err := smtp.SendMail(m.SMTPHost+":"+m.SMTPPort, auth, m.SenderEmail, []string{to}, msg)
	if err != nil {
		logger.Error("Don't possible to send email.", err)
		return rest_err.NewInternalServerError("Don't possible to send email.")
	}

	logger.Info("Email sent", zap.String("to", to))
	logger.Info("Successful email service",
		zap.String("journey", "email"))
	return nil
}

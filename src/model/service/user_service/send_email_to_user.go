package user_service

import (
	"fmt"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/email_service"
	"os"
)

const URL = "URL"

func (ud *userDomainService) SendEmailResetPassword(email, token string) *rest_err.RestErr {

	u := email_service.NewEmailnService()

	link := fmt.Sprintf("%s/updatePassword/%s", os.Getenv(URL), token)

	subject := "Redefinição de Senha - Spartan Barbearia"
	body := fmt.Sprintf("Recebemos uma solicitação para redefinir sua senha na Spartan Barbearia.\n\n"+
		"Clique no link para redefinir sua senha: %s \n\n"+
		"⚠️ Este link é válido por apenas 15 minutos. Se não for utilizado dentro desse período, você precisará solicitar uma nova redefinição de senha.\n\n"+
		"Se você não solicitou essa alteração, ignore este e-mail. Sua senha permanecerá a mesma.\n\n"+
		"Atenciosamente,\n"+
		"Equipe Spartan Barbearia", link)

	return u.SendEmail(email, subject, body)
}

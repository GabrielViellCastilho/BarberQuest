package user_controller

import (
	"net/http"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) SendEmailResetPassword(c *gin.Context) {
	logger.Info("Init SendEmailResetPassword controller",
		zap.String("journey", "sendEmailResetPassword"))

	email := c.Param("email")

	token, err := user_domain.GeneratePasswordResetToken(email)
	if err != nil {

	}

	err = uc.service.SendEmailResetPassword(email, token)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("SendEmailResetPassword successfully",
		zap.String("journey", "sendEmailResetPassword"))

	c.JSON(http.StatusOK, nil)
}

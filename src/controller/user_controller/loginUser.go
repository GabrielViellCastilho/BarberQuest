package user_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/validation"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/request"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"))

	var userRequest request.UserRequestLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	userdomain := user_domain.NewUserDomainLogin(userRequest.Email, userRequest.Password)

	response, jwt, erro := uc.service.FindUserByEmailAndPassword(c, userdomain)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.Header("Authorization", jwt)
	c.JSON(http.StatusOK, response)
	logger.Info("Successful loginUser controller", zap.String("journey", "loginUser"))
}

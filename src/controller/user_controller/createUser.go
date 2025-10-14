package user_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/validation"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/request"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	if user.GetRole() != "admin" {
		restErr := rest_err.NewBadRequestError("You are not admin")
		c.JSON(restErr.Code, restErr)
		return
	}

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := user_domain.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Role,
		userRequest.Cellphone)

	ur, err := uc.service.CreateUser(c, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, ur)
}

func (uc *userControllerInterface) CreateCustomerUser(c *gin.Context) {
	logger.Info("Init CreateCustomerUser controller",
		zap.String("journey", "createCustomerUser"))

	var userRequest request.CustomerUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := user_domain.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		"user",
		userRequest.Cellphone)

	domain.SetDateOfBirth(userRequest.DateOfBirth)

	ur, err := uc.service.CreateUser(c, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	jwt, erro := domain.GenerateToken()
	if erro != nil {
		logger.Error("Error trying to validate user_domain info", erro)
		erro := validation.ValidateError(err)

		c.JSON(erro.Code, erro)

		return
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "createCustomerUser"))

	c.Header("Authorization", jwt)
	c.JSON(http.StatusCreated, ur)
}

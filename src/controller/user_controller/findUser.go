package user_controller

import (
	"fmt"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
	"strconv"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"))

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

	id := c.Param("userId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	response, erro := uc.service.FindUserById(c, idInt)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findUserByID controller", zap.String("journey", "findUserByID"))
}

func (uc *userControllerInterface) FindMyUserData(c *gin.Context) {
	logger.Info("Init findMyUserData controller",
		zap.String("journey", "findMyUserData"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	logger.Info(fmt.Sprintf("Find user_domain by id: %v", user))

	response, erro := uc.service.FindUserById(c, user.GetID())
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findUserByID controller", zap.String("journey", "findMyUserData"))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"))

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

	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Warn("Invalid email format", err, zap.String("email", email), zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("Invalid email format. Please provide a valid email address.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	response, err := uc.service.FindUserByEmail(c, email)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findUserByEmail controller", zap.String("journey", "findUserByEmail"))
}

func (uc *userControllerInterface) FindAllBarbers(c *gin.Context) {
	logger.Info("Init findAllBarbers controller",
		zap.String("journey", "findAllBarbers"))

	response, erro := uc.service.FindAllBarbers(c)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllBarbers controller", zap.String("journey", "findAllBarbers"))
}

func (uc *userControllerInterface) FindAllUsersByDateOfBirth(c *gin.Context) {
	logger.Info("Init findAllUsersByDateOfBirth controller",
		zap.String("journey", "findAllUsersByDateOfBirth"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	if user.GetRole() != "admin" {
		restErr := rest_err.NewUnauthorizedRequestError("You are not admin")
		c.JSON(restErr.Code, restErr)
		return
	}

	dateOfBirth := c.Param("dateOfBirth")

	response, erro := uc.service.FindAllUsersByDateOfBirth(c, dateOfBirth)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllUsersByDateOfBirth controller", zap.String("journey", "findAllUsersByDateOfBirth"))
}

package service_controller

import (
	"fmt"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (sc *serviceController) FindServiceByID(c *gin.Context) {
	logger.Info("Init findServiceByID controller",
		zap.String("journey", "findServiceByID"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	logger.Info(fmt.Sprintf("Find user_domain by id: %v", user))

	if user.GetRole() != "admin" {
		restErr := rest_err.NewBadRequestError("You are not admin")
		c.JSON(restErr.Code, restErr)
		return
	}

	id := c.Param("serviceId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "findServiceByID"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	response, erro := sc.service.FindServiceByID(c, idInt)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findServiceByID controller", zap.String("journey", "findServiceByID"))
}

func (sc *serviceController) FindServiceByName(c *gin.Context) {
	logger.Info("Init findServiceByName controller",
		zap.String("journey", "findServiceByName"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	logger.Info(fmt.Sprintf("Find user_domain by name: %v", user))

	if user.GetRole() != "admin" {
		restErr := rest_err.NewBadRequestError("You are not admin")
		c.JSON(restErr.Code, restErr)
		return
	}

	name := c.Param("serviceName")

	name = strings.ReplaceAll(name, "_", " ")

	response, erro := sc.service.FindServiceByName(c, name)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findServiceByName controller", zap.String("journey", "findServiceByName"))
}

func (sc *serviceController) FindAllServices(c *gin.Context) {
	logger.Info("Init findAllServices controller",
		zap.String("journey", "findAllServices"))

	response, erro := sc.service.FindAllServices(c)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllServices controller", zap.String("journey", "findAllServices"))
}

func (sc *serviceController) FindAllAvailableServices(c *gin.Context) {
	logger.Info("Init findAllAvailableServices controller",
		zap.String("journey", "findAllAvailableServices"))

	response, erro := sc.service.FindAllAvailableServices(c)
	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllAvailableServices controller", zap.String("journey", "findAllAvailableServices"))
}

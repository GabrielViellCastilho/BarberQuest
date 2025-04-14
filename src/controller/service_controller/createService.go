package service_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/validation"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/request"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (sc *serviceController) CreateService(c *gin.Context) {
	logger.Info("Init CreateService controller",
		zap.String("journey", "createService"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	if user.GetRole() != "admin" {
		restError := rest_err.NewUnauthorizedRequestError("Your role is not admin")
		c.JSON(restError.Code, restError)
		return
	}

	var serviceRequest request.ServiceRequest

	if err := c.ShouldBindJSON(&serviceRequest); err != nil {
		logger.Error("Error trying to validate user_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	var serviceDomain service_domain.ServiceDomainInterface

	if serviceRequest.Available {
		serviceDomain = service_domain.NewServiceDomain(serviceRequest.Name, serviceRequest.Price, serviceRequest.DurationMinutes, serviceRequest.Available)
	} else {
		serviceDomain = service_domain.NewServiceDomain(serviceRequest.Name, serviceRequest.Price, serviceRequest.DurationMinutes, false)
	}

	response, err := sc.service.CreateService(c, serviceDomain)
	if err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Service created successfully",
		zap.String("journey", "serviceUser"))

	c.JSON(http.StatusCreated, response)
}

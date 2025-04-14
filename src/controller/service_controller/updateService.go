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
	"strconv"
)

func (sc *serviceController) UpdateService(c *gin.Context) {
	logger.Info("Init UpdateService controller",
		zap.String("journey", "updateService"))

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

	id := c.Param("serviceId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "updateService"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var serviceRequest request.ServiceRequest

	if err := c.ShouldBindJSON(&serviceRequest); err != nil {
		logger.Error("Error trying to validate serviceRequest info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	var domain service_domain.ServiceDomainInterface

	if serviceRequest.Available {
		domain = service_domain.NewServiceDomain(serviceRequest.Name, serviceRequest.Price, serviceRequest.DurationMinutes, serviceRequest.Available)
	} else {
		domain = service_domain.NewServiceDomain(serviceRequest.Name, serviceRequest.Price, serviceRequest.DurationMinutes, false)
	}

	domain.SetID(idInt)

	if err := sc.service.UpdateService(c, domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Service updated successfully",
		zap.String("journey", "updateService"), zap.Int("serviceID", idInt))

	c.JSON(http.StatusOK, nil)
}

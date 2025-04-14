package appointment_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func (ac *appointmentController) FindAllAppointmentsByDateAndBarberID(c *gin.Context) {
	logger.Info("Init findAllAppointmentsByDateAndBarberID controller",
		zap.String("journey", "findAllAppointmentsByDateAndBarberID"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	dateString := c.Param("date")

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logger.Error("Failed to convert string to time.Time", err, zap.String("dateString", dateString), zap.String("journey", "findAllAppointmentsByDateAndBarberID"))
		er := *rest_err.NewBadRequestError("Failed to convert string date")
		c.JSON(er.Code, er)
		return
	}

	response, erro := ac.service.FindAllAppointmentsByDateAndBarberID(c, user.GetID(), date)

	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllAppointmentsByDateAndBarberID controller", zap.String("journey", "findAllAppointmentsByDateAndBarberID"))
}

func (ac *appointmentController) FindAllAppointmentsByUserID(c *gin.Context) {
	logger.Info("Init findAllAppointmentsByUserID controller",
		zap.String("journey", "findAllAppointmentsByUserID"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	response, erro := ac.service.FindAllAppointmentsByUserID(c, user.GetID())

	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllAppointmentsByUserID controller", zap.String("journey", "findAllAppointmentsByUserID"))
}

func (ac *appointmentController) FindAllHistoricAppointmentsByUserID(c *gin.Context) {
	logger.Info("Init findAllHistoricAppointmentsByUserID controller",
		zap.String("journey", "findAllHistoricAppointmentsByUserID"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	response, erro := ac.service.FindAllHistoricAppointmentsByUserID(c, user.GetID())

	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllHistoricAppointmentsByUserID controller", zap.String("journey", "findAllHistoricAppointmentsByUserID"))
}

func (ac *appointmentController) FindAllAppointmentsByDateAndBarberIDByHeader(c *gin.Context) {
	logger.Info("Init findAllAppointmentsByDateAndBarberIDByHeader controller",
		zap.String("journey", "findAllAppointmentsByDateAndBarberIDByHeader"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	if user.GetRole() != "admin" {
		erroUnauthorized := rest_err.NewUnauthorizedRequestError("You do not have admin role")
		c.JSON(erroUnauthorized.Code, erroUnauthorized)
		return
	}

	dateString := c.Param("date")

	idString := c.Param("id")

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logger.Error("Failed to convert string to time.Time", err, zap.String("dateString", dateString), zap.String("journey", "findAllAppointmentsByDateAndBarberIDByHeader"))
		er := *rest_err.NewBadRequestError("Failed to convert string date")
		c.JSON(er.Code, er)
		return
	}

	idInt, err0r := strconv.Atoi(idString)
	if err0r != nil {
		logger.Warn("Invalid serviceID format", err0r, zap.String("id", idString), zap.String("journey", "checkAvailableSlots"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	response, erro := ac.service.FindAllAppointmentsByDateAndBarberID(c, idInt, date)

	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllAppointmentsByDateAndBarberIDByHeader controller", zap.String("journey", "findAllAppointmentsByDateAndBarberIDByHeader"))
}

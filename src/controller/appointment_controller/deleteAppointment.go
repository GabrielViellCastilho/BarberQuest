package appointment_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (ac *appointmentController) DeleteAppointmentByIdAndUserId(c *gin.Context) {
	logger.Info("Init DeleteAppointmentByIdAndUserId controller",
		zap.String("journey", "deleteAppointmentByIdAndUserId"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	id := c.Param("appointmentId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "DeleteAppointmentByIdAndUserId"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := ac.service.DeleteAppointmentByIdAndUserID(c, idInt, user.GetID()); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Appointment deleted successfully",
		zap.String("journey", "deleteAppointmentByIdAndUserId"), zap.Int("appointmentID", idInt))

	c.JSON(http.StatusOK, nil)

}

func (ac *appointmentController) DeleteAppointmentByIdAndBarberId(c *gin.Context) {
	logger.Info("Init DeleteAppointmentByIdAndBarberId controller",
		zap.String("journey", "deleteAppointmentByIdAndBarberId"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	if user.GetRole() != "barber" {
		restErr := rest_err.NewUnauthorizedRequestError("Unauthorized")
		c.JSON(restErr.Code, restErr)
		return
	}

	id := c.Param("appointmentId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "DeleteAppointmentByIdAndBarberId"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := ac.service.DeleteAppointmentByIdAndBarberID(c, idInt, user.GetID()); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Appointment deleted successfully",
		zap.String("journey", "deleteAppointmentByIdAndBarberId"), zap.Int("appointmentID", idInt))

	c.JSON(http.StatusOK, nil)

}

package appointment_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/validation"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/request"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/appointment_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func (ac *appointmentController) UpdateAppointmentCompletedById(c *gin.Context) {
	logger.Info("Init UpdateAppointmentCompletedById controller",
		zap.String("journey", "updateAppointmentCompletedById"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	if user.GetRole() != "barber" {
		restErr := rest_err.NewUnauthorizedRequestError("You are not barber")
		c.JSON(restErr.Code, restErr)
		return
	}

	id := c.Param("appointmentId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "UpdateAppointmentCompletedById"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var appointmentRequest request.UpdateCompletedAppointmentRequest

	if err := c.ShouldBindJSON(&appointmentRequest); err != nil {
		logger.Error("Error trying to validate appointmentRequest info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := appointment_domain.NewAppointmentDomain("", "", user.GetID(), 0, time.Time{}, 0, appointmentRequest.Completed)

	domain.SetId(idInt)

	if err := ac.service.UpdateAppointmentCompletedById(c, domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Appointment updated successfully",
		zap.String("journey", "updateAppointmentCompletedById"), zap.Int("appointmentID", idInt))

	c.JSON(http.StatusOK, nil)
}

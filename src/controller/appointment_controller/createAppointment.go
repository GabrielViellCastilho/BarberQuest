package appointment_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/validation"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/request"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/appointment_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (ac *appointmentController) CreateAppointment(c *gin.Context) {
	logger.Info("Init CreateAppointment controller",
		zap.String("journey", "createAppointment"))

	// 🔹 Verifica e valida o JWT do usuário
	jwtUser, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	// 🔹 Obtém os dados do usuário pelo ID
	userDomain, err := ac.userService.FindUserById(c, jwtUser.GetID())
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	// 🔹 Faz o bind do JSON para a struct request
	var appointmentRequest request.CreateAppointmentRequest
	if err := c.ShouldBindJSON(&appointmentRequest); err != nil {
		logger.Error("Error trying to validate appointment_domain info", err)
		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// 🔹 Define os dados do agendamento com base na role
	var appointmentDomain appointment_domain.AppointmentDomainInterface
	if userDomain.Role == "user" {
		appointmentRequest.ClientName = userDomain.Name
		appointmentRequest.ClientContact = userDomain.Cellphone
		appointmentDomain = appointment_domain.NewAppointmentDomain(
			appointmentRequest.ClientName, appointmentRequest.ClientContact,
			appointmentRequest.BarberID, appointmentRequest.ServiceID,
			appointmentRequest.AppointmentDate, userDomain.ID, nil,
		)
	} else if userDomain.Role == "barber" || userDomain.Role == "admin" {
		appointmentDomain = appointment_domain.NewAppointmentDomain(
			appointmentRequest.ClientName, appointmentRequest.ClientContact,
			appointmentRequest.BarberID, appointmentRequest.ServiceID,
			appointmentRequest.AppointmentDate, 0, nil,
		)
	}

	// 🔹 Chama o service para criar o agendamento
	response, err := ac.service.CreateAppointment(c, appointmentDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Appointment created successfully",
		zap.String("journey", "createAppointment"))

	c.JSON(http.StatusCreated, response)
}

package appointment_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service/appointment_service"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service/user_service"
	"github.com/gin-gonic/gin"
)

type appointmentController struct {
	service     appointment_service.AppointmentService
	userService user_service.UserDomainService
}

func NewAppointmentController(service appointment_service.AppointmentService, userService user_service.UserDomainService) *appointmentController {
	return &appointmentController{
		service:     service,
		userService: userService,
	}
}

type AppointmentController interface {
	CreateAppointment(c *gin.Context)
	FindAllAppointmentsByDateAndBarberID(c *gin.Context)
	FindAllAppointmentsByUserID(c *gin.Context)
	FindAllHistoricAppointmentsByUserID(c *gin.Context)
	FindAllAppointmentsByDateAndBarberIDByHeader(c *gin.Context)
	UpdateAppointmentCompletedById(c *gin.Context)
	CheckAvailableSlots(ctx *gin.Context)
	DeleteAppointmentByIdAndUserId(c *gin.Context)
	DeleteAppointmentByIdAndBarberId(c *gin.Context)
}

package availability_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service/availability_service"
	"github.com/gin-gonic/gin"
)

type availabilityController struct {
	service availability_service.AvailabilityServiceService
}

func NewAvailabilityController(service availability_service.AvailabilityServiceService) *availabilityController {
	return &availabilityController{service: service}
}

type AvailabilityController interface {
	CreateAvailability(c *gin.Context)
	FindAllAvailabilityById(c *gin.Context)
	UpdateAvailabilityByIdAndIdBarber(c *gin.Context)
	DeleteAvailabilityById(c *gin.Context)
}

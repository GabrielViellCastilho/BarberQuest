package service_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/barber_service_service"
	"github.com/gin-gonic/gin"
)

type serviceController struct {
	service barber_service_service.BarberServiceService
}

func NewServiceController(service barber_service_service.BarberServiceService) *serviceController {
	return &serviceController{service: service}
}

type ServiceController interface {
	CreateService(c *gin.Context)
	FindServiceByID(c *gin.Context)
	FindServiceByName(c *gin.Context)
	FindAllServices(c *gin.Context)
	FindAllAvailableServices(c *gin.Context)
	UpdateService(c *gin.Context)
	DeleteService(c *gin.Context)
}

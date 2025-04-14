package special_schedule_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service/special_schedule_service"
	"github.com/gin-gonic/gin"
)

type specialScheduleController struct {
	service special_schedule_service.SpecialScheduleService
}

func NewSpecialScheduleController(service special_schedule_service.SpecialScheduleService) *specialScheduleController {
	return &specialScheduleController{
		service: service,
	}
}

type SpecialScheduleController interface {
	CreateSpecialSchedule(c *gin.Context)
	FindAllSpecialScheduleById(c *gin.Context)
	DeleteSpecialScheduleById(c *gin.Context)
}

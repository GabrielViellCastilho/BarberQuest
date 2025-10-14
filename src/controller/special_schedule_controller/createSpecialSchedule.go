package special_schedule_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/validation"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/request"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/special_schedule_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (ssc *specialScheduleController) CreateSpecialSchedule(c *gin.Context) {
	logger.Info("Init CreateSpecialSchedule controller",
		zap.String("journey", "createSpecialSchedule"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	var specialScheduleRequest request.SpecialSchedulerRequest

	if err := c.ShouldBindJSON(&specialScheduleRequest); err != nil {
		logger.Error("Error trying to validate special_schedule_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	specialScheduleDomain := special_schedule_domain.NewSpecialScheduleDomain(user.GetID(), specialScheduleRequest.Date, specialScheduleRequest.OpeningTime, specialScheduleRequest.ClosedTime, specialScheduleRequest.BreakStartTime, specialScheduleRequest.BreakEndTime)

	response, err := ssc.service.CreateSpecialSchedule(c, specialScheduleDomain)
	if err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Service special schedule successfully",
		zap.String("journey", "createSpecialSchedule"))

	c.JSON(http.StatusCreated, response)
}

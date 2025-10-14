package availability_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/validation"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/request"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/availability_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (ac *availabilityController) CreateAvailability(c *gin.Context) {
	logger.Info("Init CreateAvailability controller",
		zap.String("journey", "createAvailability"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	var availabilityRequest request.AvailabilityRequest

	if err := c.ShouldBindJSON(&availabilityRequest); err != nil {
		logger.Error("Error trying to validate availability_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	availabilityDomain := availability_domain.NewAvailabilityDomain(user.GetID(), availabilityRequest.DayOfWeek,
		availabilityRequest.StartTime, availabilityRequest.EndTime, availabilityRequest.BreakStartTime, availabilityRequest.BreakEndTime)

	response, err := ac.service.CreateAvailability(c, availabilityDomain)
	if err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Availability created successfully",
		zap.String("journey", "createAvailability"))

	c.JSON(http.StatusCreated, response)
}

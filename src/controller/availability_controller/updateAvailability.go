package availability_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/validation"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/request"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/availability_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (ac *availabilityController) UpdateAvailabilityByIdAndIdBarber(c *gin.Context) {
	logger.Info("Init UpdateAvailabilityByIdAndIdBarber controller",
		zap.String("journey", "updateAvailabilityByIdAndIdBarber"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	id := c.Param("availabilityId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "updateAvailabilityByIdAndIdBarber"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var availabilityRequest request.AvailabilityRequestUpdate

	if err := c.ShouldBindJSON(&availabilityRequest); err != nil {
		logger.Error("Error trying to validate availabilityRequest info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := availability_domain.NewAvailabilityDomain(
		user.GetID(),
		0,
		availabilityRequest.StartTime,
		availabilityRequest.EndTime,
		availabilityRequest.BreakStartTime,
		availabilityRequest.BreakEndTime,
	)

	domain.SetID(idInt)

	if err := ac.service.UpdateAvailabilityByIdAndIdBarber(c, domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Availability updated successfully",
		zap.String("journey", "updateAvailabilityByIdAndIdBarber"), zap.Int("availabilityID", idInt))

	c.JSON(http.StatusOK, nil)
}

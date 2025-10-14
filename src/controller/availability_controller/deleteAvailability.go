package availability_controller

import (
	"net/http"
	"strconv"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *availabilityController) DeleteAvailabilityById(c *gin.Context) {
	logger.Info("Init DeleteAvailabilityById controller",
		zap.String("journey", "deleteAvailabilityById"))

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

	id := c.Param("availabilityId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "deleteAvailabilityById"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := ac.service.DeleteAvailability(c, idInt); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Availability deleted successfully",
		zap.String("journey", "deleteAvailabilityById"), zap.Int("availabilityID", idInt))

	c.JSON(http.StatusOK, nil)

}

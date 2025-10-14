package special_schedule_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (ssc *specialScheduleController) DeleteSpecialScheduleById(c *gin.Context) {
	logger.Info("Init DeleteSpecialSchedule controller",
		zap.String("journey", "deleteSpecialSchedule"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	id := c.Param("specialScheduleId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "deleteSpecialSchedule"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := ssc.service.DeleteSpecialSchedule(c, idInt, user.GetID()); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Special schedule deleted successfully",
		zap.String("journey", "deleteSpecialSchedule"), zap.Int("specialScheduleID", idInt))

	c.JSON(http.StatusOK, nil)

}

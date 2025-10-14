package special_schedule_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (ssc *specialScheduleController) FindAllSpecialScheduleById(c *gin.Context) {
	logger.Info("Init findAllSpecialScheduleById controller",
		zap.String("journey", "findAllSpecialScheduleById"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	response, erro := ssc.service.FindAllSpecialSchedulesByID(c, user.GetID())

	if erro != nil {
		c.JSON(erro.Code, erro)
		return
	}

	c.JSON(http.StatusOK, response)
	logger.Info("Successful findAllSpecialScheduleById controller", zap.String("journey", "findAllSpecialScheduleById"))
}

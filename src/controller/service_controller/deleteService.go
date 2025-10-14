package service_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (sc *serviceController) DeleteService(c *gin.Context) {
	logger.Info("Init DeleteService controller",
		zap.String("journey", "deleteService"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	if user.GetRole() != "admin" {
		restError := rest_err.NewUnauthorizedRequestError("Your role is not admin")
		c.JSON(restError.Code, restError)
		return
	}

	id := c.Param("serviceId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "DeleteService"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := sc.service.DeleteService(c, idInt); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("journey", "deleteService"), zap.Int("serviceID", idInt))

	c.JSON(http.StatusOK, nil)

}

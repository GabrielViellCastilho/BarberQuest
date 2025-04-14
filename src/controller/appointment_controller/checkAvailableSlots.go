package appointment_controller

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (ac *appointmentController) CheckAvailableSlots(ctx *gin.Context) {
	logger.Info("Init checkAvailableSlots controller",
		zap.String("journey", "checkAvailableSlots"))

	barberID := ctx.Param("barberId")
	serviceID := ctx.Param("serviceId")
	dateString := ctx.Param("date")

	barberIDInt, err0 := strconv.Atoi(barberID)
	if err0 != nil {
		logger.Warn("Invalid barberID format", err0, zap.String("id", barberID), zap.String("journey", "checkAvailableSlots"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	serviceIDInt, err0 := strconv.Atoi(serviceID)
	if err0 != nil {
		logger.Warn("Invalid serviceID format", err0, zap.String("id", serviceID), zap.String("journey", "checkAvailableSlots"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	response, err := ac.service.CheckAvailableSlots(ctx, barberIDInt, serviceIDInt, dateString)
	if err != nil {
		ctx.JSON(err.Code, err)
	}

	ctx.JSON(http.StatusOK, response)
	logger.Info("Successful checkAvailableSlots controller", zap.String("journey", "checkAvailableSlots"))

}

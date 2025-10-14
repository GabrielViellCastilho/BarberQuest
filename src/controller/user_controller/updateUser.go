package user_controller

import (
	"net/http"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/validation"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/request"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "updateUser"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	var userRequest request.UserRequestUpdate

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user_domain info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := user_domain.NewUserDomainUpdate(
		user.GetID(),
		userRequest.Name,
		userRequest.Cellphone)

	if err := uc.service.UpdateUser(c, domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("journey", "updateUser"), zap.Int("userID", user.GetID()))

	c.JSON(http.StatusOK, nil)
}

func (uc *userControllerInterface) UpdatePasswordUser(c *gin.Context) {
	logger.Info("Init UpdatePasswordUser controller",
		zap.String("journey", "updatePasswordUser"))

	user, err0 := user_domain.ValidatePasswordResetToken(c.Param("token"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}

	var userRequest request.UserRequestUpdatePassword

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate password info", err)
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	domain := user_domain.NewUserDomainUpdatePassword(user.GetEmail(), userRequest.Password)

	if err := uc.service.UpdatePasswordUser(c, domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User password updated successfully",
		zap.String("journey", "updatePasswordUser"), zap.Int("userID", user.GetID()))

	c.JSON(http.StatusOK, nil)
}

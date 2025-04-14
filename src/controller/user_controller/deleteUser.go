package user_controller

import (
	"fmt"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "deleteUser"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	if user.GetRole() != "admin" {
		restErr := rest_err.NewBadRequestError("You are not admin")
		c.JSON(restErr.Code, restErr)
		return
	}

	id := c.Param("userId")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warn("Invalid id format", err, zap.String("id", id), zap.String("journey", "DeleteUser"))
		errorMessage := rest_err.NewBadRequestError("Invalid id format. ID must be a number.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := uc.service.DeleteUser(c, idInt); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("journey", "deleteUser"), zap.Int("userID", idInt))

	c.JSON(http.StatusOK, nil)

}

func (uc *userControllerInterface) DeleteMyUser(c *gin.Context) {
	logger.Info("Init DeleteMyUser controller",
		zap.String("journey", "deleteMyUser"))

	user, err0 := user_domain.VerifyToken(c.Request.Header.Get("Authorization"))
	if err0 != nil {
		c.JSON(err0.Code, err0)
		return
	}
	logger.Info(fmt.Sprintf("Delete user_domain by id: %v", user))

	if err := uc.service.DeleteUser(c, user.GetID()); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("journey", "deleteMyUser"), zap.Int("userID", user.GetID()))

	c.JSON(http.StatusOK, nil)

}

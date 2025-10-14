package user_controller

import (
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service/user_service"
	"github.com/gin-gonic/gin"
)

type userControllerInterface struct {
	service user_service.UserDomainService
}

func NewUserController(service user_service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	CreateCustomerUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	DeleteMyUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindAllBarbers(c *gin.Context)
	FindAllUsersByDateOfBirth(c *gin.Context)
	FindMyUserData(c *gin.Context)
	UpdateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdatePasswordUser(c *gin.Context)
	SendEmailResetPassword(c *gin.Context)
}

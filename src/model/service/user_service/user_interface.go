package user_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/user"
	user2 "github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
)

func NewUserDomainService(ur user.UserRepository) *userDomainService {
	return &userDomainService{
		userRepository: ur,
	}
}

type userDomainService struct {
	userRepository user.UserRepository
}

type UserDomainService interface {
	CreateUser(ctx context.Context, userDomain user2.UserDomainInterface) (*response.User_Response, *rest_err.RestErr)
	UpdateUser(ctx context.Context, userDomain user2.UserDomainInterface) *rest_err.RestErr
	FindUserByEmail(ctx context.Context, email string) (*response.User_Response, *rest_err.RestErr)
	FindUserById(ctx context.Context, id int) (*response.User_Response, *rest_err.RestErr)
	FindAllBarbers(ctx context.Context) ([]*response.User_Response, *rest_err.RestErr)
	FindAllUsersByDateOfBirth(ctx context.Context, dateOfBirth string) ([]*response.User_Response, *rest_err.RestErr)
	DeleteUser(ctx context.Context, id int) *rest_err.RestErr
	UpdatePasswordUser(ctx context.Context, userDomain user2.UserDomainInterface) *rest_err.RestErr
	SendEmailResetPassword(email, token string) *rest_err.RestErr

	FindUserByEmailAndPassword(ctx context.Context, userDomainLogin user2.UserDomainInterface) (*response.User_Response, string, *rest_err.RestErr)
}

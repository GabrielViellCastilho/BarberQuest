package user_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/view"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(ctx context.Context, userDomain user_domain.UserDomainInterface) (*response.User_Response, *rest_err.RestErr) {
	logger.Info("Init createUser service",
		zap.String("journey", "createuser"))

	if user, _ := ud.FindUserByEmail(ctx, userDomain.GetEmail()); user != nil {
		return nil, rest_err.NewConflictError("User already exists")
	}

	userDomain.EncryptPassword()

	result, err := ud.userRepository.CreateUser(ctx, userDomain)
	if err != nil {
		logger.Error("Cannot create user_domain", err, zap.String("journey", "createuser"))
		return nil, rest_err.NewInternalServerError("Cannot create user_domain")
	}

	logger.Info("Successful createUser service",
		zap.String("journey", "createuser"))

	return view.ConvertUserDomainToResponse(result), nil
}

package user_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailAndPassword(ctx context.Context, userDomainLogin user_domain.UserDomainInterface) (*response.User_Response, string, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword service",
		zap.String("journey", "findUserByEmailAndPassword"))

	userDomainLogin.EncryptPassword()

	user, err := ud.userRepository.FindUserByEmailAndPassword(ctx, userDomainLogin.GetEmail(), userDomainLogin.GetPassword())
	if err != nil {
		return nil, "", err
	}

	userResponse := view.ConvertUserDomainToResponse(user)

	jwt, erro := user.GenerateToken()
	if erro != nil {
		return nil, "", erro
	}

	logger.Info("Successful findUserByEmailAndPassword service", zap.String("journey", "findUserByEmailAndPassword"))

	return userResponse, jwt, nil
}

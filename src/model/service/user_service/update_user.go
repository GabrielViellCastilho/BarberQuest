package user_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(ctx context.Context, userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser service",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUserById(ctx, userDomain)
	if err != nil {
		return err
	}

	logger.Info("Successful updateUser service",
		zap.String("journey", "updateUser"))

	return nil
}

func (ud *userDomainService) UpdatePasswordUser(ctx context.Context, userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updatePasswordUse service",
		zap.String("journey", "updatePasswordUse"))

	userDomain.EncryptPassword()

	err := ud.userRepository.UpdateUserPasswordByEmail(ctx, userDomain)
	if err != nil {
		return err
	}

	logger.Info("Successful updatePasswordUser service",
		zap.String("journey", "updatePasswordUser"))

	return nil
}

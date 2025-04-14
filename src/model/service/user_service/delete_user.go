package user_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(ctx context.Context, id int) *rest_err.RestErr {
	logger.Info("Init deleteUser service",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}

	logger.Info("Successful deleteUser service",
		zap.String("journey", "deleteUser"))

	return nil
}

package availability_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ads *availabilityDomainService) DeleteAvailability(ctx context.Context, id int) *rest_err.RestErr {
	logger.Info("Init deleteAvailability service",
		zap.String("journey", "deleteAvailability"))

	err := ads.availabilityServiceRepository.DeleteAvailabilityById(ctx, id)
	if err != nil {
		return err
	}

	logger.Info("Successful deleteAvailability service",
		zap.String("journey", "deleteAvailability"))

	return nil
}

package barber_service_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (sds *serviceDomainService) DeleteService(ctx context.Context, id int) *rest_err.RestErr {
	logger.Info("Init deleteService service",
		zap.String("journey", "deleteService"))

	err := sds.barberServiceRepository.DeleteServiceById(ctx, id)
	if err != nil {
		return err
	}

	logger.Info("Successful deleteService service",
		zap.String("journey", "deleteService"))

	return nil
}

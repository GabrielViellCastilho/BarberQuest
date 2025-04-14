package barber_service_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service_domain"
	"go.uber.org/zap"
)

func (uds *serviceDomainService) UpdateService(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateService service",
		zap.String("journey", "updateService"))

	err := uds.barberServiceRepository.UpdateServiceById(ctx, serviceDomain)
	if err != nil {
		return err
	}

	logger.Info("Successful updateService service",
		zap.String("journey", "updateService"))

	return nil
}

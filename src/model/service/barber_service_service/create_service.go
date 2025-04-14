package barber_service_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/view"
	"go.uber.org/zap"
)

func (bss *serviceDomainService) CreateService(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) (*response.Service_Response, *rest_err.RestErr) {
	logger.Info("Init createService service",
		zap.String("journey", "createService"))

	service, err := bss.barberServiceRepository.CreateService(ctx, serviceDomain)
	if err != nil {
		return nil, err
	}

	logger.Info("Successful createUser service",
		zap.String("journey", "createuser"))

	return view.ConvertServiceDomainToResponse(service), nil
}

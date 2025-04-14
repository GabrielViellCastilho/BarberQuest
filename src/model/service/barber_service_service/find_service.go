package barber_service_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/view"
	"go.uber.org/zap"
)

func (sds *serviceDomainService) FindServiceByID(ctx context.Context, serviceID int) (*response.Service_Response, *rest_err.RestErr) {
	logger.Info("Init finServiceByID service",
		zap.String("journey", "findServiceByID"))

	service, err := sds.barberServiceRepository.FindServiceByID(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	serviceResponse := view.ConvertServiceDomainToResponse(service)

	logger.Info("Successful findServiceByID service", zap.String("journey", "findServiceByID"))

	return serviceResponse, nil
}

func (sds *serviceDomainService) FindServiceByName(ctx context.Context, serviceName string) (*response.Service_Response, *rest_err.RestErr) {
	logger.Info("Init finServiceByName service",
		zap.String("journey", "findServiceByName"))

	service, err := sds.barberServiceRepository.FindServiceByName(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	serviceResponse := view.ConvertServiceDomainToResponse(service)

	logger.Info("Successful findServiceByName service", zap.String("journey", "findServiceByName"))

	return serviceResponse, nil
}

func (sds *serviceDomainService) FindAllServices(ctx context.Context) ([]*response.Service_Response, *rest_err.RestErr) {
	logger.Info("Init findAllServices service",
		zap.String("journey", "findAllServices"))

	services, err := sds.barberServiceRepository.FindAllServices(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*response.Service_Response

	for _, service := range services {
		responses = append(responses, view.ConvertServiceDomainToResponse(service))
	}

	logger.Info("Successful findAllServices service", zap.String("journey", "findAllServices"))

	return responses, nil
}

func (sds *serviceDomainService) FindAllAvailableServices(ctx context.Context) ([]*response.Service_Response, *rest_err.RestErr) {
	logger.Info("Init findAllAvailableServices service",
		zap.String("journey", "findAllAvailableServices"))

	services, err := sds.barberServiceRepository.FindAllAvailableServices(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*response.Service_Response

	for _, service := range services {
		responses = append(responses, view.ConvertServiceDomainToResponse(service))
	}

	logger.Info("Successful findAllAvailableServices service", zap.String("journey", "findAllAvailableServices"))

	return responses, nil
}

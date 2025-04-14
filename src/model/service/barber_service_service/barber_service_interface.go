package barber_service_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/repository/baber_service"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service_domain"
)

type serviceDomainService struct {
	barberServiceRepository baber_service.ServiceRepository
}

func NewServiceDomainService(bsr baber_service.ServiceRepository) *serviceDomainService {
	return &serviceDomainService{
		barberServiceRepository: bsr,
	}
}

type BarberServiceService interface {
	CreateService(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) (*response.Service_Response, *rest_err.RestErr)
	FindServiceByID(ctx context.Context, serviceID int) (*response.Service_Response, *rest_err.RestErr)
	FindServiceByName(ctx context.Context, serviceName string) (*response.Service_Response, *rest_err.RestErr)
	FindAllServices(ctx context.Context) ([]*response.Service_Response, *rest_err.RestErr)
	FindAllAvailableServices(ctx context.Context) ([]*response.Service_Response, *rest_err.RestErr)
	UpdateService(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) *rest_err.RestErr
	DeleteService(ctx context.Context, id int) *rest_err.RestErr
}

package baber_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service_domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type baberServiceRepository struct {
	databaseConection *pgxpool.Pool
}

func NewServiceRepository(databaseConection *pgxpool.Pool) *baberServiceRepository {
	return &baberServiceRepository{databaseConection}
}

type ServiceRepository interface {
	CreateService(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) (service_domain.ServiceDomainInterface, *rest_err.RestErr)
	FindServiceByID(ctx context.Context, serviceID int) (service_domain.ServiceDomainInterface, *rest_err.RestErr)
	FindServiceByName(ctx context.Context, serviceName string) (service_domain.ServiceDomainInterface, *rest_err.RestErr)
	FindAllServices(ctx context.Context) ([]service_domain.ServiceDomainInterface, *rest_err.RestErr)
	FindAllAvailableServices(ctx context.Context) ([]service_domain.ServiceDomainInterface, *rest_err.RestErr)
	UpdateServiceById(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) *rest_err.RestErr
	DeleteServiceById(ctx context.Context, id int) *rest_err.RestErr
}

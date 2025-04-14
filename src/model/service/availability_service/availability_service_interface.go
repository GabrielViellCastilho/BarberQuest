package availability_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/availability_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/repository/barber_availability"
)

type availabilityDomainService struct {
	availabilityServiceRepository barber_availability.AvailabilityRepository
}

func NewAvailabilityDomainService(bar barber_availability.AvailabilityRepository) *availabilityDomainService {
	return &availabilityDomainService{
		availabilityServiceRepository: bar,
	}
}

type AvailabilityServiceService interface {
	CreateAvailability(ctx context.Context, availavilityDomain availability_domain.AvailabilityDomainInterface) (*response.AvailabilityResponse, *rest_err.RestErr)
	FindAllAvailabilityByID(ctx context.Context, barberID int) ([]*response.AvailabilityResponse, *rest_err.RestErr)
	UpdateAvailabilityByIdAndIdBarber(ctx context.Context, availabilityDomain availability_domain.AvailabilityDomainInterface) *rest_err.RestErr
	DeleteAvailability(ctx context.Context, id int) *rest_err.RestErr
}

package barber_availability

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/availability_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type baberAvailabilityRepository struct {
	databaseConection *pgxpool.Pool
}

func NewAvailabilityRepository(databaseConection *pgxpool.Pool) *baberAvailabilityRepository {
	return &baberAvailabilityRepository{databaseConection}
}

type AvailabilityRepository interface {
	FindUserByID(ctx context.Context, id int) (user_domain.UserDomainInterface, *rest_err.RestErr)
	CreateAvailability(ctx context.Context, availabilityDomain availability_domain.AvailabilityDomainInterface) (availability_domain.AvailabilityDomainInterface, *rest_err.RestErr)
	FindAllAvailabilityByID(ctx context.Context, barberID int) ([]availability_domain.AvailabilityDomainInterface, *rest_err.RestErr)
	UpdateAvailabilityByIdAndIdBarber(ctx context.Context, availabilityDomain availability_domain.AvailabilityDomainInterface) *rest_err.RestErr
	DeleteAvailabilityById(ctx context.Context, id int) *rest_err.RestErr
}

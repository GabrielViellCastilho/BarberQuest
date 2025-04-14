package appointment

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/appointment_domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type appointmentRepository struct {
	databaseConection *pgxpool.Pool
}

func NewAppointmentRepository(databaseConection *pgxpool.Pool) *appointmentRepository {
	return &appointmentRepository{databaseConection}
}

type AppointmentRepository interface {
	CreateAppointment(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) (appointment_domain.AppointmentDomainInterface, *rest_err.RestErr)
	FindAllAppointmentsByDateAndBarberID(ctx context.Context, barberID int, date time.Time) ([]appointment_domain.AppointmentDomainInterface, *rest_err.RestErr)
	FindAllAppointmentsByUserID(ctx context.Context, userID int) ([]appointment_domain.AppointmentDomainInterface, *rest_err.RestErr)
	FindAllHistoricAppointmentsByUserID(ctx context.Context, userID int) ([]appointment_domain.AppointmentDomainInterface, *rest_err.RestErr)
	UpdateAppointmentCompletedById(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) *rest_err.RestErr
	CheckAvailableSlots(ctx context.Context, barberID int, serviceID int, appointmentDate time.Time) (*WorkingPeriod, *BreakPeriod, *[]Appointment, *time.Duration, *rest_err.RestErr)
	DeleteAppointmentByIdAndUserID(ctx context.Context, id int, userId int) *rest_err.RestErr
	DeleteAppointmentByIdAndBarberID(ctx context.Context, id int, barberId int) *rest_err.RestErr
}

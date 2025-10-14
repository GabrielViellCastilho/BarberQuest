package appointment_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/appointment_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/repository/appointment"
	"time"
)

type appointmentDomainService struct {
	repository appointment.AppointmentRepository
}

func NewAppointmentDomainService(ar appointment.AppointmentRepository) *appointmentDomainService {
	return &appointmentDomainService{ar}
}

type AppointmentService interface {
	CreateAppointment(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) (*response.AppointmentResponse, *rest_err.RestErr)
	FindAllAppointmentsByDateAndBarberID(ctx context.Context, barberID int, date time.Time) ([]*response.AppointmentResponse, *rest_err.RestErr)
	FindAllAppointmentsByUserID(ctx context.Context, userID int) ([]*response.UserAppointmentResponse, *rest_err.RestErr)
	FindAllHistoricAppointmentsByUserID(ctx context.Context, userID int) ([]*response.UserAppointmentResponse, *rest_err.RestErr)
	UpdateAppointmentCompletedById(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) *rest_err.RestErr
	CheckAvailableSlots(ctx context.Context, barberID int, serviceID int, appointmentDate string) ([]response.TimeSlot, *rest_err.RestErr)
	DeleteAppointmentByIdAndUserID(ctx context.Context, id int, userId int) *rest_err.RestErr
	DeleteAppointmentByIdAndBarberID(ctx context.Context, id int, barberId int) *rest_err.RestErr
}

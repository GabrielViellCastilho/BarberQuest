package view

import (
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/appointment_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/availability_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/service_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/special_schedule_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
)

func ConvertUserDomainToResponse(domainInterface user_domain.UserDomainInterface) *response.User_Response {
	return &response.User_Response{
		ID:        domainInterface.GetID(),
		Email:     domainInterface.GetEmail(),
		Name:      domainInterface.GetName(),
		Cellphone: domainInterface.GetCellphone(),
		Role:      domainInterface.GetRole(),
	}
}

func ConvertServiceDomainToResponse(domainInterface service_domain.ServiceDomainInterface) *response.Service_Response {
	return &response.Service_Response{
		ID:              domainInterface.GetID(),
		Name:            domainInterface.GetName(),
		Price:           domainInterface.GetPrice(),
		DurationMinutes: domainInterface.GetDurationMinutes(),
		Available:       domainInterface.GetAvailable(),
	}
}

func ConvertAvailabilityDomainToResponse(domainInterface availability_domain.AvailabilityDomainInterface) *response.AvailabilityResponse {
	return &response.AvailabilityResponse{
		ID:             domainInterface.GetID(),
		BarberID:       domainInterface.GetBarberId(),
		DayOfWeek:      domainInterface.GetDayOfWeek(),
		StartTime:      domainInterface.GetStartTime(),
		EndTime:        domainInterface.GetEndTime(),
		BreakStartTime: domainInterface.GetBreakStartTime(),
		BreakEndTime:   domainInterface.GetBreakEndTime(),
	}
}

func ConvertSpecialScheduleDomainToResponse(domainInterface special_schedule_domain.SpecialScheduleInterface) *response.SpecialScheduleResponse {
	return &response.SpecialScheduleResponse{
		ID:             domainInterface.GetID(),
		BarberID:       domainInterface.GetBarberID(),
		Date:           domainInterface.GetDate(),
		OpeningTime:    domainInterface.GetOpeningTime(),
		ClosedTime:     domainInterface.GetClosedTime(),
		BreakStartTime: domainInterface.GetBreakStartTime(),
		BreakEndTime:   domainInterface.GetBreakEndTime(),
	}
}

func ConvertAppointmentDomainToResponse(domainInterface appointment_domain.AppointmentDomainInterface) *response.AppointmentResponse {
	return &response.AppointmentResponse{
		ID:            domainInterface.GetId(),
		BarberID:      domainInterface.GetBarberId(),
		ServiceID:     domainInterface.GetServiceId(),
		ClientName:    domainInterface.GetClientName(),
		ClientContact: domainInterface.GetClientContact(),
		Date:          domainInterface.GetAppointmentDate(),
		UserID:        domainInterface.GetUserId(),
		Completed:     domainInterface.GetCompleted(),
	}
}

func ConvertAppointmentDomainToUserAppointmentResponse(domainInterface appointment_domain.AppointmentDomainInterface) *response.UserAppointmentResponse {
	return &response.UserAppointmentResponse{
		ID:        domainInterface.GetId(),
		BarberID:  domainInterface.GetBarberId(),
		ServiceID: domainInterface.GetServiceId(),
		Date:      domainInterface.GetAppointmentDate(),
		UserID:    domainInterface.GetUserId(),
	}
}

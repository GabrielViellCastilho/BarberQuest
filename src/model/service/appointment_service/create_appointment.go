package appointment_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/appointment_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
)

func (ads *appointmentDomainService) CreateAppointment(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) (*response.AppointmentResponse, *rest_err.RestErr) {
	logger.Info("Init createAppointment service",
		zap.String("journey", "createAppointment"))

	result, err := ads.repository.CreateAppointment(ctx, appointmentDomain)
	if err != nil {
		return nil, err
	}

	logger.Info("Successful createAppointment service",
		zap.String("journey", "creatAppointment"))

	return view.ConvertAppointmentDomainToResponse(result), nil
}

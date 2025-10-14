package appointment_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
	"time"
)

func (ads *appointmentDomainService) FindAllAppointmentsByDateAndBarberID(ctx context.Context, barberID int, date time.Time) ([]*response.AppointmentResponse, *rest_err.RestErr) {
	logger.Info("Init findAllAppointmentsByDateAndBarberID service",
		zap.String("journey", "findAllAppointmentsByDateAndBarberID"))

	appointments, err := ads.repository.FindAllAppointmentsByDateAndBarberID(ctx, barberID, date)
	if err != nil {
		return nil, err
	}

	var responses []*response.AppointmentResponse

	for _, appointment := range appointments {
		responses = append(responses, view.ConvertAppointmentDomainToResponse(appointment))
	}

	logger.Info("Successful findAllAppointmentsByDateAndBarberID service", zap.String("journey", "findAllAppointmentsByDateAndBarberID"))

	return responses, nil
}

func (ads *appointmentDomainService) FindAllAppointmentsByUserID(ctx context.Context, userID int) ([]*response.UserAppointmentResponse, *rest_err.RestErr) {
	logger.Info("Init findAllAppointmentsByUserID service",
		zap.String("journey", "findAllAppointmentsByUserID"))

	appointments, err := ads.repository.FindAllAppointmentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []*response.UserAppointmentResponse

	for _, appointment := range appointments {
		responses = append(responses, view.ConvertAppointmentDomainToUserAppointmentResponse(appointment))
	}

	logger.Info("Successful findAllAppointmentsByUserID service", zap.String("journey", "findAllAppointmentsByUserID"))

	return responses, nil
}

func (ads *appointmentDomainService) FindAllHistoricAppointmentsByUserID(ctx context.Context, userID int) ([]*response.UserAppointmentResponse, *rest_err.RestErr) {
	logger.Info("Init findAllHistoricAppointmentsByUserID service",
		zap.String("journey", "findAllHistoricAppointmentsByUserID"))

	appointments, err := ads.repository.FindAllHistoricAppointmentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responses []*response.UserAppointmentResponse

	for _, appointment := range appointments {
		responses = append(responses, view.ConvertAppointmentDomainToUserAppointmentResponse(appointment))
	}

	logger.Info("Successful findAllHistoricAppointmentsByUserID service", zap.String("journey", "findAllHistoricAppointmentsByUserID"))

	return responses, nil
}

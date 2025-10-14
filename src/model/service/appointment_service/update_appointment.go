package appointment_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/appointment_domain"
	"go.uber.org/zap"
)

func (ads *appointmentDomainService) UpdateAppointmentCompletedById(ctx context.Context, appointmentDomain appointment_domain.AppointmentDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateAppointmentCompletedById service",
		zap.String("journey", "updateAppointmentCompletedById"))

	err := ads.repository.UpdateAppointmentCompletedById(ctx, appointmentDomain)
	if err != nil {
		return err
	}

	logger.Info("Successful updateAppointmentCompletedById service",
		zap.String("journey", "updateAppointmentCompletedById"))

	return nil
}

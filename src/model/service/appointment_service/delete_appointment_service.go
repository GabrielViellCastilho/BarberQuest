package appointment_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ads *appointmentDomainService) DeleteAppointmentByIdAndUserID(ctx context.Context, id int, userId int) *rest_err.RestErr {
	logger.Info("Init deleteAppointmentByIdAndUserID service",
		zap.String("journey", "deleteAppointmentByIdAndUserID"))

	err := ads.repository.DeleteAppointmentByIdAndUserID(ctx, id, userId)
	if err != nil {
		return err
	}

	logger.Info("Successful deleteAppointmentByIdAndUserID service",
		zap.String("journey", "deleteAppointmentByIdAndUserID"))

	return nil
}

func (ads *appointmentDomainService) DeleteAppointmentByIdAndBarberID(ctx context.Context, id int, barberId int) *rest_err.RestErr {
	logger.Info("Init deleteAppointmentByIdAndBarberID service",
		zap.String("journey", "deleteAppointmentByIdAndBarberID"))

	err := ads.repository.DeleteAppointmentByIdAndBarberID(ctx, id, barberId)
	if err != nil {
		return err
	}

	logger.Info("Successful deleteAppointmentByIdAndBarberID service",
		zap.String("journey", "deleteAppointmentByIdAndBarberID"))

	return nil
}

package availability_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/availability_domain"
	"go.uber.org/zap"
)

func (adas *availabilityDomainService) UpdateAvailabilityByIdAndIdBarber(ctx context.Context, availabilityDomain availability_domain.AvailabilityDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateAvailabilityByIdAndIdBarber service",
		zap.String("journey", "updateAvailabilityByIdAndIdBarber"))

	err := adas.availabilityServiceRepository.UpdateAvailabilityByIdAndIdBarber(ctx, availabilityDomain)
	if err != nil {
		return err
	}

	logger.Info("Successful updateAvailabilityByIdAndIdBarber service",
		zap.String("journey", "updateAvailabilityByIdAndIdBarber"))

	return nil
}

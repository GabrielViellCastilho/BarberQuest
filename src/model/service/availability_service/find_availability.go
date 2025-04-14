package availability_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/view"
	"go.uber.org/zap"
)

func (ads *availabilityDomainService) FindAllAvailabilityByID(ctx context.Context, barberID int) ([]*response.AvailabilityResponse, *rest_err.RestErr) {
	logger.Info("Init findAllAvailabilityByID service",
		zap.String("journey", "findAllAvailabilityByID"))

	availabilities, err := ads.availabilityServiceRepository.FindAllAvailabilityByID(ctx, barberID)
	if err != nil {
		return nil, err
	}

	var responses []*response.AvailabilityResponse

	for _, availability := range availabilities {
		responses = append(responses, view.ConvertAvailabilityDomainToResponse(availability))
	}

	logger.Info("Successful findAllAvailabilityByID service", zap.String("journey", "findAllAvailabilityByID"))

	return responses, nil
}

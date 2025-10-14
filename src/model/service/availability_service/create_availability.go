package availability_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/availability_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
)

func (ads *availabilityDomainService) CreateAvailability(ctx context.Context, availavilityDomain availability_domain.AvailabilityDomainInterface) (*response.AvailabilityResponse, *rest_err.RestErr) {
	logger.Info("Init createAvailability service",
		zap.String("journey", "createAvailability"))

	userDomain, err := ads.availabilityServiceRepository.FindUserByID(ctx, availavilityDomain.GetBarberId())
	if err != nil {
		logger.Error("Error fetching user", err, zap.String("journey", "createAvailability"))
		return nil, rest_err.NewNotFoundError("User not found")
	}

	if userDomain.GetRole() != "barber" {
		logger.Error("Cannot create availability_domain", err, zap.String("journey", "createAvailability"))
		return nil, rest_err.NewBadRequestError("Cannot create availability_domain without role barber")
	}

	result, err := ads.availabilityServiceRepository.CreateAvailability(ctx, availavilityDomain)
	if err != nil {
		logger.Error("Cannot create availability_domain", err, zap.String("journey", "createAvailability"))
		return nil, rest_err.NewInternalServerError("Cannot create availability_domain")
	}

	logger.Info("Successful createAvailability service",
		zap.String("journey", "createAvailability"))

	return view.ConvertAvailabilityDomainToResponse(result), nil
}

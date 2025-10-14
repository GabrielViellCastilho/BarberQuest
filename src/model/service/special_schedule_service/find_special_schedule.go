package special_schedule_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
)

func (ssds *specialScheduleDomainService) FindAllSpecialSchedulesByID(ctx context.Context, barberID int) ([]*response.SpecialScheduleResponse, *rest_err.RestErr) {
	logger.Info("Init findAllSpecialSchedulesByID service",
		zap.String("journey", "findAllSpecialSchedulesByID"))

	specialSchedules, err := ssds.specialScheduleRepository.FindAllSpecialScheduleByID(ctx, barberID)
	if err != nil {
		return nil, err
	}

	var responses []*response.SpecialScheduleResponse

	for _, specialSchedule := range specialSchedules {
		responses = append(responses, view.ConvertSpecialScheduleDomainToResponse(specialSchedule))
	}

	logger.Info("Successful findAllSpecialSchedulesByID service", zap.String("journey", "findAllSpecialSchedulesByID"))

	return responses, nil
}

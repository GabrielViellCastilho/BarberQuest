package special_schedule_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/special_schedule_domain"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/view"
	"go.uber.org/zap"
)

func (ssds *specialScheduleDomainService) CreateSpecialSchedule(ctx context.Context, specialScheduleDomain special_schedule_domain.SpecialScheduleInterface) (*response.SpecialScheduleResponse, *rest_err.RestErr) {
	logger.Info("Init specialSchedule service",
		zap.String("journey", "specialSchedule"))

	specialSchedule, err := ssds.specialScheduleRepository.CreateSpecialSchedule(ctx, specialScheduleDomain)
	if err != nil {
		return nil, err
	}

	logger.Info("Successful specialSchedule service",
		zap.String("journey", "specialSchedule"))

	return view.ConvertSpecialScheduleDomainToResponse(specialSchedule), nil
}

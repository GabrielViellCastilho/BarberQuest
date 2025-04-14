package special_schedule_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ssds *specialScheduleDomainService) DeleteSpecialSchedule(ctx context.Context, id int, barber_id int) *rest_err.RestErr {
	logger.Info("Init deleteSpecialSchedule service",
		zap.String("journey", "deleteSpecialSchedule"))

	err := ssds.specialScheduleRepository.DeleteSpecialScheduleById(ctx, id, barber_id)
	if err != nil {
		return err
	}

	logger.Info("Successful deleteSpecialSchedule service",
		zap.String("journey", "deleteSpecialSchedule"))

	return nil
}

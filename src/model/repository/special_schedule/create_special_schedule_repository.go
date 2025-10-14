package special_schedule

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/special_schedule_domain"
	"go.uber.org/zap"
)

func (ssr *specialScheduleRepository) CreateSpecialSchedule(ctx context.Context, specialScheduleDomain special_schedule_domain.SpecialScheduleInterface) (special_schedule_domain.SpecialScheduleInterface, *rest_err.RestErr) {
	logger.Info("Init specialScheduleDomain repository",
		zap.String("journey", "specialScheduleDomain"))
	query := `
    INSERT INTO special_schedule (barber_id, date, opening_time, closing_time, break_start_time, break_end_time)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id;
`
	var id int

	err := ssr.databaseConection.QueryRow(ctx, query, specialScheduleDomain.GetBarberID(), specialScheduleDomain.GetDate(), specialScheduleDomain.GetOpeningTime(), specialScheduleDomain.GetClosedTime(), specialScheduleDomain.GetBreakStartTime(), specialScheduleDomain.GetBreakEndTime()).Scan(&id)
	if err != nil {
		logger.Error("Error create specialScheduleDomain", err, zap.String("journey", "specialScheduleDomain"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	specialScheduleDomain.SetID(id)

	logger.Info("Successful specialScheduleDomain repository", zap.String("journey", "specialScheduleDomain"),
		zap.String("journey", "specialScheduleDomain"))

	return specialScheduleDomain, nil
}

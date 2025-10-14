package special_schedule

import (
	"context"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/special_schedule_domain"
	"go.uber.org/zap"
)

func (ssr *specialScheduleRepository) FindAllSpecialScheduleByID(ctx context.Context, barberID int) ([]special_schedule_domain.SpecialScheduleInterface, *rest_err.RestErr) {
	logger.Info("Init findAllSpecialScheduleByID repository",
		zap.String("journey", "findAllSpecialScheduleByID"))

	var specialSchedules []struct {
		ID             int
		BarberID       int
		Date           string
		OpeningTime    string
		ClosingTime    string
		BreakStartTime string
		BreakEndTime   string
	}

	query := `
SELECT id, barber_id, TO_CHAR(date, 'YYYY-MM-DD') AS date, opening_time, closing_time,break_start_time,break_end_time
FROM special_schedule
WHERE barber_id = $1
ORDER BY date ASC;
`

	rows, err := ssr.databaseConection.Query(ctx, query, barberID)
	if err != nil {
		logger.Error("Error finding special_schedule_domain", err, zap.String("journey", "findAllSpecialScheduleByID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var specialSchedule struct {
			ID             int
			BarberID       int
			Date           string
			OpeningTime    string
			ClosingTime    string
			BreakStartTime string
			BreakEndTime   string
		}

		if err := rows.Scan(&specialSchedule.ID, &specialSchedule.BarberID, &specialSchedule.Date, &specialSchedule.OpeningTime, &specialSchedule.ClosingTime, &specialSchedule.BreakStartTime, &specialSchedule.BreakEndTime); err != nil {
			logger.Error("Error scanning special schedule data", err, zap.String("journey", "findAllSpecialScheduleByID"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		specialSchedules = append(specialSchedules, specialSchedule)
	}
	
	if len(specialSchedules) == 0 {
		return nil, rest_err.NewNotFoundError("No special schedules found")
	}

	var specialSchedulesDomain []special_schedule_domain.SpecialScheduleInterface

	for _, specialSchedule := range specialSchedules {
		ss := special_schedule_domain.NewSpecialScheduleDomain(barberID, specialSchedule.Date, specialSchedule.OpeningTime, specialSchedule.ClosingTime, specialSchedule.BreakStartTime, specialSchedule.BreakEndTime)
		ss.SetID(specialSchedule.ID)
		specialSchedulesDomain = append(specialSchedulesDomain, ss)
	}

	logger.Info("Successful findAllSpecialScheduleByID repository", zap.String("journey", "findAllSpecialScheduleByID"),
		zap.String("journey", "findAllSpecialScheduleByID"))

	return specialSchedulesDomain, nil
}

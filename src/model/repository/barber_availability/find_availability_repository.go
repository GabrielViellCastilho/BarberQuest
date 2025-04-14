package barber_availability

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/availability_domain"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (bar *baberAvailabilityRepository) FindAllAvailabilityByID(ctx context.Context, barberID int) ([]availability_domain.AvailabilityDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllAvailabilityByID repository",
		zap.String("journey", "findAllAvailabilityByID"))

	var availabilities []struct {
		ID             int
		BarberID       int
		DayOfWeek      int
		StartTime      string
		EndTime        string
		BreakStartTime string
		BreakEndTime   string
	}

	query := `
SELECT id,barber_id,day_of_week, start_time, end_time,break_start_time,break_end_time
FROM barber_availability
WHERE barber_id = $1
ORDER BY day_of_week ASC;
`

	rows, err := bar.databaseConection.Query(ctx, query, barberID)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Info("No availability found",
				zap.String("journey", "findAllAvailabilityByID"))
			return nil, rest_err.NewNotFoundError("No availability found")
		}
		logger.Error("Error finding availability_domain", err, zap.String("journey", "findAllAvailabilityByID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	// Iterando sobre as linhas retornadas e populando a lista
	for rows.Next() {
		var availability struct {
			ID             int
			BarberID       int
			DayOfWeek      int
			StartTime      string
			EndTime        string
			BreakStartTime string
			BreakEndTime   string
		}

		if err := rows.Scan(&availability.ID, &availability.BarberID, &availability.DayOfWeek, &availability.StartTime, &availability.EndTime, &availability.BreakStartTime, &availability.BreakEndTime); err != nil {
			logger.Error("Error scanning availability data", err, zap.String("journey", "findAllAvailabilityByID"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		availabilities = append(availabilities, availability)
	}

	var availabilitiesDomain []availability_domain.AvailabilityDomainInterface

	for _, availability := range availabilities {
		aval := availability_domain.NewAvailabilityDomain(availability.BarberID, availability.DayOfWeek, availability.StartTime, availability.EndTime, availability.BreakStartTime, availability.BreakEndTime)
		aval.SetID(availability.ID)
		availabilitiesDomain = append(availabilitiesDomain, aval)
	}

	logger.Info("Successful findAllAvailabilityByID repository", zap.String("journey", "findAllAvailabilityByID"),
		zap.String("journey", "findAllAvailabilityByID"))

	return availabilitiesDomain, nil
}

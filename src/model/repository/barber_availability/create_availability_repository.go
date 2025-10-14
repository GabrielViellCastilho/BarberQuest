package barber_availability

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/availability_domain"
	"go.uber.org/zap"
)

func (bar *baberAvailabilityRepository) CreateAvailability(ctx context.Context, availabilityDomain availability_domain.AvailabilityDomainInterface) (availability_domain.AvailabilityDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createAvailability repository",
		zap.String("journey", "createAvailability"))

	query := `
	        INSERT INTO barber_availability (barber_id, day_of_week, start_time, end_time, break_start_time,break_end_time)
        VALUES ($1, $2, $3,$4,$5,$6)
        RETURNING id;
`
	var id int

	err := bar.databaseConection.QueryRow(ctx, query, availabilityDomain.GetBarberId(), availabilityDomain.GetDayOfWeek(), availabilityDomain.GetStartTime(), availabilityDomain.GetEndTime(), availabilityDomain.GetBreakStartTime(), availabilityDomain.GetBreakEndTime()).Scan(&id)
	if err != nil {
		logger.Error("Error create createAvailability", err, zap.String("journey", "createAvailability"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	availabilityDomain.SetID(id)

	logger.Info("Successful createAvailability repository", zap.String("journey", "createAvailability"),
		zap.String("journey", "createAvailability"))

	return availabilityDomain, nil
}

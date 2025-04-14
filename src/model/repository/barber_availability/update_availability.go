package barber_availability

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/availability_domain"
	"go.uber.org/zap"
)

func (bar baberAvailabilityRepository) UpdateAvailabilityByIdAndIdBarber(ctx context.Context, availabilityDomain availability_domain.AvailabilityDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateAvailabilityByIdAndIdBarber repository",
		zap.String("journey", "UpdateAvailabilityByIdAndIdBarber"))

	query := `
	UPDATE barber_availability
	SET start_time = $1, end_time = $2
	WHERE id = $3 AND barber_id = $4;
	`

	result, err := bar.databaseConection.Exec(ctx, query, availabilityDomain.GetStartTime(), availabilityDomain.GetEndTime(), availabilityDomain.GetID(), availabilityDomain.GetBarberId())
	if err != nil {
		logger.Error("Error updating availability_domain", err, zap.String("journey", "updateAvailabilityByIdAndIdBarber"))
		return rest_err.NewInternalServerError("Database error when updating availability_domain")
	}

	noRow := result.RowsAffected()
	if noRow == 0 {
		logger.WarnWithoutError("No availability_domain found with given ID", zap.Int("id", availabilityDomain.GetID()), zap.Int("barber_id", availabilityDomain.GetBarberId()), zap.String("journey", "updateAvailabilityByIdAndIdBarber"))
		return rest_err.NewNotFoundError("Availability not found")
	}

	logger.Info("Successful updateAvailabilityByIdAndIdBarber repository",
		zap.Int64("rowsAffected", noRow),
		zap.String("journey", "updateAvailabilityByIdAndIdBarber"))

	return nil
}

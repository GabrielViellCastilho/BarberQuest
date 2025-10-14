package baber_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service_domain"
	"go.uber.org/zap"
)

func (bsr baberServiceRepository) UpdateServiceById(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateServiceById repository",
		zap.String("journey", "updateServiceById"))

	query := `
	UPDATE services
	SET name = $1, price = $2, duration_minutes = $3,available = $4
	WHERE id = $5;
	`

	result, err := bsr.databaseConection.Exec(ctx, query, serviceDomain.GetName(), serviceDomain.GetPrice(), serviceDomain.GetDurationMinutes(), serviceDomain.GetAvailable(), serviceDomain.GetID())
	if err != nil {
		logger.Error("Error updating service_domain", err, zap.String("journey", "updateServiceById"))
		return rest_err.NewInternalServerError("Database error when updating service_domain")
	}

	noRow := result.RowsAffected()
	if noRow == 0 {
		logger.WarnWithoutError("No service_domain found with given ID", zap.Int("id", serviceDomain.GetID()), zap.String("journey", "updateUserById"))
		return rest_err.NewNotFoundError("Service not found")
	}

	logger.Info("Successful updateServiceById repository",
		zap.Int64("rowsAffected", noRow),
		zap.String("journey", "updateServiceById"))

	return nil
}

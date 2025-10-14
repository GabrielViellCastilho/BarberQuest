package barber_availability

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (bar *baberAvailabilityRepository) DeleteAvailabilityById(ctx context.Context, id int) *rest_err.RestErr {
	logger.Info("Init DeleteAvailabilityById repository", zap.String("journey", "DeleteAvailabilityById"))

	tx, err := bar.databaseConection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error("Error starting transaction", err)
		return rest_err.NewInternalServerError("Error starting transaction")
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `DELETE FROM barber_availability WHERE id = $1;`
	result, err := tx.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Error executing delete", err)
		return rest_err.NewInternalServerError("Database error during availability_domain deletion")
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected > 1 {
		logger.Error("Multiple rows affected during delete", pgx.ErrTooManyRows, zap.Int64("rowsAffected", rowsAffected))
		_ = tx.Rollback(ctx)
		return rest_err.NewInternalServerError("Multiple rows affected. Rolling back transaction.")
	} else if rowsAffected == 0 {
		logger.WarnWithoutError("No availability_domain found with given ID", zap.Int("id", id))
		return rest_err.NewNotFoundError("Availability not found")
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("Error committing transaction", err)
		return rest_err.NewInternalServerError("Error committing transaction")
	}

	logger.Info("Successful DeleteAvailabilityById repository", zap.Int("id", id))
	return nil
}

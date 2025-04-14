package user

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUserById(ctx context.Context, id int) *rest_err.RestErr {
	logger.Info("Init DeleteUserById repository", zap.String("journey", "DeleteUserById"))

	tx, err := ur.databaseConection.BeginTx(ctx, pgx.TxOptions{})
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

	query := `DELETE FROM users WHERE id = $1;`
	result, err := tx.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Error executing delete", err)
		return rest_err.NewInternalServerError("Database error during user_domain deletion")
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected > 1 {
		logger.Error("Multiple rows affected during delete", pgx.ErrTooManyRows, zap.Int64("rowsAffected", rowsAffected))
		_ = tx.Rollback(ctx)
		return rest_err.NewInternalServerError("Multiple rows affected. Rolling back transaction.")
	} else if rowsAffected == 0 {
		logger.WarnWithoutError("No user_domain found with given ID", zap.Int("id", id))
		return rest_err.NewNotFoundError("User not found")
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("Error committing transaction", err)
		return rest_err.NewInternalServerError("Error committing transaction")
	}

	logger.Info("Successful DeleteUserById repository", zap.Int("id", id))
	return nil
}

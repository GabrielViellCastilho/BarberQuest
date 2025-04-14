package user

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUserById(ctx context.Context, userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUserById repository",
		zap.String("journey", "updateUserById"))

	query := `
	UPDATE users
	SET name = $1, cellphone=$2
	WHERE id = $3;
	`

	result, err := ur.databaseConection.Exec(ctx, query, userDomain.GetName(), userDomain.GetCellphone(), userDomain.GetID())
	if err != nil {
		logger.Error("Error updating user_domain", err, zap.String("journey", "updateUserById"))
		return rest_err.NewInternalServerError("Database error when updating user_domain")
	}

	noRow := result.RowsAffected()
	if noRow == 0 {
		logger.WarnWithoutError("No user_domain found with given ID", zap.Int("id", userDomain.GetID()), zap.String("journey", "updateUserById"))
		return rest_err.NewNotFoundError("User not found")
	}

	logger.Info("Successful updateUserById repository",
		zap.Int64("rowsAffected", noRow),
		zap.String("journey", "updateUserById"))

	return nil
}

func (ur *userRepository) UpdateUserPasswordByEmail(ctx context.Context, userDomain user_domain.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUserPasswordByEmail repository",
		zap.String("journey", "updateUserPasswordByEmail"))

	query := `
	UPDATE users
	SET password = $1
	WHERE email = $2;
	`

	result, err := ur.databaseConection.Exec(ctx, query, userDomain.GetPassword(), userDomain.GetEmail())
	if err != nil {
		logger.Error("Error updating user_domain", err, zap.String("journey", "updateUserPasswordByEmail"))
		return rest_err.NewInternalServerError("Database error when updating user_domain")
	}

	noRow := result.RowsAffected()
	if noRow == 0 {
		logger.WarnWithoutError("No user_domain found with given ID", zap.Int("id", userDomain.GetID()), zap.String("journey", "updateUserPasswordById"))
		return rest_err.NewNotFoundError("User not found")
	}

	logger.Info("Successful updateUserPasswordByEmail repository",
		zap.Int64("rowsAffected", noRow),
		zap.String("journey", "updateUserByEmail"))

	return nil
}

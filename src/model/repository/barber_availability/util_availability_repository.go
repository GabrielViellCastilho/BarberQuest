package barber_availability

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/model/user_domain"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (bar *baberAvailabilityRepository) FindUserByID(ctx context.Context, id int) (user_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))

	var user struct {
		ID        int
		Email     string
		Name      string
		Role      string
		Cellphone string
	}

	query := `
    SELECT id, email, name, role, cellphone
    FROM users 
    WHERE id = $1;
`
	err := bar.databaseConection.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Name, &user.Role, &user.Cellphone)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("user_domain not found")
		}
		logger.Error("Error finding user_domain", err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	domainUser := user_domain.NewUserDomain(user.Email, "", user.Name, user.Role, user.Cellphone)

	domainUser.SetID(user.ID)

	logger.Info("Successful findUserByID repository", zap.String("journey", "findUserByID"),
		zap.String("journey", "findUserByID"))

	return domainUser, nil
}

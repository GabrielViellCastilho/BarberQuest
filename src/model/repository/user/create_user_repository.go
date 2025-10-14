package user

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
)

func (ur *userRepository) CreateUser(ctx context.Context, userDomain user_domain.UserDomainInterface) (user_domain.UserDomainInterface, *rest_err.RestErr) {
	query := `
        INSERT INTO users (name,email,password,role,cellphone,date_of_birth)
        VALUES ($1, $2, $3, $4,$5,$6)
        RETURNING id;
    `

	var id int

	if userDomain.GetDateOfBirth() == "" {
		err := ur.databaseConection.QueryRow(ctx, query,
			userDomain.GetName(),
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetRole(),
			userDomain.GetCellphone(),
			nil,
		).Scan(&id)

		if err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	} else {
		err := ur.databaseConection.QueryRow(ctx, query,
			userDomain.GetName(),
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetRole(),
			userDomain.GetCellphone(),
			userDomain.GetDateOfBirth(),
		).Scan(&id)

		if err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
	}

	userDomain.SetID(id)

	return userDomain, nil
}

package user

import (
	"context"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(ctx context.Context, email string) (user_domain.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))

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
    WHERE email = $1;
`
	err := ur.databaseConection.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Name, &user.Role, &user.Cellphone)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("user_domain not found")
		}
		logger.Error("Error finding user_domain", err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	domainUser := user_domain.NewUserDomain(user.Email, "", user.Name, user.Role, user.Cellphone)

	domainUser.SetID(user.ID)

	logger.Info("Successful findUserByEmail repository", zap.String("journey", "findUserByEmail"),
		zap.String("journey", "findUserByEmail"))

	return domainUser, nil
}

func (ur *userRepository) FindUserByID(ctx context.Context, id int) (user_domain.UserDomainInterface, *rest_err.RestErr) {
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
	err := ur.databaseConection.QueryRow(ctx, query, id).Scan(
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

func (ur *userRepository) FindUserByEmailAndPassword(ctx context.Context, email string, password string) (user_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"))

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
    WHERE email = $1 AND password = $2;
`
	err := ur.databaseConection.QueryRow(ctx, query, email, password).Scan(
		&user.ID, &user.Email, &user.Name, &user.Role, &user.Cellphone)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("email or password incorrect")
		}
		logger.Error("Error finding user_domain", err, zap.String("journey", "findUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	domainUser := user_domain.NewUserDomain(user.Email, "", user.Name, user.Role, user.Cellphone)

	domainUser.SetID(user.ID)

	logger.Info("Successful findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("journey", "findUserByEmailAndPassword"))

	return domainUser, nil
}

func (ur *userRepository) FindAllBarbers(ctx context.Context) ([]user_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllBarbers repository",
		zap.String("journey", "findAllBarbers"))

	var users []struct {
		ID   int
		Name string
	}

	query := `
    SELECT id, name
    FROM users 
    WHERE role='barber';
`

	rows, err := ur.databaseConection.Query(ctx, query)
	if err != nil {
		logger.Error("Error finding user_domain", err, zap.String("journey", "findAllBarbers"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user struct {
			ID   int
			Name string
		}

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			logger.Error("Error scanning user_domain data", err, zap.String("journey", "findAllBarbers"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, rest_err.NewNotFoundError("No barbers found")
	}

	var usersDomain []user_domain.UserDomainInterface

	for _, user := range users {
		userDomain := user_domain.NewUserDomain("", "", user.Name, "", "")
		userDomain.SetID(user.ID)
		usersDomain = append(usersDomain, userDomain)
	}

	logger.Info("Successful findAllBarbers repository", zap.String("journey", "findAllBarbers"),
		zap.String("journey", "findAllBarbers"))

	return usersDomain, nil
}

func (ur *userRepository) FindAllUsersByDateOfBirth(ctx context.Context, dateOfBirth string) ([]user_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllUsersByDateOfBirth repository",
		zap.String("journey", "findAllUsersByDateOfBirth"))

	var users []struct {
		ID        int
		Name      string
		Cellphone string
	}

	query := `
    SELECT id, name, cellphone
    FROM users 
    WHERE TO_CHAR(date_of_birth, 'MM-DD') = $1;
`

	rows, err := ur.databaseConection.Query(ctx, query, dateOfBirth)
	if err != nil {
		logger.Error("Error finding user_domain", err, zap.String("journey", "findAllUsersByDateOfBirth"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user struct {
			ID        int
			Name      string
			Cellphone string
		}

		if err := rows.Scan(&user.ID, &user.Name, &user.Cellphone); err != nil {
			logger.Error("Error scanning user_domain data", err, zap.String("journey", "findAllUsersByDateOfBirth"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, rest_err.NewNotFoundError("No users found")
	}

	var usersDomain []user_domain.UserDomainInterface

	for _, user := range users {
		userDomain := user_domain.NewUserDomain("", "", user.Name, "", user.Cellphone)
		userDomain.SetID(user.ID)
		usersDomain = append(usersDomain, userDomain)
	}

	logger.Info("Successful findAllUsersByDateOfBirth repository", zap.String("journey", "findAllUsersByDateOfBirth"),
		zap.String("journey", "findAllUsersByDateOfBirth"))

	return usersDomain, nil
}

func (ur *userRepository) FindUsersByRole(role string) ([]user_domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUsersByRole repository",
		zap.String("journey", "findUsersByRole"))

	query := `
		SELECT id, email, name, role, cellphone
		FROM users
		WHERE role = $1;
	`

	ctx := context.Background()

	rows, err := ur.databaseConection.Query(ctx, query, role)
	if err != nil {
		logger.Error("Error querying users by role", err, zap.String("journey", "findUsersByRole"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var users []user_domain.UserDomainInterface

	for rows.Next() {
		var u struct {
			ID        int
			Email     string
			Name      string
			Role      string
			Cellphone string
		}

		if err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.Role, &u.Cellphone); err != nil {
			logger.Error("Error scanning user row", err, zap.String("journey", "findUsersByRole"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		domainUser := user_domain.NewUserDomain(u.Email, "", u.Name, u.Role, u.Cellphone)
		domainUser.SetID(u.ID)

		users = append(users, domainUser)
	}

	if len(users) == 0 {
		return nil, rest_err.NewNotFoundError("no users found with the specified role")
	}

	logger.Info("Successful findUsersByRole repository",
		zap.String("journey", "findUsersByRole"),
		zap.Int("usersFound", len(users)))

	return users, nil
}

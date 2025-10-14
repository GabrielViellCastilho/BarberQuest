package user

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	databaseConection *pgxpool.Pool
}

func NewUserRepository(databaseConection *pgxpool.Pool) *userRepository {
	return &userRepository{databaseConection}
}

type UserRepository interface {
	CreateUser(ctx context.Context, userDomain user_domain.UserDomainInterface) (user_domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(ctx context.Context, email string) (user_domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(ctx context.Context, id int) (user_domain.UserDomainInterface, *rest_err.RestErr)
	FindAllBarbers(ctx context.Context) ([]user_domain.UserDomainInterface, *rest_err.RestErr)
	FindAllUsersByDateOfBirth(ctx context.Context, dateOfBirth string) ([]user_domain.UserDomainInterface, *rest_err.RestErr)
	UpdateUserById(ctx context.Context, userDomain user_domain.UserDomainInterface) *rest_err.RestErr
	DeleteUserById(ctx context.Context, id int) *rest_err.RestErr
	UpdateUserPasswordByEmail(ctx context.Context, userDomain user_domain.UserDomainInterface) *rest_err.RestErr

	FindUserByEmailAndPassword(ctx context.Context, email string, password string) (user_domain.UserDomainInterface, *rest_err.RestErr)
}

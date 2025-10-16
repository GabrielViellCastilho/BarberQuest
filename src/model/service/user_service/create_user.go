package user_service

import (
	"context"
	"net/http"
	"os"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/user_domain"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
)

const (
	ADMIN_EMAIL    = "ADMIN_EMAIL"
	ADMIN_PASSWORD = "ADMIN_PASSWORD"
	ADMIN_NAME     = "ADMIN_NAME"
	ADMIN_PHONE    = "ADMIN_PHONE"
)

func (ud *userDomainService) CreateUser(ctx context.Context, userDomain user_domain.UserDomainInterface) (*response.User_Response, *rest_err.RestErr) {
	logger.Info("Init createUser service",
		zap.String("journey", "createuser"))

	if user, _ := ud.FindUserByEmail(ctx, userDomain.GetEmail()); user != nil {
		return nil, rest_err.NewConflictError("User already exists")
	}

	userDomain.EncryptPassword()

	result, err := ud.userRepository.CreateUser(ctx, userDomain)
	if err != nil {
		logger.Error("Cannot create user_domain", err, zap.String("journey", "createuser"))
		return nil, rest_err.NewInternalServerError("Cannot create user_domain")
	}

	logger.Info("Successful createUser service",
		zap.String("journey", "createuser"))

	return view.ConvertUserDomainToResponse(result), nil
}

func (ud *userDomainService) CreateAdminIfNotExists() *rest_err.RestErr {
	logger.Info("Init createAdminIfNotExists service",
		zap.String("journey", "createAdminIfNotExists"))

	admins, err := ud.userRepository.FindUsersByRole("admin")
	if err != nil {
		if err.Code == http.StatusNotFound {
			logger.Info("No admin found, creating default admin",
				zap.String("journey", "createAdminIfNotExists"))
		} else {
			logger.Error("Error checking for existing admin", err,
				zap.String("journey", "createAdminIfNotExists"))
			return rest_err.NewInternalServerError("Cannot check for existing admin")
		}
	}

	if len(admins) > 0 {
		logger.Info("Admin already exists, skipping creation",
			zap.String("journey", "createAdminIfNotExists"))
		return nil
	}

	adminDomain := user_domain.NewUserDomain(
		os.Getenv(ADMIN_EMAIL),
		os.Getenv(ADMIN_PASSWORD),
		os.Getenv(ADMIN_NAME),
		"admin",
		os.Getenv(ADMIN_PHONE),
	)

	adminDomain.EncryptPassword()

	_, createErr := ud.userRepository.CreateUser(context.Background(), adminDomain)
	if createErr != nil {
		logger.Error("Cannot create admin user", createErr,
			zap.String("journey", "createAdminIfNotExists"))
		return rest_err.NewInternalServerError("Cannot create admin user")
	}

	logger.Info("Admin user created successfully",
		zap.String("journey", "createAdminIfNotExists"))

	return nil
}

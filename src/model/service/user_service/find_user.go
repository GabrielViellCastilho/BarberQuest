package user_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"github.com/GabrielViellCastilho/BarberQuest/src/view"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmail(ctx context.Context, email string) (*response.User_Response, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service",
		zap.String("journey", "findUserByEmail"))

	user, err := ud.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	userResponse := view.ConvertUserDomainToResponse(user)

	logger.Info("Successful findUserByEmail service", zap.String("journey", "findUserByEmail"))

	return userResponse, nil

}

func (ud *userDomainService) FindUserById(ctx context.Context, id int) (*response.User_Response, *rest_err.RestErr) {
	logger.Info("Init findUserByID service",
		zap.String("journey", "findUserByID"))

	user, err := ud.userRepository.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userResponse := view.ConvertUserDomainToResponse(user)

	logger.Info("Successful findUserByID service", zap.String("journey", "findUserByID"))

	return userResponse, nil

}

func (ud *userDomainService) FindAllBarbers(ctx context.Context) ([]*response.User_Response, *rest_err.RestErr) {
	logger.Info("Init findAllBarbers service",
		zap.String("journey", "findAllBarbers"))

	users, err := ud.userRepository.FindAllBarbers(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*response.User_Response

	for _, user := range users {
		responses = append(responses, view.ConvertUserDomainToResponse(user))
	}

	logger.Info("Successful findAllBarbers service", zap.String("journey", "findAllBarbers"))

	return responses, nil
}

func (ud *userDomainService) FindAllUsersByDateOfBirth(ctx context.Context, dateOfBirth string) ([]*response.User_Response, *rest_err.RestErr) {
	logger.Info("Init findAllUsersByDateOfBirth service",
		zap.String("journey", "findAllUsersByDateOfBirth"))

	users, err := ud.userRepository.FindAllUsersByDateOfBirth(ctx, dateOfBirth)
	if err != nil {
		return nil, err
	}

	var responses []*response.User_Response

	for _, user := range users {
		responses = append(responses, view.ConvertUserDomainToResponse(user))
	}

	logger.Info("Successful findAllUsersByDateOfBirth service", zap.String("journey", "findAllUsersByDateOfBirth"))

	return responses, nil
}

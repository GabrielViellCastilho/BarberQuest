package baber_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service_domain"
	"go.uber.org/zap"
)

func (bsr *baberServiceRepository) CreateService(ctx context.Context, serviceDomain service_domain.ServiceDomainInterface) (service_domain.ServiceDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createBarberService repository",
		zap.String("journey", "createBarberService"))
	query := `
	        INSERT INTO services (name,price,duration_minutes,available)
        VALUES ($1, $2, $3,$4)
        RETURNING id;
`
	var id int

	err := bsr.databaseConection.QueryRow(ctx, query, serviceDomain.GetName(), serviceDomain.GetPrice(), serviceDomain.GetDurationMinutes(), serviceDomain.GetAvailable()).Scan(&id)
	if err != nil {
		logger.Error("Error create barber_service", err, zap.String("journey", "createBarberService"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	serviceDomain.SetID(id)

	logger.Info("Successful createBarberService repository", zap.String("journey", "createBarberService"),
		zap.String("journey", "createBarberService"))

	return serviceDomain, nil
}

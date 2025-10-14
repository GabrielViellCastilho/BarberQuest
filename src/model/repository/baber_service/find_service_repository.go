package baber_service

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/service_domain"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (bsr *baberServiceRepository) FindServiceByID(ctx context.Context, serviceID int) (service_domain.ServiceDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findServiceByID repository",
		zap.String("journey", "findServiceByID"))

	var service struct {
		ID              int
		Name            string
		Price           float32
		DurationMinutes int
		Available       bool
	}

	query := `
	SELECT  name,price,duration_minutes,available
	FROM services 
	WHERE id = $1;
`
	err := bsr.databaseConection.QueryRow(ctx, query, serviceID).Scan(&service.Name, &service.Price, &service.DurationMinutes, &service.Available)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("service_domain not found")
		}
		logger.Error("Error finding service_domain", err, zap.String("journey", "serviceServiceByID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	serviceDomain := service_domain.NewServiceDomain(service.Name, service.Price, service.DurationMinutes, service.Available)
	serviceDomain.SetID(serviceID)

	logger.Info("Successful findServiceByID repository", zap.String("journey", "findServiceByID"),
		zap.String("journey", "findServiceByID"))

	return serviceDomain, nil
}

func (bsr *baberServiceRepository) FindServiceByName(ctx context.Context, serviceName string) (service_domain.ServiceDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findServiceByName repository",
		zap.String("journey", "findServiceByName"))

	var service struct {
		ID              int
		Name            string
		Price           float32
		DurationMinutes int
		Available       bool
	}

	query := `
	SELECT  id,name,price,available
	FROM services 
	WHERE name = $1;
`
	err := bsr.databaseConection.QueryRow(ctx, query, serviceName).Scan(&service.ID, &service.Name, &service.Price, &service.DurationMinutes, &service.Available)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("service_domain not found")
		}
		logger.Error("Error finding service_domain", err, zap.String("journey", "serviceServiceByName"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	serviceDomain := service_domain.NewServiceDomain(service.Name, service.Price, service.DurationMinutes, service.Available)
	serviceDomain.SetID(service.ID)

	logger.Info("Successful findServiceByName repository", zap.String("journey", "findServiceBynAME"),
		zap.String("journey", "findServiceByName"))

	return serviceDomain, nil
}

func (bsr *baberServiceRepository) FindAllServices(ctx context.Context) ([]service_domain.ServiceDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllServices repository",
		zap.String("journey", "findAllServices"))

	var services []struct {
		ID              int
		Name            string
		Price           float32
		DurationMinutes int
		Available       bool
	}

	query := `
	SELECT  id,name,price,duration_minutes,available
	FROM services;
`

	rows, err := bsr.databaseConection.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("No services found")
		}
		logger.Error("Error finding service_domain", err, zap.String("journey", "findAllServices"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var service struct {
			ID              int
			Name            string
			Price           float32
			DurationMinutes int
			Available       bool
		}

		if err := rows.Scan(&service.ID, &service.Name, &service.Price, &service.DurationMinutes, &service.Available); err != nil {
			logger.Error("Error scanning service data", err, zap.String("journey", "findAllServices"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		services = append(services, service)
	}

	var servicesDomain []service_domain.ServiceDomainInterface

	for _, service := range services {
		serviceDomain := service_domain.NewServiceDomain(service.Name, service.Price, service.DurationMinutes, service.Available)
		serviceDomain.SetID(service.ID)
		servicesDomain = append(servicesDomain, serviceDomain)
	}

	logger.Info("Successful findAllServices repository", zap.String("journey", "findAllServices"),
		zap.String("journey", "findAllServices"))

	return servicesDomain, nil
}

func (bsr *baberServiceRepository) FindAllAvailableServices(ctx context.Context) ([]service_domain.ServiceDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllAvailableServices repository",
		zap.String("journey", "findAllAvailableServices"))

	var services []struct {
		ID              int
		Name            string
		Price           float32
		DurationMinutes int
		Available       bool
	}

	query := `
	SELECT  id,name,price,duration_minutes,available
	FROM services
	WHERE available = true;
`

	rows, err := bsr.databaseConection.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, rest_err.NewNotFoundError("No services found")
		}
		logger.Error("Error finding service_domain", err, zap.String("journey", "findAllAvailableServices"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var service struct {
			ID              int
			Name            string
			Price           float32
			DurationMinutes int
			Available       bool
		}

		if err := rows.Scan(&service.ID, &service.Name, &service.Price, &service.DurationMinutes, &service.Available); err != nil {
			logger.Error("Error scanning service data", err, zap.String("journey", "findAllAvailableServices"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		services = append(services, service)
	}

	var servicesDomain []service_domain.ServiceDomainInterface

	for _, service := range services {
		serviceDomain := service_domain.NewServiceDomain(service.Name, service.Price, service.DurationMinutes, service.Available)
		serviceDomain.SetID(service.ID)
		servicesDomain = append(servicesDomain, serviceDomain)
	}

	logger.Info("Successful findAllAvailableServices repository", zap.String("journey", "findAllAvailableServices"),
		zap.String("journey", "findAllAvailableServices"))

	return servicesDomain, nil
}

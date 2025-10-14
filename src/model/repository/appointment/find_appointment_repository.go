package appointment

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/model/appointment_domain"
	"go.uber.org/zap"
)

func (ar *appointmentRepository) FindAllAppointmentsByDateAndBarberID(ctx context.Context, barberID int, date time.Time) ([]appointment_domain.AppointmentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllAppointmentsByDateAndBarberID repository",
		zap.String("journey", "findAllAppointmentsByDateAndBarberID"))

	var appointments []struct {
		ID              int
		UserID          sql.NullInt64
		ClientName      string
		ClientContact   string
		AppointmentDate time.Time
		BarberID        int
		ServiceId       int
		Completed       *bool
	}

	query := `
SELECT id, client_name, client_contact, appointment_date, barber_id, service_id,user_id,completed
FROM appointments
WHERE DATE(appointment_date) = $1 AND barber_id = $2
ORDER BY appointment_date::time;
`
	fmt.Println(date.Format("2006-01-02"))
	rows, err := ar.databaseConection.Query(ctx, query, date.Format("2006-01-02"), barberID)
	if err != nil {
		logger.Error("Error finding appointment_domain", err, zap.String("journey", "findAllAppointmentsByDateAndBarberID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var appointment struct {
			ID              int
			UserID          sql.NullInt64
			ClientName      string
			ClientContact   string
			AppointmentDate time.Time
			BarberID        int
			ServiceId       int
			Completed       *bool
		}

		if err := rows.Scan(&appointment.ID, &appointment.ClientName, &appointment.ClientContact, &appointment.AppointmentDate, &appointment.BarberID, &appointment.ServiceId, &appointment.UserID, &appointment.Completed); err != nil {
			logger.Error("Error scanning availability data", err, zap.String("journey", "findAllAppointmentsByDateAndBarberID"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		appointments = append(appointments, appointment)
	}

	if len(appointments) == 0 {
		return nil, rest_err.NewNotFoundError("No appointments found")
	}

	var appointmentsDomain []appointment_domain.AppointmentDomainInterface

	for _, appointment := range appointments {
		var userID int
		if appointment.UserID.Valid {
			userID = int(appointment.UserID.Int64)
		} else {
			userID = 0
		}

		app := appointment_domain.NewAppointmentDomain(appointment.ClientName, appointment.ClientContact, appointment.BarberID, appointment.ServiceId, appointment.AppointmentDate, userID, appointment.Completed)
		app.SetId(appointment.ID)
		appointmentsDomain = append(appointmentsDomain, app)
	}

	logger.Info("Successful findAllAppointmentsByDateAndBarberID repository", zap.String("journey", "findAllAppointmentsByDateAndBarberID"),
		zap.String("journey", "findAllAppointmentsByDateAndBarberID"))

	return appointmentsDomain, nil
}

func (ar *appointmentRepository) FindAllAppointmentsByUserID(ctx context.Context, userID int) ([]appointment_domain.AppointmentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllAppointmentsByUserID repository",
		zap.String("journey", "findAllAppointmentsByUserID"))

	var appointments []struct {
		ID              int
		UserID          int
		ClientName      string
		ClientContact   string
		AppointmentDate time.Time
		BarberID        int
		ServiceId       int
	}

	query := `
	SELECT id, client_name, client_contact, appointment_date, barber_id, service_id, user_id
	FROM appointments
	WHERE user_id = $1 
	AND appointment_date >= CURRENT_DATE
	ORDER BY appointment_date ASC ;
`
	rows, err := ar.databaseConection.Query(ctx, query, userID)
	if err != nil {
		logger.Error("Error finding appointment_domain", err, zap.String("journey", "findAllAppointmentsByUserID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var appointment struct {
			ID              int
			UserID          int
			ClientName      string
			ClientContact   string
			AppointmentDate time.Time
			BarberID        int
			ServiceId       int
		}

		if err := rows.Scan(&appointment.ID, &appointment.ClientName, &appointment.ClientContact, &appointment.AppointmentDate, &appointment.BarberID, &appointment.ServiceId, &appointment.UserID); err != nil {
			logger.Error("Error scanning availability data", err, zap.String("journey", "findAllAppointmentsByUserID"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		appointments = append(appointments, appointment)
	}

	if len(appointments) == 0 {
		return nil, rest_err.NewNotFoundError("No appointments found")
	}

	var appointmentsDomain []appointment_domain.AppointmentDomainInterface

	for _, appointment := range appointments {

		app := appointment_domain.NewAppointmentDomain(appointment.ClientName, appointment.ClientContact, appointment.BarberID, appointment.ServiceId, appointment.AppointmentDate, userID, nil)
		app.SetId(appointment.ID)
		appointmentsDomain = append(appointmentsDomain, app)
	}

	logger.Info("Successful findAllAppointmentsByUserID repository", zap.String("journey", "findAllAppointmentsByUserID"),
		zap.String("journey", "findAllAppointmentsByUserID"))

	return appointmentsDomain, nil
}

func (ar *appointmentRepository) FindAllHistoricAppointmentsByUserID(ctx context.Context, userID int) ([]appointment_domain.AppointmentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findAllHistoricAppointmentsByUserID repository",
		zap.String("journey", "findAllHistoricAppointmentsByUserID"))

	var appointments []struct {
		ID              int
		UserID          int
		ClientName      string
		ClientContact   string
		AppointmentDate time.Time
		BarberID        int
		ServiceId       int
	}

	query := `
	SELECT id, client_name, client_contact, appointment_date, barber_id, service_id, user_id
	FROM appointments
	WHERE user_id = $1 
	ORDER BY appointment_date ASC ;
`
	rows, err := ar.databaseConection.Query(ctx, query, userID)
	if err != nil {
		logger.Error("Error finding appointment_domain", err, zap.String("journey", "findAllHistoricAppointmentsByUserID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var appointment struct {
			ID              int
			UserID          int
			ClientName      string
			ClientContact   string
			AppointmentDate time.Time
			BarberID        int
			ServiceId       int
		}

		if err := rows.Scan(&appointment.ID, &appointment.ClientName, &appointment.ClientContact, &appointment.AppointmentDate, &appointment.BarberID, &appointment.ServiceId, &appointment.UserID); err != nil {
			logger.Error("Error scanning availability data", err, zap.String("journey", "findAllAppointmentsByUserID"))
			return nil, rest_err.NewInternalServerError(err.Error())
		}

		appointments = append(appointments, appointment)
	}

	if len(appointments) == 0 {
		return nil, rest_err.NewNotFoundError("No appointments found")
	}

	var appointmentsDomain []appointment_domain.AppointmentDomainInterface

	for _, appointment := range appointments {

		app := appointment_domain.NewAppointmentDomain(appointment.ClientName, appointment.ClientContact, appointment.BarberID, appointment.ServiceId, appointment.AppointmentDate, userID, nil)
		app.SetId(appointment.ID)
		appointmentsDomain = append(appointmentsDomain, app)
	}

	logger.Info("Successful findAllHistoricAppointmentsByUserID repository",
		zap.String("journey", "findAllHistoricAppointmentsByUserID"))

	return appointmentsDomain, nil
}

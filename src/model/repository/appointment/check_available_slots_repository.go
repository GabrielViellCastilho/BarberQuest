package appointment

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"time"
)

// Struct para representar o período de funcionamento
type WorkingPeriod struct {
	StartTime time.Time
	EndTime   time.Time
}

type BreakPeriod struct {
	StartTime time.Time
	EndTime   time.Time
}

// Struct para representar um agendamento
type Appointment struct {
	StartTime time.Time
	Duration  time.Duration
}

// Struct para representar um slot de tempo disponível
type TimeSlot struct {
	Slot        time.Time
	IsAvailable bool
}

func (ar *appointmentRepository) CheckAvailableSlots(ctx context.Context, barberID int, serviceID int, appointmentDate time.Time) (*WorkingPeriod, *BreakPeriod, *[]Appointment, *time.Duration, *rest_err.RestErr) {
	logger.Info("Init checkAvailableSlots repository",
		zap.String("journey", "checkAvailableSlots"))

	var workingPeriod WorkingPeriod

	var breakPeriod BreakPeriod

	// Primeiro, busca um horário especial para o barbeiro
	query := `
		SELECT opening_time, closing_time,break_start_time,break_end_time
		FROM special_schedule
		WHERE DATE(date) = $1 AND barber_id = $2;
	`
	rows, err := ar.databaseConection.Query(ctx, query, appointmentDate.Format("2006-01-02"), barberID)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Info("No check available slots",
				zap.String("journey", "checkAvailableSlots"))
			return nil, nil, nil, nil, rest_err.NewNotFoundError("No check available slots")
		}
		logger.Error("Error executing query for special_schedule", err)
		return nil, nil, nil, nil, rest_err.NewInternalServerError("Error executing query for special_schedule")
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&workingPeriod.StartTime, &workingPeriod.EndTime, &breakPeriod.StartTime, &breakPeriod.EndTime); err != nil {
			logger.Error("Error scanning special_schedule", err)
			return nil, nil, nil, nil, rest_err.NewInternalServerError("Error scanning special_schedule")
		}
	} else {
		// Caso não haja horário especial, busca a disponibilidade regular
		query = `
			SELECT start_time, end_time ,break_start_time,break_end_time
			FROM barber_availability 
			WHERE day_of_week = $1 AND barber_id = $2;
		`
		err = ar.databaseConection.QueryRow(ctx, query, int(appointmentDate.Weekday()), barberID).Scan(&workingPeriod.StartTime, &workingPeriod.EndTime, &breakPeriod.StartTime, &breakPeriod.EndTime)
		if err != nil {
			if err == pgx.ErrNoRows {
				logger.Info("No check available slots",
					zap.String("journey", "checkAvailableSlots"))
				return nil, nil, nil, nil, rest_err.NewNotFoundError("No check available slots")
			}
			logger.Error("Error finding barber availability", err)
			return nil, nil, nil, nil, rest_err.NewInternalServerError("Error finding barber availability")
		}
	}

	// Busca os horários agendados e suas durações
	query = `
		SELECT a.appointment_date::TIME AS scheduled_time, s.duration_minutes
		FROM appointments a
		JOIN services s ON a.service_id = s.id
		WHERE DATE(a.appointment_date) = $1;
	`

	rows.Close()
	rows, err = ar.databaseConection.Query(ctx, query, appointmentDate.Format("2006-01-02"))
	if err != nil {
		logger.Error("Error executing query for appointments", err)
		return nil, nil, nil, nil, rest_err.NewInternalServerError("Error executing query for appointments")
	}
	defer rows.Close()

	var appointments []Appointment
	for rows.Next() {
		var appointment Appointment
		var durationMinutes int

		if err := rows.Scan(&appointment.StartTime, &durationMinutes); err != nil {
			logger.Error("Error scanning appointment", err)
			return nil, nil, nil, nil, rest_err.NewInternalServerError("Error scanning appointment")
		}

		appointment.Duration = time.Duration(durationMinutes) * time.Minute
		appointments = append(appointments, appointment)
	}

	// Obtém a duração do serviço selecionado
	var durationMinutes int
	query = `
		SELECT duration_minutes
		FROM services
		WHERE id = $1;
	`
	err = ar.databaseConection.QueryRow(ctx, query, serviceID).Scan(&durationMinutes)
	if err != nil {
		logger.Error("Error retrieving service duration", err)
		return nil, nil, nil, nil, rest_err.NewInternalServerError("Error retrieving service duration")
	}

	serviceDuration := time.Duration(durationMinutes) * time.Minute

	logger.Info("Successful checkAvailableSlots repository", zap.String("journey", "checkAvailableSlots"))

	return &workingPeriod, &breakPeriod, &appointments, &serviceDuration, nil
}

package appointment_service

import (
	"context"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/logger"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/SpartanBarbearia/src/controller/model/response"
	"go.uber.org/zap"
	"time"
)

func (ads *appointmentDomainService) CheckAvailableSlots(ctx context.Context, barberID int, serviceID int, appointmentDate string) ([]response.TimeSlot, *rest_err.RestErr) {
	logger.Info("Init checkAvailableSlots service",
		zap.String("journey", "checkAvailableSlots"))

	// Converter a string para time.Time
	date, err := time.Parse("2006-01-02", appointmentDate)
	if err != nil {
		logger.Error("Error parsing date", err)
		return nil, rest_err.NewInternalServerError("Invalid date format")
	}

	// Obter período de funcionamento, intervalo de pausa e agendamentos do repositório
	workingPeriod, breakPeriod, appointments, consultationDuration, err0 := ads.repository.CheckAvailableSlots(ctx, barberID, serviceID, date)
	if err0 != nil {
		return nil, err0
	}

	// Verificar se os dados retornados são válidos
	if workingPeriod == nil || appointments == nil || consultationDuration == nil {
		logger.WarnWithoutError("Error: Missing data from repository")
		return nil, rest_err.NewInternalServerError("Could not retrieve available slots")
	}

	var slots []response.TimeSlot

	// Gera slots de 15 em 15 minutos respeitando a duração do serviço
	for currentTime := workingPeriod.StartTime; currentTime.Add(*consultationDuration).Before(workingPeriod.EndTime) || currentTime.Add(*consultationDuration).Equal(workingPeriod.EndTime); currentTime = currentTime.Add(15 * time.Minute) {
		isAvailable := true

		// Verificar se o slot está ocupado por algum agendamento
		for _, appointment := range *appointments {
			if Overlaps(currentTime, *consultationDuration, appointment.StartTime, appointment.Duration) {
				isAvailable = false
				break
			}
		}

		// Verificar se o slot cai dentro do intervalo de pausa
		if breakPeriod != nil && Overlaps(currentTime, *consultationDuration, breakPeriod.StartTime, breakPeriod.EndTime.Sub(breakPeriod.StartTime)) {
			isAvailable = false
		}

		// Adiciona o slot apenas se houver tempo suficiente para o serviço e não estiver no intervalo
		if isAvailable {
			slots = append(slots, response.TimeSlot{
				Slot:        currentTime,
				IsAvailable: true,
			})
		}
	}

	logger.Info("Successful checkAvailableSlots service", zap.String("journey", "checkAvailableSlots"))
	return slots, nil
}

// Função auxiliar para verificar sobreposição de horários
func Overlaps(start1 time.Time, duration1 time.Duration, start2 time.Time, duration2 time.Duration) bool {
	end1 := start1.Add(duration1)
	end2 := start2.Add(duration2)
	return start1.Before(end2) && end1.After(start2)
}

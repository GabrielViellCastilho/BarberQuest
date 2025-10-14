package appointment_service

import (
	"context"
	"time"

	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/GabrielViellCastilho/BarberQuest/src/controller/model/response"
	"go.uber.org/zap"
)

func (ads *appointmentDomainService) CheckAvailableSlots(ctx context.Context, barberID int, serviceID int, appointmentDate string) ([]response.TimeSlot, *rest_err.RestErr) {
	logger.Info("Init checkAvailableSlots service",
		zap.String("journey", "checkAvailableSlots"))

	date, err := time.Parse("2006-01-02", appointmentDate)
	if err != nil {
		logger.Error("Error parsing date", err)
		return nil, rest_err.NewInternalServerError("Invalid date format")
	}

	workingPeriod, breakPeriod, appointments, consultationDuration, err0 := ads.repository.CheckAvailableSlots(ctx, barberID, serviceID, date)
	if err0 != nil {
		return nil, err0
	}

	if workingPeriod == nil || appointments == nil || consultationDuration == nil {
		logger.WarnWithoutError("Error: Missing data from repository")
		return nil, rest_err.NewInternalServerError("Could not retrieve available slots")
	}

	var slots []response.TimeSlot

	for currentTime := workingPeriod.StartTime; currentTime.Add(*consultationDuration).Before(workingPeriod.EndTime) || currentTime.Add(*consultationDuration).Equal(workingPeriod.EndTime); currentTime = currentTime.Add(15 * time.Minute) {
		isAvailable := true

		for _, appointment := range *appointments {
			if Overlaps(currentTime, *consultationDuration, appointment.StartTime, appointment.Duration) {
				isAvailable = false
				break
			}
		}

		if breakPeriod != nil && Overlaps(currentTime, *consultationDuration, breakPeriod.StartTime, breakPeriod.EndTime.Sub(breakPeriod.StartTime)) {
			isAvailable = false
		}

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

func Overlaps(start1 time.Time, duration1 time.Duration, start2 time.Time, duration2 time.Duration) bool {
	end1 := start1.Add(duration1)
	end2 := start2.Add(duration2)
	return start1.Before(end2) && end1.After(start2)
}

package request

import "time"

type CreateAppointmentRequest struct {
	ClientName      string    `json:"client_name" binding:"max=100"`
	ClientContact   string    `json:"client_contact"`
	AppointmentDate time.Time `json:"appointment_date" binding:"required" time_format:"2006-01-02T15:04:05"`
	BarberID        int       `json:"barber_id" binding:"required,gt=0"`
	ServiceID       int       `json:"service_id" binding:"required,gt=0"`
}

type UpdateCompletedAppointmentRequest struct {
	Completed *bool `json:"completed"`
}

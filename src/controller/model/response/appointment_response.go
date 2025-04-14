package response

import "time"

type AppointmentResponse struct {
	ID            int       `json:"id"`
	BarberID      int       `json:"barber_id"`
	ServiceID     int       `json:"service_id"`
	UserID        int       `json:"user_id"`
	ClientName    string    `json:"client_name"`
	ClientContact string    `json:"client_contact"`
	Date          time.Time `json:"date"`
	Completed     *bool     `json:"completed"`
}

type UserAppointmentResponse struct {
	ID        int       `json:"id"`
	BarberID  int       `json:"barber_id"`
	ServiceID int       `json:"service_id"`
	UserID    int       `json:"user_id"`
	Date      time.Time `json:"date"`
}

type TimeSlot struct {
	Slot        time.Time `json:"slot"`
	IsAvailable bool      `json:"is_available"`
}

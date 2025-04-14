package appointment_domain

import "time"

type appointmentDomain struct {
	id               int
	client_name      string
	client_contact   string
	appointment_date time.Time
	barber_id        int
	service_id       int
	user_id          int
	completed        *bool
}

func NewAppointmentDomain(client_name, client_contact string, barber_id int, service_id int, appointmentDate time.Time, user_id int, completed *bool) *appointmentDomain {
	return &appointmentDomain{
		client_name:      client_name,
		client_contact:   client_contact,
		barber_id:        barber_id,
		service_id:       service_id,
		appointment_date: appointmentDate,
		user_id:          user_id,
		completed:        completed,
	}
}

func (ad *appointmentDomain) GetId() int {
	return ad.id
}
func (ad *appointmentDomain) GetClientName() string {
	return ad.client_name
}
func (ad *appointmentDomain) GetClientContact() string {
	return ad.client_contact
}
func (ad *appointmentDomain) GetBarberId() int {
	return ad.barber_id
}
func (ad *appointmentDomain) GetServiceId() int {
	return ad.service_id
}
func (ad *appointmentDomain) GetAppointmentDate() time.Time {
	return ad.appointment_date
}
func (ad *appointmentDomain) GetUserId() int {
	return ad.user_id
}
func (ad *appointmentDomain) GetCompleted() *bool { return ad.completed }
func (ad *appointmentDomain) SetId(id int) {
	ad.id = id
}

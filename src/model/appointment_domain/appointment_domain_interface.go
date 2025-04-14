package appointment_domain

import "time"

type AppointmentDomainInterface interface {
	GetId() int
	GetClientName() string
	GetClientContact() string
	GetBarberId() int
	GetServiceId() int
	GetAppointmentDate() time.Time
	GetUserId() int
	GetCompleted() *bool
	SetId(id int)
}

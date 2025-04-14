package service_domain

type ServiceDomainInterface interface {
	GetPrice() float32
	GetID() int
	GetName() string
	GetDurationMinutes() int
	GetAvailable() bool
	SetID(id int)
}

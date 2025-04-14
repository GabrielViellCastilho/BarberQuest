package availability_domain

type AvailabilityDomainInterface interface {
	GetID() int
	GetBarberId() int
	GetDayOfWeek() int
	GetStartTime() string
	GetEndTime() string
	GetBreakStartTime() string
	GetBreakEndTime() string
	SetID(id int)
}

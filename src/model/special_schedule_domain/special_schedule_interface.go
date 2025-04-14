package special_schedule_domain

type SpecialScheduleInterface interface {
	GetID() int
	GetBarberID() int
	GetDate() string
	GetOpeningTime() string
	GetClosedTime() string
	GetBreakStartTime() string
	GetBreakEndTime() string
	SetID(id int)
}

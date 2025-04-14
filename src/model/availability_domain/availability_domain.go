package availability_domain

type availabilityDomain struct {
	id               int
	barber_id        int
	day_of_week      int
	start_time       string
	end_time         string
	break_start_time string
	break_end_time   string
}

func NewAvailabilityDomain(barber_id int, day_of_week int, start_time string, end_time string, break_start_time string, break_end_time string) *availabilityDomain {
	return &availabilityDomain{
		barber_id:        barber_id,
		day_of_week:      day_of_week,
		start_time:       start_time,
		end_time:         end_time,
		break_start_time: break_start_time,
		break_end_time:   break_end_time,
	}
}

func (ad *availabilityDomain) GetID() int {
	return ad.id
}
func (ad *availabilityDomain) GetBarberId() int {
	return ad.barber_id
}
func (ad *availabilityDomain) GetDayOfWeek() int {
	return ad.day_of_week
}
func (ad *availabilityDomain) GetStartTime() string {
	return ad.start_time
}
func (ad *availabilityDomain) GetEndTime() string {
	return ad.end_time
}
func (ad *availabilityDomain) GetBreakStartTime() string { return ad.break_start_time }
func (ad *availabilityDomain) GetBreakEndTime() string   { return ad.break_end_time }

func (ad *availabilityDomain) SetID(id int) {
	ad.id = id
}

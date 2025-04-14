package special_schedule_domain

type specialScheduleDomain struct {
	id               int
	barber_id        int
	date             string
	opening_time     string
	closed_time      string
	break_start_time string
	break_end_time   string
}

func NewSpecialScheduleDomain(barber_id int, date, opening_time, closed_time, break_start_time, break_end_time string) *specialScheduleDomain {
	return &specialScheduleDomain{
		barber_id:        barber_id,
		date:             date,
		opening_time:     opening_time,
		closed_time:      closed_time,
		break_start_time: break_start_time,
		break_end_time:   break_end_time,
	}
}

func (d *specialScheduleDomain) GetID() int {
	return d.id
}
func (d *specialScheduleDomain) GetBarberID() int {
	return d.barber_id
}
func (d *specialScheduleDomain) GetDate() string {
	return d.date
}
func (d *specialScheduleDomain) GetOpeningTime() string {
	return d.opening_time
}
func (d *specialScheduleDomain) GetClosedTime() string {
	return d.closed_time
}
func (d *specialScheduleDomain) GetBreakStartTime() string { return d.break_start_time }
func (d *specialScheduleDomain) GetBreakEndTime() string   { return d.break_end_time }
func (d *specialScheduleDomain) SetID(id int) {
	d.id = id
}

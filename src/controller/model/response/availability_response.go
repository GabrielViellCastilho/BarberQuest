package response

type AvailabilityResponse struct {
	ID             int    `json:"id"`
	BarberID       int    `json:"barber_id"`
	DayOfWeek      int    `json:"day_of_week"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	BreakStartTime string `json:"break_start_time"`
	BreakEndTime   string `json:"break_end_time"`
}

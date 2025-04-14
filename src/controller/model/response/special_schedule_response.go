package response

type SpecialScheduleResponse struct {
	ID             int    `json:"ID"`
	BarberID       int    `json:"barberID"`
	Date           string `json:"date"`
	OpeningTime    string `json:"opening_time"`
	ClosedTime     string `json:"closed_time"`
	BreakStartTime string `json:"break_start_time"`
	BreakEndTime   string `json:"break_end_time"`
}

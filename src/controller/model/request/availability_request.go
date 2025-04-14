package request

type AvailabilityRequest struct {
	DayOfWeek      int    `json:"day_of_week" binding:"required,numeric,oneof=1 2 3 4 5 6"`
	StartTime      string `json:"start_time" binding:"required,datetime=15:04:05"`
	EndTime        string `json:"end_time" binding:"required,datetime=15:04:05"`
	BreakStartTime string `json:"break_start_time" binding:"required,datetime=15:04:05"`
	BreakEndTime   string `json:"break_end_time" binding:"required,datetime=15:04:05"`
}

type AvailabilityRequestUpdate struct {
	StartTime      string `json:"start_time" binding:"required,datetime=15:04:05"`
	EndTime        string `json:"end_time" binding:"required,datetime=15:04:05"`
	BreakStartTime string `json:"break_start_time" binding:"required,datetime=15:04:05"`
	BreakEndTime   string `json:"break_end_time" binding:"required,datetime=15:04:05"`
}

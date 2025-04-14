package request

type SpecialSchedulerRequest struct {
	Date           string `json:"date" binding:"required,datetime=2006-01-02"`
	OpeningTime    string `json:"opening_time" binding:"required,datetime=15:04:05"`
	ClosedTime     string `json:"closed_time" binding:"required,datetime=15:04:05"`
	BreakStartTime string `json:"break_start_time" binding:"required,datetime=15:04:05"`
	BreakEndTime   string `json:"break_end_time" binding:"required,datetime=15:04:05"`
}

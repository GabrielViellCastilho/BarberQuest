package response

type Service_Response struct {
	ID              int     `json:"ID"`
	Name            string  `json:"name"`
	Price           float32 `json:"price"`
	DurationMinutes int     `json:"duration_minutes"`
	Available       bool    `json:"available"`
}

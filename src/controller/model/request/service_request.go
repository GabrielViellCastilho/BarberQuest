package request

type ServiceRequest struct {
	Name            string  `json:"name" binding:"required,min=4,max=100"`
	Price           float32 `json:"price" binding:"required,numeric,gte=0"`
	DurationMinutes int     `json:"duration_minutes" binding:"required,numeric,gte=0"`
	Available       bool    `json:"available"`
}

package response

type User_Response struct {
	ID        int    `json:"ID"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Cellphone string `json:"cellphone"`
	Role      string `json:"role"`
}

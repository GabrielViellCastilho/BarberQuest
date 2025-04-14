package request

type UserRequest struct {
	Name      string `json:"name" binding:"required,min=4,max=100"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Role      string `json:"role" binding:"required,max=6,oneof=admin barber"`
	Cellphone string `json:"cellphone" binding:"required,min=11,max=11"`
}

type CustomerUserRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=100"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Cellphone   string `json:"cellphone" binding:"required,min=11,max=11"`
	DateOfBirth string `json:"dateOfBirth" binding:"required,datetime=2006-01-02"`
}

type UserRequestUpdate struct {
	Name      string `json:"name" binding:"required,min=4,max=100"`
	Cellphone string `json:"cellphone" binding:"required,min=11,max=11"`
}

type UserRequestUpdatePassword struct {
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
}

type UserRequestLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

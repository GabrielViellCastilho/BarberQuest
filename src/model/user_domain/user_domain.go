package user_domain

type userDomain struct {
	id          int
	name        string
	email       string
	password    string
	role        string
	cellphone   string
	dateOfBirth string
}

func NewUserDomain(email string, password string, name string, role string, cellphone string) *userDomain {
	return &userDomain{
		email:     email,
		password:  password,
		name:      name,
		role:      role,
		cellphone: cellphone,
	}

}

func NewUserDomainUpdate(id int, name string, cellphone string) *userDomain {
	return &userDomain{
		id:        id,
		name:      name,
		cellphone: cellphone,
	}

}

func NewUserDomainUpdatePassword(email string, password string) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
	}

}

func NewUserDomainLogin(email, password string) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetRole() string { return ud.role }
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetID() int {
	return ud.id
}
func (ud *userDomain) GetCellphone() string {
	return ud.cellphone
}
func (ud *userDomain) GetDateOfBirth() string { return ud.dateOfBirth }

func (ud *userDomain) SetDateOfBirth(dateOfBirth string) { ud.dateOfBirth = dateOfBirth }
func (ud *userDomain) SetID(id int)                      { ud.id = id }

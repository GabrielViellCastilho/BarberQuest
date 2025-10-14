package user_domain

import "github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"

type UserDomainInterface interface {
	EncryptPassword()
	GetEmail() string
	GetName() string
	GetPassword() string
	GetRole() string
	GetID() int
	GetCellphone() string
	GetDateOfBirth() string
	SetDateOfBirth(dateOfBirth string)
	SetID(id int)
	GenerateToken() (string, *rest_err.RestErr)
}

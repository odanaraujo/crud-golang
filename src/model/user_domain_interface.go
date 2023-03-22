package model

type UserDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8
	SetID(string)

	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		age:      age,
	}
}

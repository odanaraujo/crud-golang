package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int8) UserDomainInterface {
	return &userDomain{
		name, email, password, age,
	}
}

type userDomain struct {
	name     string
	email    string
	password string
	age      int8
}

func (user *userDomain) GetName() string {
	return user.name
}
func (user *userDomain) GetEmail() string {
	return user.email
}
func (user *userDomain) GetPassword() string {
	return user.password
}
func (user *userDomain) GetAge() int8 {
	return user.age
}

func (user *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.password))                 //Transformando a senha string em um array de byte
	user.password = hex.EncodeToString(hash.Sum(nil)) //troca a senha que ele vai enviar por uma ecryptada

}

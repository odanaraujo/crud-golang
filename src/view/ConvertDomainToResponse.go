package view

import (
	"github.com/odanaraujo/crud-golang/src/controller/model/response"
	"github.com/odanaraujo/crud-golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}

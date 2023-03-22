package converter

import (
	"github.com/odanaraujo/crud-golang/src/entity"
	"github.com/odanaraujo/crud-golang/src/model"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}

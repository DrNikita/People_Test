package model

import (
	"gorm.io/gorm"
)

type Person struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic"`
}

type SupplementedPerson struct {
	Person
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	CountryId string `json:"country_id"`
}

type PersonFullInfo struct {
	gorm.Model
	SupplementedPerson
}

func (*PersonFullInfo) TableName() string {
	return "persons"
}

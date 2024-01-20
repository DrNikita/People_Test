package model

import "gorm.io/gorm"

type Persons struct {
	gorm.Model
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	CountryId  string
}

package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName 	string
	MiddleName 	string
	LastName 	string
	Email		string
}

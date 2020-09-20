package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code 		string
	Sku			string
	Name 		string
	Description	string
	Price		float64
}
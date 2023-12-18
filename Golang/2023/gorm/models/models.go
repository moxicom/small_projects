package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	UserRefer uint
}

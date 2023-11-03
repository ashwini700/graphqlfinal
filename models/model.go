package models

import "gorm.io/gorm"

type NewUser struct {
	Email    string `json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	Name         string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	HashPassword string `json:"hash_password" validate:"required"`
}
type Company struct {
	gorm.Model
	Name     string `json:"companyname" validate:"required"`
	Location string `json:"companylocation" validate:"required"`
}
type Job struct {
	gorm.Model
	Company Company `json:"-" gorm:"foreignKey:cid"`
	Cid     uint    `json:"cid"`
	Role    string  `json:"role"`
	Salary  string  `json:"salary"`
}

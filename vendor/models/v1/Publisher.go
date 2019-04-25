package v1

import "github.com/jinzhu/gorm"

type Publisher struct {
	gorm.Model

	// has many
	Books 	[]Book

	// columns
	Name	string	`gorm:"type:varchar(100);not null"`
}
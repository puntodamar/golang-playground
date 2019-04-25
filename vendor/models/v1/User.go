package v1

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name        string	`gorm:"type:varchar(150);unique_index;not null"`
	Username	string  `gorm:"type:varchar(20);unique_index;not null"`
	Email       string  `gorm:"type:varchar(100);unique_index;not null"`
}

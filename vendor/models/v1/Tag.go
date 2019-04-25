package v1

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model

	// has many
	BookTags []BookTag

	// columns
	Name string	`gorm:"type:varchar(20);unique_index;not null"`
}


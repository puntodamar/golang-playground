package v1

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model

	// has many
	BookAuthors []BookAuthor

	// columns
	Name        string	`gorm:"type:varchar(150);not null"`
}


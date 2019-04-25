package v1

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model

	// has many
	BookAuthors []BookAuthor
	BookTags	[]Tag

	// columns
	Name	string	`gorm:"type:varchar(200);unique_index;not null"`
}

package v1

import "github.com/jinzhu/gorm"

type Author struct {
	// columns
	ID			int     `gorm:"primary_key" json:"id"`
	Name        string	`gorm:"type:varchar(150);not null" json:"name"`
	Books		[]Book 	`gorm:"many2many:book_authors" json:"books"`
}

type AuthorOnly struct {
	ID			int     `gorm:"primary_key" json:"id"`
	Name        string	`gorm:"type:varchar(150);not null" json:"name"`
}

func CompleteAuthorList(db *gorm.DB) *gorm.DB {

	db. Table("authors a").
		Joins("JOIN book_authors ba 	ON ba.author_id = a.id").
		Joins("JOIN books b 			ON b.id 		= ba.book_id").
		Joins("JOIN book_tags bt 		ON bt.book_id 	= b.id").
		Joins("JOIN tags t				ON t.id			= bt.tag_id").
		Select("a.*").
		Preload("Books.Tags")
	return db
}


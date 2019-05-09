package v1

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Book struct {

	// columns
	ID			int     	`gorm:"primary_key" json:"id"`
	PublisherID uint		`gorm:"type:int(11);not null" json:"publisher_id"`
	Name		string		`gorm:"type:varchar(200);unique_index;not null" json:"name"`
	PublishedAt	time.Time	`gorm:"type:date" json:"published_at"`
	CreatedAt	time.Time 	`gorm:"type:datetime;" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"type:datetime;" json:"updated_at"`
	DeletedAt	*time.Time 	`gorm:"type:datetime;" json:"deleted_at"`

	Authors		[]Author	`gorm:"many2many:book_authors" json:"authors"`
	Publisher 	Publisher	`json:"publisher"`
	Tags 		[]Tag 		`gorm:"many2many:book_tags" json:"tags"`
}

type BookOnly struct {

	// columns
	ID			int     		`gorm:"primary_key" json:"id"`
	PublisherID uint			`gorm:"type:int(11);not null" json:"-"`
	Name		string			`gorm:"type:varchar(200);unique_index;not null" json:"name"`
	PublishedAt	time.Time		`gorm:"type:date" json:"published_at"`
	CreatedAt	time.Time 		`gorm:"type:datetime;" json:"created_at"`
	UpdatedAt	time.Time 		`gorm:"type:datetime;" json:"updated_at"`
	DeletedAt	*time.Time 		`gorm:"type:datetime;" json:"deleted_at"`

	Authors		[]AuthorOnly	`json:"authors"`
	Publisher  	PublisherOnly 	`json:"publisher"`
}

func (BookOnly) TableName() string {
	return "books"
}

type BookQuery struct {

}

func BookList (db *gorm.DB) *gorm.DB {
	db.	Table("books b").
		Joins("JOIN book_authors ba 	ON ba.book_id 	= b.id").
		Joins("JOIN authors a 		ON a.id 		= ba.author_id").
		Joins("JOIN book_tags bt		ON bt.book_id	= b.id").
		Joins("JOIN tags t			ON t.id			= bt.tag_id").
		Select("b.*")

	return db
}
package book

import (
	"time"
	model "models/v1"
)

type Book struct {
	// columns
	ID			int     			`json:"id"`
	PublisherID uint				`json:"publisher_id"`
	Name		string				`json:"name"`
	PublishedAt	time.Time			`json:"published_at"`
	CreatedAt	time.Time 			`json:"created_at"`
	UpdatedAt	time.Time 			`json:"updated_at"`
	DeletedAt	*time.Time 			`json:"deleted_at"`

	Author		model.Author		`json:"author"`
	Publisher 	model.Publisher		`json:"publisher"`
	Tags 		[]model.Publisher 	`gorm:"many2many:book_tags" json:"tags"`
}

package v1

import "time"

type Book struct {

	// columns
	ID			int     	`gorm:"primary_key" json:"id"`
	Name		string		`gorm:"type:varchar(200);unique_index;not null" json:"name"`
	PublishedAt	time.Time	`gorm:"type:date" json:"published_at"`
	CreatedAt	time.Time 	`gorm:"type:datetime;" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"type:datetime;" json:"updated_at"`
	DeletedAt	*time.Time 	`gorm:"type:datetime;" json:"deleted_at"`

	Tags []Tag `gorm:"many2many:book_tags" json:"tags"`
}

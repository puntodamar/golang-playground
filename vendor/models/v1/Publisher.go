package v1

import "time"

type Publisher struct {
	// columns
	ID 			int 		`gorm:"primary_key" json:"id"`
	Name		string		`gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt	time.Time 	`gorm:"type:datetime;" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"type:datetime;" json:"updated_at"`

	Books 		[]Book
}

type PublisherOnly struct {
	ID 			int 		`json:"id"`
	Name		string		`json:"name"`
}

func (PublisherOnly) TableName() string {
	return "publishers"
}
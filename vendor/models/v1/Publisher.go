package v1

import "time"

type Publisher struct {
	// has many
	Books 	[]Book

	// columns
	ID 			int 		`gorm:"primary_key" json:"id"`
	Name		string		`gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt	time.Time 	`gorm:"type:datetime;" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"type:datetime;" json:"updated_at"`

}
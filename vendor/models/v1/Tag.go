package v1

type Tag struct {

	// has many
	BookTags []BookTag

	// columns
	ID 		int 	`gorm:"primary_key" json:"id"`
	Name 	string	`gorm:"type:varchar(20);unique_index;not null" json:"name"`
}


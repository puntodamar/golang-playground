package v1

type Author struct {
	// has many
	BookAuthors []BookAuthor

	// columns
	ID			int     `gorm:"primary_key" json:"id"`
	Name        string	`gorm:"type:varchar(150);not null" json:"name"`
}


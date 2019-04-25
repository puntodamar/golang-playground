package v1

type BookAuthor struct {

	// has one
	BookID		uint
	AuthorID 	uint

	// columns
	Name        string	`gorm:"type:varchar(200);not null"`
}


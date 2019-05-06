package v1

type BookAuthor struct {

	// has one
	BookID		uint 	`gorm:"type:int(11);not null" json:"book_id"`
	AuthorID 	uint	`gorm:"type:int(11);not null" json:"author_id"`

	Book 		Book
	Author		Author
}


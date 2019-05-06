package v1

type BookTag struct {

	// has one
	BookID 	uint `gorm:"type:int(11);not null" json:"book_id"`
	TagID 	uint `gorm:"type:int(11);not null" json:"tag_id"`

	Tag 	Tag
	Book 	Book
}


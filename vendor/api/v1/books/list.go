package books

import (
	"github.com/gin-gonic/gin"

	//"strconv"

	. "configs/database"
	//. "helpers"
	model "models/v1"
)

func List(c *gin.Context) {
	db 			:= Db
	//page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	//limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))
	//
	//books 		:= Paginate(Param{
	//	DB 		: db.Scopes(model.BookList),
	//	Page 	: page,
	//	Limit 	: limit,
	//	Model	: &model.Book{},
	//})


	books := []model.BookOnly{}
	db.Find(&books)

	for i := 0; i < len(books); i++ {

		authors := []model.AuthorOnly{}
		db. Table("authors a").
			Joins("JOIN book_authors ba ON ba.author_id = a.id").
			Where("ba.book_id = ?", books[i].ID).
			Select("a.*").Scan(&authors)

		publisher := model.PublisherOnly{}
		db.	Table("books b").
			Joins("JOIN publishers p ON p.id = b.publisher_id").
			Where("p.id = ?", books[i].PublisherID).
			Select("p.*").
			Scan(&publisher)

		books[i].Authors 	= authors
		books[i].Publisher 	= publisher
	}

	c.JSON(200,&books)
	return
	c.JSON(200, gin.H{"data" : &books})
}
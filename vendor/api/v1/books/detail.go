package books

import (
	. "configs/database"
	"github.com/gin-gonic/gin"
	//model "models/v1"
	"net/http"
)

func Detail(c *gin.Context){
	db := Db

	id 		:= c.Param("id")
	//book 	:= model.Book{}

	//result :=
	//	db.	Table("books").
	//		Joins("JOIN book_authors ba ON ba.book_id 		= b.id").
	//		Joins("JOIN authors a 		ON a.id 			= ba.author_id").
	//		Joins("JOIN publishers p 	ON b.publisher_id 	= p.id").
	//		Select("b.*, a.name, p.name").
	//		Where("b.id = ?", &id).RowsAffected


	result :=
		db.Table("books").Where("id = ?", &id)


	//if err := db.Find(&book,id).Error; err != nil {
	//	c.AbortWithStatusJSON(
	//		http.StatusUnprocessableEntity,
	//		gin.H{"message" : "can't find book"},
	//	)
	//}

	c.JSON(http.StatusOK, result)
}

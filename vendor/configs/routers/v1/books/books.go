package books

import (
	"api/v1/books"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine){
	a := route.Group("/api/v1/books")
	{
		a.GET("/", books.List)
	}
}


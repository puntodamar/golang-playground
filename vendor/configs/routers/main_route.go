package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"

	"configs/routers/v1/auth"
	"configs/routers/v1/users"
	"configs/routers/v1/authors"
	"configs/routers/v1/books"
)

func GetRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//v.RegisterValidation("bookabledate", bookableDate)
	}

	auth.Routes(r)
	users.Routes(r)
	authors.Routes(r)
	books.Routes(r)

	return r
}

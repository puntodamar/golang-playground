package routers

import (
	"configs/routers/v1/auth"
	"configs/routers/v1/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
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

	return r
}

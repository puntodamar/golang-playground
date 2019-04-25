package routers

import (
	"github.com/gin-gonic/gin"

	"configs/routers/v1/auth"
	"configs/routers/v1/users"
)

func GetRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	auth.Routes(r)
	users.Routes(r)

	return r
}

package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"net/http"

	. "api/v1/authentication/struct"

)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		if len(token) == 0{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message" : "no authorization token found",
			})
		}

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error)  {
			return JwtKey, nil
		})

		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message" : "invalid token",
			})

		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message" : "invalid token signature",
				})

			}
			c.AbortWithStatus(http.StatusBadRequest)

		}

		c.Next()
	}
}
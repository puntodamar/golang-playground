package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"net/http"

	. "api/v1/authentication/struct"
)

func Welcome(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// For any other type of error, return a bad request status
		c.AbortWithStatus(http.StatusBadRequest)
	}


	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error)  {
		return JwtKey, nil
	})
	if !tkn.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)

	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatus(http.StatusUnauthorized)

		}
		c.AbortWithStatus(http.StatusBadRequest)

	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	c.JSON(http.StatusOK, claims.Username)
}

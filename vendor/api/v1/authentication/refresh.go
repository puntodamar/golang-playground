package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Refresh(c *gin.Context) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	token, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// For any other type of error, return a bad request status
		c.AbortWithStatus(http.StatusBadRequest)
	}



	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "unauthorized access"})
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "unauthorized access"})
		}
		c.AbortWithStatus(http.StatusBadRequest)

	}
	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime		:= time.Now().Add(5 * time.Minute)
	claims.ExpiresAt	 = expirationTime.Unix()
	newToken 			:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err 	:= newToken.SignedString(jwtKey)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	// Set the new token as the users `token` cookie
	c.JSON(http.StatusOK, gin.H{
		"message" 	: "authorization success",
		"token"		: tokenString,
	})
}



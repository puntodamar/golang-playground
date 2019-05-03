package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"

	"net/http"

	"helpers/validator/requests/v1/auth"
	. "api/v1/authentication/struct"
	. "configs/database"

	global_helper 	"helpers/global"
	model 			"models/v1"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time


func Auth(c *gin.Context) {

	db := Db
	var user auth.LoginForm

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	userLogin 	:= &model.User{}
	exists 		:=
		db.Model(&model.User{}).
			Where(map[string]interface{}{"username": user.Username}).
			First(&userLogin).RowsAffected

	if exists != 1 || !global_helper.CheckPasswordHash(user.Password, userLogin.Password) {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message" : "invalid credentials",
		})
		return
	}


	//Declare the expiration time of the token
	//here, we have kept it as 5 minutes

	var expirationTime time.Time
	if user.Username == "jamal"{
		expirationTime = time.Now().Add(24 * time.Hour)
	} else {
		expirationTime = time.Now().Add(5 * time.Minute)
	}

//	Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username		: user.Username,
		StandardClaims	: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt	: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	//http.SetCookie(c.Writer, &http.Cookie{
	//	Name	: "token",
	//	Value	: tokenString,
	//	Expires	: expirationTime,
	//})

	c.JSON(http.StatusOK, gin.H{
		"message" 	: "authorization success",
		"token"		: tokenString,
	})

}
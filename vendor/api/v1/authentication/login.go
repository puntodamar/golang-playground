package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
	"net/http"

					"helpers/validator/requests/v1/auth"
	. 				"configs/database"
	global_helper 	"helpers/global"
	model 			"models/v1"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}


// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


func Login(c *gin.Context) {

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
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username		: user.Username,
		StandardClaims	: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt	: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	//c.SetCookie("token", tokenString)
	http.SetCookie(c.Writer, &http.Cookie{
		Name	: "token",
		Value	: tokenString,
		Expires	: expirationTime,
	})

	c.JSON(http.StatusOK, gin.H{
		"message" 	: "authorization success",
		"token"		: tokenString,
	})
}
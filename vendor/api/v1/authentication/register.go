package authentication

import (

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"net/http"

	"helpers/validator"
	. "configs/database"
	model 				"models/v1"
	request_validator "helpers/validator/requests/v1/auth"
)


func Register(c *gin.Context) {
	db 					:= Db
	req 				:= &request_validator.RegisterForm{}
	jsonFormatter 		:= &validator.JsonFormatter{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			jsonFormatter.FormatJsonErrorFromString(err.Error()))
	}

	hash, _ 	:= HashPassword(req.Password)
	newUser 	:= model.User{
		Name		: req.Name,
		Username	: req.Username,
		Email		: req.Email,
		Password	: hash,
	}


	var ok, err = req.Validate(db, &newUser)

	if ok {
		tx := db.Begin()

		if dbc := tx.Create(&newUser); dbc.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusUnprocessableEntity,dbc.Error )
		} else {
			tx.Commit()
			newUser.Password = ""
			json := jsonFormatter.FormatJsonSuccess("user","create",&newUser)
			c.JSON(http.StatusOK, &json)

		}
	} else {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

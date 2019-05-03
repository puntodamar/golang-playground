package authentication

import (
	"github.com/gin-gonic/gin"
	"helpers/formatters"

	"net/http"

	. "configs/database"
	global_helper "helpers/global"
	auth_validator "helpers/validator/requests/v1/auth"
	model "models/v1"
)


func Register(c *gin.Context) {
	db 					:= Db
	req 				:= &auth_validator.RegisterForm{}
	jsonFormatter 		:= &formatters.JsonFormatter{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			jsonFormatter.FormatJsonErrorFromString(err.Error()))
	}

	hash, _ 	:= global_helper.HashPassword(req.Password)

	req.Password = hash


	var ok, err = req.Validate(db)

	if ok {

		tx := db.Begin()
		newUser 	:= model.User{
			Name		: req.Name,
			Username	: req.Username,
			Email		: req.Email,
			Password	: req.Password,
		}

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




package authentication

import (
	"api/requests/v1/auth"
	. "configs/database"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"helpers/validator"
	model_validator "helpers/validator/model"
	model "models/v1"
	"net/http"
)


func Register(c *gin.Context) {
	db 					:= Db
	req 				:= &auth.RegisterForm{}
	jsonFormatter 		:= &validator.JsonFormatter{}
	validationErrors 	:= map[string]string{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		json 				:= jsonFormatter.FormatJsonErrorFromString(err.Error())
		validationErrors 	= json
	}

	hash, _ 	:= HashPassword(req.Password)
	newUser 	:= model.User{
		Name		: req.Name,
		Username	: req.Username,
		Email		: req.Email,
		Password	: hash,
	}


	var ok, msg = validRegistration(db, &newUser)
	if ok {
		tx := db.Begin()

		if dbc := tx.Create(&newUser); dbc.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusUnprocessableEntity,dbc.Error )
		} else {
			//tx.Commit()
			tx.Rollback()
			newUser.Password = ""
			json := jsonFormatter.FormatJsonSuccess("user","create",&newUser)
			c.JSON(http.StatusOK, &json)

		}
	} else {
		for k , e := range msg {
			validationErrors[k] = e
		}
		c.JSON(http.StatusOK,jsonFormatter.FormatJsonErrorFromMap(validationErrors))
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

////////////////////////////
// PRIVATE
///////////////////////////

func validRegistration(db *gorm.DB,req *model.User) (bool, map[string]string){
	qe := map[string]string{}

	if model_validator.Exists(db, model.User{}, "username", req.Username){
		qe["username"] = model_validator.DataExists("username")
	}

	if(len(qe) > 0){
		return false,qe
	} else {
		return true,qe
	}

}
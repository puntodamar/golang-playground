package authentication

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"time"
	"net/http"

	"api/requests/v1/auth"
	. "configs/database"
	validation "helpers/validations/model"
	model "models/v1"
)

func Register(c *gin.Context) {
	db 	:= Db
	req := &auth.RegisterForm{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errs := req.Validate("")

	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	hash, _ 	:= HashPassword(req.Password)

	if validation.Exists(db, model.User{},"username", req.Username){
		c.JSON(http.StatusUnprocessableEntity, errors.New("exists"))
		return
	} else {
			newUser := model.User{
				Name		: req.Name,
				Username	: req.Username,
				Email		: req.Email,
				Password	: hash,
				CreatedAt	: time.Now(),
			}

			tx := db.Begin()
			if dbc := tx.Create(&newUser); dbc.Error != nil {
				tx.Rollback()
				c.JSON(http.StatusUnprocessableEntity,dbc.Error )
				return
			} else {
				tx.Commit()
				c.JSON(http.StatusOK, newUser)
			}
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

package v1

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {

	// columns
	ID 			int 		`gorm:"primary_key" json:"id"`
	Name        string		`gorm:"type:varchar(150);unique_index;not null" json:"name"`
	Username	string  	`gorm:"type:varchar(20);unique_index;not null" json:"username"`
	Email       string  	`gorm:"type:varchar(100);unique_index;not null" json:"email"`
	Password	string		`gorm:"type:varchar(255);" json:"password"`
	CreatedAt	time.Time 	`gorm:"type:datetime;" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"type:datetime;" json:"updated_at"`
	DeletedAt	*time.Time 	`gorm:"type:datetime;" json:"deleted_at"`
}

func ListUsers(db *gorm.DB) *gorm.DB {

	return db.Select("id, name, username, email, created_at, updated_at, deleted_at")
}
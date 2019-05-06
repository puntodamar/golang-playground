package v1

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {

	// columns
	ID 			uint 		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name        string		`gorm:"type:varchar(150);unique_index;not null" json:"name"`
	Username	string  	`gorm:"type:varchar(20);unique_index;not null" index:"users_username_uindex" json:"username"`
	Email       string  	`gorm:"type:varchar(100);unique_index;not null" index:"users_email_uindex" json:"email"`
	Password	string		`gorm:"type:varchar(255);" json:"password"`
	CreatedAt	*time.Time 	`gorm:"type:datetime;" sql:"DEFAULT:'current_timestamp'" json:"created_at"`
	UpdatedAt	*time.Time 	`gorm:"type:datetime;" json:"updated_at"`
	DeletedAt	*time.Time 	`gorm:"type:datetime;" json:"deleted_at"`
}

func ListUsers(db *gorm.DB) *gorm.DB {

	return db.Select("id, name, username, email, created_at, updated_at, deleted_at")
}
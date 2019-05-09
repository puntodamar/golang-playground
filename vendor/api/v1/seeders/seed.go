package seeders

import (
	"github.com/gin-gonic/gin"
	//"time"

)
//
func Seed(c *gin.Context) {
//	db 	:= Db
//
//	passwordPunto,_ := HashPassword("puntodamarpassword")
//	passwordHeru,_ 	:= HashPassword("herupassword")
//
//	users := []User{
//		{
//			Name	: "Punto Damar P",
//			Username: "puntodamar",
//			Email	: "puntodamar@gmail.com",
//			Password: passwordPunto,
//
//		},
//		{
//			Name	: "Dwi Heru",
//			Username: "heru",
//			Email	: "heru@gmail.com",
//			Password: passwordHeru,
//		},
//	}
//
//	publishers := []Publisher{
//		{Name : "Bloomsbury"},
//		{Name : "Gramedia"},
//		{Name : "Bentang"},
//	}
//
//	books := []Book{
//		{
//
//			Publisher	: publishers[0],
//			Name		: "Harry Potter",
//			PublishedAt	: time.Date(1995,1,01,0, 0, 0,0, time.UTC),
//		},
//		{
//			Publisher	: publishers[2],
//			Name 		: "Laskar Pelangi",
//			PublishedAt	: time.Date(2006,3,20,0, 0, 0,0, time.UTC),
//		},
//	}
//
//	for _, user := range users{
//		db.Debug().Create(&user)
//	}
//
//	authors := []Author{
//		{Name: "J.K. Rowling"},
//		{Name: "Andrea Hirata"},
//	}
//
//	bookAuthors := []BookAuthor{
//		{
//			Book	: books[0],
//			Author	: authors[0],
//		},
//		{
//			Book	: books[1],
//			Author 	: authors[1],
//		},
//	}
//
//	tags := []Tag{
//		{Name: "Action"},
//		{Name: "Adventure"},
//		{Name: "Mystery"},
//	}
//
//
//	db.Debug().Create(publishers[1])
//
//	for _, bookAuthor := range bookAuthors{
//		db.Debug().Create(&bookAuthor)
//	}
//
//	for _, tag := range tags{
//		db.Debug().Create(&tag)
//	}
}
package helpers

import (
	"github.com/jinzhu/gorm"
	"math"
)


type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	Model 	interface{}
}

type Paginator struct {
	TotalResults 	int         `json:"total_results"`
	TotalPage   	int         `json:"total_page"`
	Results     	interface{} `json:"results"`
	Page        	int         `json:"page"`
	PrevPage    	int         `json:"prev_page"`
	NextPage    	int         `json:"next_page"`
}


func Paginate(p Param)  interface{} {

	var offset 		int
	var count 		int
	var paginator 	Paginator

	done 	:= make(chan bool, 1)
	db 		:= p.DB

	go countRecords(db, p.Model, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(p.Model)
	<-done

	paginator.TotalPage 	= int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	paginator.Results		= p.Model
	paginator.Page 			= p.Page
	paginator.TotalResults 	= count

	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}


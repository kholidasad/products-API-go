package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	// gorm.Model
	ID int64 `json:"id"`
	ProductCode string `json:"product_code"`
	Name  string `json:"name"`
    SubCategory string `json:"subcategory"`
    Brand string `json:"brand"`
    Price float32 `json:"price"`
	Status bool `json:"status" sql:"DEFAULT:true"`
}

type ProductsResponse struct {
	Data []Product `json:"data"`
	Count int64 `json:"count"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Product{})
	return db
}
package model

import (
	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model
	ID int64 `gorm: "unique" json: "ID"`
	product_code   string  `json:"product_code"`
	name  string  `json:"name"`
    subcategory  string  `json:"subcategory"`
    brand  string  `json:"brand"`
    price  float32  `json:"price"`
	status bool `json:"status"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Products{})
	return db
}
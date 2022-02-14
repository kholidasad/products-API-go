package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kholidasad/products-API-go/app/model"
)

func GetProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Products{}

	db.Find(products)
	respondJSON(w, http.StatusOK, products)
}

func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 8)

	product := getProductOr404(db, id, w, r)
	if product == nil {
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func PostProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := model.Products{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&products); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&products).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, products)
}

// func updateProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
// }

// func deleteProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
// }

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func getProductOr404(db *gorm.DB, id int64, w http.ResponseWriter, r *http.Request) *model.Products {
	product := model.Products{}
	if err := db.First(&product, model.Products{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &product
}

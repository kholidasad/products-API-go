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
	products := []model.Product{}
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "8"
	}
	limitInt, _ := strconv.ParseInt(limit, 10, 8)
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.ParseInt(page, 10, 8)
	
	offset := (pageInt - 1) * limitInt
	var count int64
	db.Offset(offset).Limit(limitInt).Find(&products).Count(&count)
	productResponse := &model.ProductsResponse{
		Data : products,
		Count : count,
		Page : int(pageInt),
		Limit : int(limitInt),
	}
	respondJSON(w, http.StatusOK, productResponse)
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
	products := model.Product{}

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

func UpdateProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 8)

	products := getProductOr404(db, id, w, r)
	if products == nil {
		return
	}

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

func DeleteProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 8)

	products := getProductOr404(db, id, w, r)
	if products == nil {
		return
	}

	if err := db.Delete(&products).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, nil)
}

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

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func getProductOr404(db *gorm.DB, id int64, w http.ResponseWriter, r *http.Request) *model.Product {
	product := model.Product{}
	if err := db.First(&product, model.Product{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &product
}

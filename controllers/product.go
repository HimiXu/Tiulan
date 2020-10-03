package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HimiXu/Tiulan/stock-service/db"
	"github.com/HimiXu/Tiulan/stock-service/models"

	"github.com/julienschmidt/httprouter"
)

func GetProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	fmt.Println("got -", id)
	res := db.GetProduct(id)
	fmt.Println("found -", res.OK)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if !res.OK {
		json.NewEncoder(w).Encode("Product does not exist")
		return
	}
	json.NewEncoder(w).Encode(*res.P)
}

func CreateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)
	success := db.CreateProduct(p)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if !success {
		json.NewEncoder(w).Encode("Product already exists")
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	toUpdate := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&toUpdate)
	success := db.UpdateProduct(id, toUpdate)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if !success {
		json.NewEncoder(w).Encode("Product not found or fields incorrect")
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	db.DeleteProduct(id)
	w.WriteHeader(http.StatusOK)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sp := db.GetAllProducts()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sp)
}

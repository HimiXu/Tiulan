package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HimiXu/Tiulan/stock-service/db"
	"github.com/HimiXu/Tiulan/stock-service/item"

	"github.com/julienschmidt/httprouter"
)

func GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemName := params.ByName("name")
	itm := db.GetItemByName(itemName)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if itm == nil {
		json.NewEncoder(w).Encode("Item does not exist")
		return
	}
	json.NewEncoder(w).Encode(*itm)
}

func CreateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var it item.Item
	json.NewDecoder(r.Body).Decode(&it)
	success := db.CreateItem(it)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if !success {
		json.NewEncoder(w).Encode("Item already exists")
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemName := params.ByName("name")
	toUpdate := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&toUpdate)
	success := db.UpdateItemByName(itemName, toUpdate)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if !success {
		json.NewEncoder(w).Encode("Item not found or fields incorrect")
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemName := params.ByName("name")
	db.DeleteItemByName(itemName)
	w.WriteHeader(http.StatusOK)
}

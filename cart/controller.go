package cart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HimiXu/Tiulan/models"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

// TODO: solve concurrency issues using a sessions manager
var store *sessions.CookieStore

func init() {
	// TODO: update this to a configurable key
	store = sessions.NewCookieStore([]byte("test"))
}

func SetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productID := params.ByName("productID")
	amount := params.ByName("amount")
	fmt.Println("GET:", "id:", productID, "amount:", amount)
	session, _ := store.Get(r, "cart")
	session.Values[productID], _ = strconv.Atoi(amount)
	fmt.Println(session.Values)
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cart", http.StatusOK)
}

func DeleteItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productID := params.ByName("productID")
	session, _ := store.Get(r, "cart")
	fmt.Println("DELETE", "id:", productID)
	delete(session.Values, productID)
	fmt.Println(session.Values)
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cart", http.StatusOK)
}

func GetCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session, _ := store.Get(r, "cart")
	fmt.Println("GET ALL")
	fmt.Println(session.Values)
	w.Header().Set("Content-Type", "application/json")
	cartMap := make(map[string]int)
	for productID, amount := range session.Values {
		cartMap[productID.(string)] = amount.(int)
	}
	crt := models.ToCart(cartMap)
	fmt.Println(crt)
	JSON, err := json.Marshal(crt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSON))
	json.NewEncoder(w).Encode(crt)
}

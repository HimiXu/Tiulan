package main

import (
	"net/http"

	"github.com/HimiXu/Tiulan/cart"
	"github.com/HimiXu/Tiulan/product"
	"github.com/HimiXu/Tiulan/reservations"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/cart/", http.StripPrefix("/cart", cart.NewRouter()))
	mux.Handle("/products/", http.StripPrefix("/products", product.NewRouter()))
	mux.Handle("/reservations/", http.StripPrefix("/reservations", reservations.NewRouter()))

	http.ListenAndServe(":8080", mux)
}

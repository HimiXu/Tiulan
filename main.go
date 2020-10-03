package main

import (
	"net/http"

	"github.com/HimiXu/Tiulan/cart"
	"github.com/HimiXu/Tiulan/product"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/cart/", http.StripPrefix("/cart", cart.NewRouter()))
	mux.Handle("/products/", http.StripPrefix("/products", product.NewRouter()))

	http.ListenAndServe(":8080", mux)
}

package main

import (
	"net/http"

	"github.com/HimiXu/Tiulan/stock-service/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/products/:id", controllers.GetProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.POST("/products/create", controllers.CreateProduct)
	r.GET("/products", controllers.GetAllProducts)
	http.ListenAndServe(":8081", r)
}

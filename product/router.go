package product

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/:id", GetProduct)
	r.PUT("/:id", UpdateProduct)
	r.DELETE("/:id", DeleteProduct)
	r.POST("/create", CreateProduct)
	r.GET("/", GetProducts)
	return r
}

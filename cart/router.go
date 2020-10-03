package cart

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	r := httprouter.New()
	r.POST("/:productID/:amount", SetItem)
	r.DELETE("/:productID", DeleteItem)
	r.GET("/", GetCart)
	return r
}

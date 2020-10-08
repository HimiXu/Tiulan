package reservations

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/:id", GetReservation)
	r.DELETE("/:id", DeleteReservation)
	r.POST("/create", CreateReservation)
	return r
}

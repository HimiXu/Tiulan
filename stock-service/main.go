package main

import (
	"fmt"
	"net/http"

	"github.com/HimiXu/Tiulan/stock-service/controllers"
	"github.com/HimiXu/Tiulan/stock-service/db"
	"github.com/HimiXu/Tiulan/stock-service/item"
	"github.com/julienschmidt/httprouter"
)

func main() {

	it := item.NewBuilder().
		Name("tent").
		Price(100).
		Amount(5).
		Description("for 4 people or hilla").
		Build()
	fmt.Println(it)

	db.CreateItem(*it)
	fmt.Println(db.GetItemByName("tent"))

	r := httprouter.New()
	r.GET("/get-item/:name", controllers.GetItem)
	r.POST("/create-item", controllers.CreateItem)
	r.PUT("/update-item/:name", controllers.UpdateItem)
	r.DELETE("/delete-item/:name", controllers.DeleteItem)
	http.ListenAndServe(":8081", r)
}

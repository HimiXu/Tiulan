package main

import (
	"fmt"

	"github.com/HimiXu/Tiulan/stock-service/db"
	"github.com/HimiXu/Tiulan/stock-service/item"
)

func main() {

	item := item.NewBuilder().
		Name("tent").
		Price(100).
		Amount(5).
		Description("for 4 people or hilla").
		Build()
	fmt.Println(item)

	db.Connect()

	// r := httprouter.New()
	// http.ListenAndServe(":8081", r)
}

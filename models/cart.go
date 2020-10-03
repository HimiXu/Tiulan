package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
	ID     primitive.ObjectID `json:"_id,omitempty"`
	Name   string             `json:"name"`
	Price  float64            `json:"price"`
	Amount int                `josn:"amount"`
}

type Cart struct {
	Items      map[string]CartItem `json:"items"`
	TotalPrice float64             `json:"total_price"`
	// Discounts
}

func ToCart(cartMap map[string]int) Cart {
	cart := Cart{Items: make(map[string]CartItem), TotalPrice: 0}
	for productID, amount := range cartMap {
		res := GetProduct(productID)
		if res.OK {
			prod := res.P
			cart.TotalPrice += prod.Price * float64(amount)
			cart.Items[productID] = CartItem{
				ID:     prod.ID,
				Name:   prod.Name,
				Price:  prod.Price,
				Amount: amount,
			}
		}
	}
	return cart
}

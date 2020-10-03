package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Price       float64            `json:"price" bson:"price"`
	Amount      float64            `json:"amount" bson:"amount"`
	Description string             `json:"description,omitempty" bson:"description"`
}

func (p *Product) Update(toUpdate map[string]interface{}) {
	for k, v := range toUpdate {
		switch k {
		case "name":
			p.Name = v.(string)
		case "price":
			p.Price = v.(float64)
		case "amount":
			p.Amount = v.(float64)
		case "description":
			p.Description = v.(string)
		}
	}
}

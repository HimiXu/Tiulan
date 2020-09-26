package item

import (
	"fmt"
)

// Item - basic item type
type Item struct {
	Name        string  `json:"name" bson:"name"`
	Price       float64 `json:"price" bson:"price"`
	Amount      float64 `json:"amount" bson:"amount"`
	Description string  `json:"description" bson:"description"`
}

// Builder - object that creates an item
type Builder struct {
	item    Item
	actions []func(*Item)
}

// NewBuilder - create a new ItemBuilder
func NewBuilder() *Builder {
	return &Builder{item: Item{}}
}

func (b *Builder) Name(name string) *Builder {
	b.actions = append(b.actions, func(item *Item) {
		item.Name = name
	})
	return b
}

func (b *Builder) Price(price float64) *Builder {
	b.actions = append(b.actions, func(item *Item) {
		item.Price = price
	})
	return b
}

func (b *Builder) Amount(amount float64) *Builder {
	b.actions = append(b.actions, func(item *Item) {
		item.Amount = amount
	})
	return b
}

func (b *Builder) Description(description string) *Builder {
	b.actions = append(b.actions, func(item *Item) {
		item.Description = description
	})
	return b
}

func (b *Builder) Build() *Item {
	for _, action := range b.actions {
		action(&b.item)
	}
	return &b.item
}

func (it *Item) UpdateItem(toUpdate map[string]interface{}) {
	for k, v := range toUpdate {
		fmt.Println(k, v)
		switch k {
		case "name":
			it.Name = v.(string)
		case "price":
			it.Price = v.(float64)
		case "amount":
			it.Amount = v.(float64)
		case "description":
			it.Description = v.(string)
		}
	}
}

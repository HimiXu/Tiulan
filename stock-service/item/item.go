package item

import (
	"gopkg.in/mgo.v2/bson"
)

// Item - basic item type
type Item struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Price       uint          `json:"price" bson:"price"`
	Amount      uint          `json:"amount" bson:"amount"`
	Description string        `json:"description" bson:"description"`
}

// Builder - object that creates an item
type Builder struct {
	item    Item
	actions []func(*Item)
}

// NewBuilder - create a new ItemBuilder
func NewBuilder() *Builder {
	return &Builder{item: Item{Id: bson.NewObjectId()}}
}

func (b *Builder) Name(name string) *Builder {
	b.actions = append(b.actions, func(item *Item) {
		item.Name = name
	})
	return b
}

func (b *Builder) Price(price uint) *Builder {
	b.actions = append(b.actions, func(item *Item) {
		item.Price = price
	})
	return b
}

func (b *Builder) Amount(amount uint) *Builder {
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

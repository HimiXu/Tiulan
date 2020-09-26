package db

import (
	"context"
	"log"
	"time"

	"github.com/HimiXu/Tiulan/stock-service/item"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createSession() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}

func CreateItem(it item.Item) bool {
	if itemExistByName(it.Name) {
		return false
	}
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Items")
	itemBSON, err := bson.Marshal(it)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = collection.InsertOne(ctx, itemBSON)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetItemByName(name string) *item.Item {
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Items")

	cur, _ := collection.Find(ctx, bson.M{"name": name})
	itms := make([]item.Item, 0)
	cur.All(ctx, &itms)
	if len(itms) > 1 {
		log.Fatal("more than one item with the same name")
		return nil
	}
	if len(itms) == 0 {
		return nil
	}
	return &itms[0]
}

func UpdateItemByName(name string, toUpdate map[string]interface{}) bool {
	it := GetItemByName(name)
	if it == nil {
		return false
	}
	it.UpdateItem(toUpdate)
	DeleteItemByName(name)
	res := CreateItem(*it)
	return res
}

func DeleteItemByName(name string) {
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Items")
	collection.DeleteMany(ctx, bson.M{"name": name})
}

func itemExistByName(name string) bool {
	if GetItemByName(name) == nil {
		return false
	}
	return true
}

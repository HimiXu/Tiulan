package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Result struct {
	P  *Product
	OK bool
}

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

func CreateProduct(p Product) bool {
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Products")
	pBSON, err := bson.Marshal(p)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = collection.InsertOne(ctx, pBSON)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
func GetAllProducts() []Product {
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Products")

	cur, _ := collection.Find(ctx, bson.M{})
	sp := make([]Product, 0)
	cur.All(ctx, &sp)
	return sp
}

func GetProduct(id string) Result {
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Products")
	oid, _ := primitive.ObjectIDFromHex(id)
	cur, _ := collection.Find(ctx, bson.M{"_id": oid})
	sp := make([]Product, 0)
	cur.All(ctx, &sp)
	if len(sp) > 1 {
		log.Fatal("more than one item with the same id")
		return Result{P: nil, OK: false}
	}
	if len(sp) == 0 {
		return Result{P: nil, OK: false}
	}
	return Result{P: &sp[0], OK: true}
}

func UpdateProduct(id string, toUpdate map[string]interface{}) bool {
	r := GetProduct(id)
	if !r.OK {
		return false
	}
	p := r.P
	p.Update(toUpdate)
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Products")
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": p})
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DeleteProduct(id string) {
	client, ctx := createSession()
	defer client.Disconnect(ctx)
	collection := client.Database("Stock").Collection("Products")
	oid, _ := primitive.ObjectIDFromHex(id)
	collection.DeleteMany(ctx, bson.M{"_id": oid})
}

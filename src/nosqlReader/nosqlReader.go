package nosqlReader

import (
	"context"
	"github.com/google/uuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

const (
	mongoClient     = "mongodb://localhost:27017"
	datebase        = "Orders"
	orderCollection = "Order"
)

type Orders struct {
	OrderId uuid.UUID `bson:"OrderId"`
	Req     int       `bson:"Req"`
}

func getDateBase() *mongo.Database {
	client, err := mongo.NewClient(mongoClient)
	if err != nil {
		panic("cannot connect to MongoServer")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		defer cancel()
		panic("cannot connect to MongoServer")
	}
	return client.Database(datebase)
}

func GetOne() Orders {
	db := getDateBase()
	defer db.Client().Disconnect(context.Background())
	collection := db.Collection(orderCollection)
	filter := bson.M{"Req": 3}
	var orders Orders
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&orders)
	if err != nil {
		cancel()
		panic("find error")
	}

	return orders
}

func GetMany() []Orders {
	db := getDateBase()
	defer db.Client().Disconnect(context.Background())
	collection := db.Collection(orderCollection)
	filter := bson.M{"Req": bson.M{"$lte": 3}}
	var orders []Orders
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order Orders
		cursor.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	return orders
}

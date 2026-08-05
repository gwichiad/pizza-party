package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func startDB(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	
	pingCtx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

type store struct {
	collection *mongo.Collection
}

func newStore(client *mongo.Client) *store {
	collection := client.Database("pizza_party").Collection("satellite_data") 
	return &store{collection: collection}
}

func (store *store) insert(ctx context.Context, data SatelliteResponse) error {
	_, err := store.collection.InsertOne(ctx, &data)
	return err
}

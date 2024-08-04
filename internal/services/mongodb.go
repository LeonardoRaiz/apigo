package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Result struct {
	Terms     []string  `json:"terms"`
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
	Domains   []string  `json:"domains"`
}

func saveResults(terms []string, email string, domains []string) error {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerSelectionTimeout(10 * time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("brand_monitor").Collection("results")
	_, err = collection.InsertOne(context.Background(), Result{
		Terms:     terms,
		Email:     email,
		Timestamp: time.Now(),
		Domains:   domains,
	})
	return err
}

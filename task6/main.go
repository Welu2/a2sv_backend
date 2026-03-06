package main

import (
	"context"
	"log"
	"task_manager/data"
	"task_manager/router"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 1. Setup MongoDB Connection
	uri := "mongodb+srv://welelabekeleug_db_user:password@cluster0.n6xzoew.mongodb.net//task6?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
    log.Fatal("Could not connect to MongoDB Atlas. Check your IP Whitelist: ", err)
}

	// 2. Initialize both services with the database client
	// This ensures both Task and User collections are ready
	data.ConnectDB(client)        // Sets up Task collection
	data.SetUserCollection(client) // Sets up User collection

	// 3. Start Router
	r := router.SetupRouter()
	r.Run(":8080")
}

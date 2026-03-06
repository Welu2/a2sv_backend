package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	// Use your local MongoDB URI
	data.ConnectDB("mongodb+srv://welelabekeleug_db_user:q9sC8pS2xZJYfWZJ@cluster0.n6xzoew.mongodb.net/task5?retryWrites=true&w=majority")

	r := router.SetupRouter()
	r.Run(":8080")
}

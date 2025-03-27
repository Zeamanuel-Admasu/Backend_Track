package main

import (
	"context"
	"log"
	"time"

	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/router"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := data.ConnectDB(ctx); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	r := router.SetupRouter()

	r.Run()
}

package main

import (
	configs "api-project/src/internal/config"
	"api-project/src/internal/router"
	"log"
)

func main() {
	cfg := configs.Load()

	if cfg != nil {
		log.Fatal(cfg)
	}

	route := router.SetupRouter()

	port := configs.GetPort()

	log.Println("API running on port", port)
	if err := route.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

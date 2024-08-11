package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/config"
	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/db"
	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/router"
)

func main() {
	// Define possible locations for the .env file
	envLocations := []string{
		".env",
		"../.env",
		"../../.env",
		"/root/.env",
	}

	// Try to load the .env file from each location
	var envLoaded bool
	for _, loc := range envLocations {
		if err := godotenv.Load(loc); err == nil {
			log.Printf("Loaded .env file from %s", loc)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		log.Fatal("Error loading .env file")
	}

	_, configErr := config.New()
	if configErr != nil {
		log.Fatal(configErr)
	}
	log.Printf("Config initialization succeeded")

	db.Init()
	defer db.Close()

	r := gin.Default()
	router.SetRoutes(r)
	gatewayPort := os.Getenv("GATEWAY_PORT")
	r.Run(gatewayPort)
}

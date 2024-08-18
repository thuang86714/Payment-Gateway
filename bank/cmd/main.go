package main

import (
	"log"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/config"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/db"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/router"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/repository"
    "github.com/processout-hiring/payment-gateway-thuang86714/bank/service"
    "github.com/processout-hiring/payment-gateway-thuang86714/bank/controller"
)

func main() {
    // Load .env file
    if err := loadEnv(); err != nil {
        log.Fatal(err)
    }

	// Load configuration
	conf, err := config.New(config.OsEnvReader{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Config initialization succeeded")

	//Initialize database
	database, err := db.NewDB(conf)
	if err != nil {
        log.Fatal(err)
    }
	defer db.CloseDB(database)

	// Create repository
    repo := repository.NewGormRepository(database)

    // Create service
    svc := service.NewService(repo)

    // Create controller
    ctrl := controller.NewController(svc)

	// Set up router
	r := gin.Default()
	router.SetRoutes(r, ctrl)

	//Start server
	bankPort := os.Getenv("BANK_PORT")
	if err := r.Run(bankPort); err != nil {
        log.Fatal(err)
    }
}

func loadEnv() error {
    // Define possible locations for the .env file
    envLocations := []string{
        ".env",
        "../.env",
        "../../.env",
        "/root/.env",
    }

    // Try to load the .env file from each location
    for _, loc := range envLocations {
        if err := godotenv.Load(loc); err == nil {
            log.Printf("Loaded .env file from %s", loc)
            return nil
        }
    }

    return fmt.Errorf("error loading .env file")
}
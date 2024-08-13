package db

import (
	"fmt"
	"log"

	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/config"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable that holds the database connection
var DB *gorm.DB

// Init initializes the database connection and performs schema migration
func Init() {
	var err error
	// Construct the database connection string (DSN) using configuration values
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Conf.PostgresHost,
		config.Conf.PostgresUser,
		config.Conf.PostgresPassword,
		config.Conf.PostgresDB,
		config.Conf.PostgresPort,
	)
	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Automatically migrate the database schema for PostResponse model
	DB.AutoMigrate(&models.PostResponse{})
	log.Println("DB initialize succeeded")
}

// Close closes the database connection
func Close() {
	// Get the underlying sql.DB object
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	// Close the database connection
	sqlDB.Close()
}

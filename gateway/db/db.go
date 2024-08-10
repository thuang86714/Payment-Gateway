package db

import (
	"fmt"
	"log"

	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/config"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Conf.PostgresHost,
		config.Conf.PostgresUser,
		config.Conf.PostgresPassword,
		config.Conf.PostgresDB,
		config.Conf.PostgresPort,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.PostResponse{})
	log.Println("DB initialize succeeded")
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	sqlDB.Close()
}

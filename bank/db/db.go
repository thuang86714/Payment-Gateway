package db

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/config"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
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
	DB.AutoMigrate(&models.Balance{}, &models.InvoiceID{})
	log.Println("DB initialize succeeded")
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	sqlDB.Close()
}

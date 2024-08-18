package db

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/config"
	"github.com/processout-hiring/payment-gateway-thuang86714/shared/models"
)

func NewDB(conf *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.PostgresHost,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDB,
		conf.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Balance{}, &models.InvoiceID{})
	if err != nil {
        return nil, fmt.Errorf("failed to migrate database: %v", err)
    }

	log.Println("DB initialize succeeded")
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
        log.Printf("failed to get sql.DB from gorm.DB: %v", err)
        return
    }
    if err := sqlDB.Close(); err != nil {
        log.Printf("failed to close database connection: %v", err)
    }
}

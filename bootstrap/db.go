package bootstrap

import (
	"fmt"
	"gin-app/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "user"),
		utils.GetEnv("DB_PASSWORD", "user"),
		utils.GetEnv("DB_NAME", "db"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	fmt.Println("Database connection string:", connectionString)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("connected to database")
	return db
}

package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

type DatabaseConn struct {
	db *gorm.DB
}

var dbInstance *DatabaseConn
var once sync.Once

func GetDatabaseConnection() *gorm.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found — assuming production environment")
		}
		databaseURL := os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			log.Fatalf("DATABASE_URL not set in the environment")
		}
		conn, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		dbInstance = &DatabaseConn{db: conn}
	})

	return dbInstance.db
}

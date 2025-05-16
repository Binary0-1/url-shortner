package migrations

import (
    "fmt"
    "urlshort/db"
    "urlshort/models"
    "log"
)

func RunMigrations() {
    dbConn := db.GetDatabaseConnection()

    err := dbConn.AutoMigrate(
        &models.URL{},
    )
    if err != nil {
        log.Fatalf("Error running migrations: %v", err)
    }

    fmt.Println("Migrations ran successfully!")
}

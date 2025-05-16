package main

import (
	"fmt"
	"log"
	"net/http"
	"urlshort/routes"
	"urlshort/db/migrations"
)

func main() {
	fmt.Println("Hello, Go!")
	migrations.RunMigrations();
	mux := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8081", mux))
}

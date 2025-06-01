package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"urlshort/db/migrations"
	"urlshort/routes"
)

func main() {
	var runMigrations bool

	flag.BoolVar(&runMigrations, "migrate", false, "Set this flag to run migrations")
	flag.Parse()

	if runMigrations {
		migrations.RunMigrations()
	} else {
		fmt.Print("Skipping migrations")

	}
	mux := routes.RegisterRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	log.Printf("Running server on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}

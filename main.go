package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":8082", mux))
}

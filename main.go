package main

import (
	"database/sql"
	"fmt"
	"locationapi/http"
	"locationapi/locations"
	"locationapi/postgres"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	postgresHost     = "postgres"
	postgresPort     = 5432
	postgresUser     = "locationapi"
	postgresPassword = "postgres"
	postgresDBname   = "locationapi"
)

func main() {

	db := setupPostgres(
		postgresHost,
		postgresUser,
		postgresPassword,
		postgresDBname,
		postgresPort,
	)
	defer db.Close()

	repo := postgres.NewRepository(db)
	locationService := locations.NewLocationService(repo)

	httpServer := http.NewServer(locationService)
	httpServer.Open()
}

func setupPostgres(host, user, password, dbname string, port uint) *sql.DB {

	postgresConfig := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresPort,
		postgresDBname,
		"disable",
	)
	db, err := sql.Open("postgres", postgresConfig)
	if err != nil {
		log.Println("could not open connection to postgres")
		panic(err)
	}
	log.Println("opened postgres connection")
	time.Sleep(5 * time.Second)
	err = db.Ping()
	if err != nil {
		log.Println("could not ping database")
		panic(err)
	}

	log.Println("Successfully connected!")
	return db
}

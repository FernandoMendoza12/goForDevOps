package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	webPort = "80"
)

type Config struct {
}

func main() {

	app := Config{}

	dsn := "host=postgres user=admin password=password dbname=users port=5432 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("No se pudo conectar a la base de datos")
	}
	log.Println("Starting on port :", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}

}

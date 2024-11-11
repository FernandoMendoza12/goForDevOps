package main

import (
	"fmt"
	"net/http"
)

const (
	webPort = "80"
)

type Config struct{}

func main() {

	app := Config{}
	fmt.Println("Starting on webPort:" + webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	srv.ListenAndServe()
}

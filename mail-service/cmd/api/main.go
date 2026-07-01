package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const WebPort = "80"

func main() {
	app := Config{}
	log.Println("Starting mail service on port", WebPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

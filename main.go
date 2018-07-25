package main

import (
	"net/http"
	"os"
	"time"
	"./controller"
	_ "./db/mongodb"
	_ "./option"

)

const (
	serverUrl = "0.0.0.0"
	defaultPort = "8000"
	readTimeout = 100 * time.Second
	writeTimeout = 200 * time.Second
	)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	handler := controllers.New()

	server := &http.Server{
		Addr:   serverUrl + port,
		Handler: handler,
		WriteTimeout:      writeTimeout,
		ReadTimeout:       readTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


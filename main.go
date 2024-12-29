package main

import (
	"log"
	"net/http"

	"github.com/Kaltner/oauth_test/app"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	mux := app.ServeMultiplexer()
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"io"
	"ligmanstark/pastebin-go/handlers"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"
)

func main() { // This is a placeholder for the main function.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, (" Host: " + appHost + " Port: " + appPort))
	}

	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/pastebin/create", handlers.CreatePastebinHandler)

	http.HandleFunc("/pastebin", handlers.GetPastebinBySlug)

	log.Fatal(http.ListenAndServe(":5555", nil))

}

package main

import (
	"io"
	"ligmanstark/pastebin-go/handlers"
	"log"
	"net/http"
)

func main() { // This is a placeholder for the main function.

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, ("Hello from 5555!"))
	}

	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/users/create", handlers.CreateUserHandler)

	http.HandleFunc("/users", handlers.GetAllUsersHandler)

	log.Fatal(http.ListenAndServe(":5555", nil))

}

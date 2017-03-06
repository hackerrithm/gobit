package main

import (
	"log"
	"net/http"

	"restful-api-standard-library/handlers"
	"restful-api-standard-library/storage"
)

func main() {

	db := storage.NewInMemoryDB()
	mux := http.NewServeMux()
	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))

	log.Printf("serving on port 5000")

	err := http.ListenAndServe(":5000", mux)

	log.Fatal(err)
}

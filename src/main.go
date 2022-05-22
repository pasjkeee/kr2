package main

import (
	"github.com/gorilla/mux"
	"kr2/pkg/handler"
	"log"
	"net/http"
)

func main() {

	for i := range handler.Field {
		handler.Field[i] = make([]int, 3, 3)
	}

	router := mux.NewRouter()

	router.HandleFunc("/move", handler.GetMove).Methods(http.MethodGet).Queries("x", "{x:[0-9]+}", "y", "{y:[0-9]+}")
	router.HandleFunc("/reset", handler.Reset).Methods(http.MethodGet)
	router.Handle("/", http.FileServer(http.Dir("./src")))

	log.Fatal(http.ListenAndServe(":8080", router))
}

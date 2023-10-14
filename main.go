package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func rootTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First backend test\n")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootTest).Methods("GET")
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}

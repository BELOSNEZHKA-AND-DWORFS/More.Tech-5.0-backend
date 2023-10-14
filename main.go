package main

import (
	"fmt"
	"net/http"
	"time"

	"more.tech/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.RootTest).Methods("GET")
	router.HandleFunc("/office/{officeID}", handlers.GetOfficeInfo).Methods("GET") // TODO: what officeID?
	router.HandleFunc("/get_offices", handlers.GetOffices).Methods("POST")
	router.HandleFunc("/atm/{atmID}", handlers.GetAtmInfo).Methods("GET") // TODO: what atmID?
	router.HandleFunc("/get_atms", handlers.GetAtms).Methods("POST")
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}

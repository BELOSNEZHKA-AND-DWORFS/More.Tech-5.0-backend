package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"more.tech/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.RootTest).Methods("GET")
	router.Path("/office").Queries("officeid", "{officeid}").HandlerFunc(handlers.GetOfficeInfo).Methods("GET")
	router.HandleFunc("/get_offices", handlers.GetOffices).Methods("POST")
	router.Path("/atm").Queries("atmid", "{atmid}").HandlerFunc(handlers.GetAtmInfo).Methods("GET")
	router.HandleFunc("/get_atms", handlers.GetAtms).Methods("POST")
	router.HandleFunc("/voice", handlers.VoiceHandler).Methods("POST")
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}

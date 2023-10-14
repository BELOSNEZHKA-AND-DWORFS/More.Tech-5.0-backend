package handlers

import (
	"fmt"
	"net/http"
)

func RootTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First backend test\n")
}

func GetOfficeInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented\n")
}

func GetOffices(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented\n")
}

func GetAtmInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented\n")
}

func GetAtms(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented\n")
}

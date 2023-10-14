package handlers

import (
	"fmt"
	"net/http"
)

func RootTest(w http.ResponseWriter, r *http.Request) {
	result := getOfficeByTitleInArea("втб", 55.649865, 37.622646)
	fmt.Fprintf(w, "%v\n", result)
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

package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed.")
		return
	}
	fmt.Fprintf(w, "Hello there %s", "visitor")
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func addCountries(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	if countryExist(country.Name) {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Country exists.")
		log.Println("Country exists.")
		return
	} else {
		countries = append(countries, country)
		fmt.Fprintf(w, "Country was added.")
		log.Println("Country was added.")
	}
}

func countryExist(Name string) bool {
	for _, v := range countries {
		if v.Name == Name {
			return true
		}
	}
	return false
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Biodata struct {
	Nama string   `json:"nama"`
	Kode []string `json:"kode"`
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/haiqal", greet)

	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/haiqal/middleware1", middleware1(endpoint))
	mux.Handle("/haiqal/middleware2", middleware2(endpoint))

	err := http.ListenAndServe(PORT, mux)

	if err != nil {
		log.Fatal(err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var biodata Biodata

	biodata.Nama = "Haiqal Ramanizar Al Fajri"
	biodata.Kode = append(biodata.Kode, "149368582101-637", "149368582101-639", "149368582100-406")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(biodata)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("halo saya haiqal")
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var biodata Biodata
		biodata.Nama = "Haiqal Ramanizar Al Fajri"
		biodata.Kode = append(biodata.Kode, "149368582101-637", "149368582101-639", "149368582100-406")

		fmt.Printf("%v\n", biodata)
		next.ServeHTTP(w, r)
	})
}

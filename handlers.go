package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"io"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Rest API for Elastic Search!")
}

func SearchAllSeries(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Searching for all series:")
	results := getAllSeries()
	fmt.Fprintln(w, "Results:")
    fmt.Fprintln(w, "", results)
}

func SearchSeriesByName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    serieName := vars["serieName"]
    fmt.Fprintln(w, "Searching for series with name:", serieName)
	results := getSeriesByName(serieName)
	fmt.Fprintln(w, "Results:")
    fmt.Fprintln(w, "", results)
}

func SearchSeriesByRate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    rate := vars["rate"]
    fmt.Fprintln(w, "Searching for series with rate:", rate)
	results := getSeriesByRate(rate)
	fmt.Fprintln(w, "Results:")
    fmt.Fprintln(w, "", results)
}

func CreateSerie(w http.ResponseWriter, r *http.Request) {
	var serie Serie
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &serie); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

	results := createSerie(serie)
	fmt.Fprintln(w, "Creating serie in the index:", serie)
	fmt.Fprintln(w, "Results:")
    fmt.Fprintln(w, "", results)
}
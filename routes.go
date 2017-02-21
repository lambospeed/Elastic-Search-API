package main

import (
    "net/http"
    "github.com/gorilla/mux"
	
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },Route{
        "SearchAllSeries",
        "GET",
        "/all_series",
        SearchAllSeries,
    },Route{
        "SearchSeriesByName",
        "GET",
        "/series_byname/{serieName}",
        SearchSeriesByName,
    },Route{
        "SearchSeriesByRate",
        "GET",
        "/series_byrate/{rate}",
        SearchSeriesByRate,
    },Route{
        "CreateSerie",
        "POST",
        "/create_serie",
        CreateSerie,
    },
}
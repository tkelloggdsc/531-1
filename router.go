package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a REST endpoint that maps common methods
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of routes
type Routes []Route

var routes = Routes{
// TODO: implement
}

// NewRouter establishes the root application router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		// handler = config.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

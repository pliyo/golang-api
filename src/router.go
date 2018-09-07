package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter initialises a new router
func NewRouter() *mux.Router {
	// When StrictSlash is true, if the route path is "/path/", accessing "/path" will redirect to the former and vice versa
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

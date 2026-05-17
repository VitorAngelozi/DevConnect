package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ROUTEs represents the list of routes that will be used in the application
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := UserRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}

package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//its going to generate a router with with routes configured in the future

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	return routes.ConfigureRoutes(router)
}

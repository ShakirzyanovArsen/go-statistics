package middleware

import (
	"github.com/gorilla/mux"
	"slurm_statistics/handler"
	"net/http"
)

/*
RegisterRoutes register routes and return router
 */
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.IndexHandler).Methods("GET")

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs)).Methods("GET")

	return router;
}

package middleware

import "github.com/gorilla/mux"

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	//router.HandleFunc()
	return router;
}

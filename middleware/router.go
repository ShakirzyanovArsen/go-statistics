package middleware

import "github.com/gorilla/mux"

/*
RegisterRoutes register routes and return router
 */
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	//router.HandleFunc()
	return router;
}

package handler

import "github.com/gorilla/mux"

var Router *mux.Router

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/users/{id:[1-9]+}", FindUserById).Methods("GET")
	router.HandleFunc("/users", InsertUser).Methods("POST")

	// products
	router.HandleFunc("/products/{id:[1-9]+}", FindProductById).Methods("GET")
	router.HandleFunc("/products", FindAllProduct).Methods("GET")

	Router = router
}

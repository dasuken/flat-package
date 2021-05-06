package router

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/flat-package/handler"
)

var Router *mux.Router

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/users/{id:[1-9]+}", handler.FindUserById).Methods("GET")
	router.HandleFunc("/users", handler.InsertUser).Methods("POST")

	// products
	router.HandleFunc("/products/{id:[1-9]+}", handler.FindProductById).Methods("GET")
	router.HandleFunc("/products", handler.FindAllProduct).Methods("GET")

	Router = router
}
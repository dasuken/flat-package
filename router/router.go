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

	Router = router
}
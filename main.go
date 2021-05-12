package main

import (
	"github.com/noguchidaisuke/flat-package/handler"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", handler.Router))
}
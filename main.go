package main

import (
	"github.com/noguchidaisuke/flat-package/router"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
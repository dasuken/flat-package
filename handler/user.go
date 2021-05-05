package handler

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/flat-package/db"
	"github.com/noguchidaisuke/flat-package/handler/response"
	"github.com/noguchidaisuke/flat-package/user"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	response.ResponseHelloWorld(w, "Hello World")
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	conn, err := db.NewDBConn()
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := user.GetById(conn, id)
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseUser(w, user)
}


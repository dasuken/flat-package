package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/flat-package/db"
	"github.com/noguchidaisuke/flat-package/handler/response"
	"github.com/noguchidaisuke/flat-package/user"
	"io"
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

func InsertUser(w http.ResponseWriter, r *http.Request) {
	conn, err := db.NewDBConn()
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	name := getUserName(r)

	_, err = user.Add(conn, name)
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "Success")
}


func getUserName(r *http.Request) string {
	var data map[string]string
	b, _ := io.ReadAll(r.Body)
	json.Unmarshal(b, &data)

	return data["name"]
}
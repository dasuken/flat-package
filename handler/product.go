package handler

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/flat-package/db"
	"github.com/noguchidaisuke/flat-package/handler/response"
	"github.com/noguchidaisuke/flat-package/product"
	"net/http"
	"strconv"
)

func FindProductById(w http.ResponseWriter, r *http.Request) {
	conn, err := db.NewDBConn()
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	product, err := product.GetById(conn, id)
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseProduct(w, product)
}

func FindAllProduct(w http.ResponseWriter, r *http.Request) {
	conn, err := db.NewDBConn()
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}


	products, err := product.GetAll(conn)
	if err != nil {
		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseProducts(w, products)
}

//func InsertUser(w http.ResponseWriter, r *http.Request) {
//	conn, err := db.NewDBConn()
//	if err != nil {
//		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	name := getUserName(r)
//
//	_, err = user.Add(conn, name)
//	if err != nil {
//		response.ResponseError(w, err.Error(), http.StatusInternalServerError)
//	}
//
//	fmt.Fprintln(w, "Success")
//}
//
//
//func getUserName(r *http.Request) string {
//	var data map[string]string
//	b, _ := io.ReadAll(r.Body)
//	json.Unmarshal(b, &data)
//
//	return data["name"]
//}
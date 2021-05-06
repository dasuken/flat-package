package response

import (
	"encoding/json"
	"github.com/noguchidaisuke/flat-package/product"
	"net/http"
)

func ResponseProduct(w http.ResponseWriter, product product.Product) {
	if  err := json.NewEncoder(w).Encode(product); err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func ResponseProducts(w http.ResponseWriter, products []product.Product) {
	if  err := json.NewEncoder(w).Encode(products); err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
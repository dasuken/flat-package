package response

import (
	"encoding/json"
	"github.com/noguchidaisuke/flat-package/user"
	"net/http"
)

func ResponseHelloWorld(w http.ResponseWriter, str string) {
	w.Write([]byte(str))
}

func ResponseUser(w http.ResponseWriter, user user.User) {
	if  err := json.NewEncoder(w).Encode(user); err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
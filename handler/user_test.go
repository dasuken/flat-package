package handler

import (
	"encoding/json"
	"github.com/noguchidaisuke/flat-package/user"
	"net/http"
	"testing"
)

func TestFindUserById(t *testing.T) {
	res, err := http.Get(srv.URL + "/users/1")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status code want: 200, got: %d", res.StatusCode)
	}

	var user user.User
	json.NewDecoder(res.Body).Decode(&user)

	if user.Name != "JOKER" {
		t.Fatalf("got: %s, want: JOKER", user.Name)
	}
}


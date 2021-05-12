package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

var srv *httptest.Server

func TestMain(m *testing.M) {
	setup()

	m.Run()
	srv.Close()
}


func setup() {
	fmt.Println("Setup")

	srv = httptest.NewServer(Router)
}
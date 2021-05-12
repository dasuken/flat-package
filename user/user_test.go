package user

import (
	"database/sql"
	"github.com/noguchidaisuke/flat-package/db"
	"testing"
)

var conn *sql.DB

func TestMain(m *testing.M) {
	setup()

	m.Run()
}

func setup() {
	conn, _ = db.NewTestDBConn()
}

func TestAddAndGetById(t *testing.T) {
	id, err := Add(conn, "Joker")
	if err != nil {
		t.Fatal(err)
	}

	got, err := GetById(conn, id)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != id && got.Name == "Joker" {
		t.Fatalf("want: Joker, got: %s", got.Name)
	}
}
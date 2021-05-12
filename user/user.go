package user

import (
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

// 引数にconnとると、毎回コネクション渡す必要がある。
// どうやってrepositoryと分けられるんだろ,,,
// 1. sql自体を自作
// interfaceに分ける
func GetById(conn *sql.DB, userId int) (user User, err error) {
	row := conn.QueryRow(`SELECT id, name FROM users WHERE id = ?`, userId)
	var id int
	var name string
	if err = row.Scan(&id, &name); err != nil {
		return
	}

	user.ID = id
	user.Name = name

	return
}

func Add(conn *sql.DB, name string) (id int, err error) {
	result, err := conn.Exec(`INSERT INTO users(name) VALUES (?)`, name)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	id = int(id64)
	return
}
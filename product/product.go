package product

import (
	"database/sql"
	"github.com/noguchidaisuke/flat-package/user"
)

type Product struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt string    `json:"created_at"`
	User      user.User `json:"user"`
}

func GetById(conn *sql.DB, productId int) (product Product, err error) {
	row := conn.QueryRow(`SELECT * FROM products WHERE id = ?`, productId)
	if err = row.Scan(&product.ID, &product.UserID, &product.Name, &product.Price, &product.CreatedAt); err != nil {
		return
	}

	return
}

func GetAll(conn *sql.DB) (products []Product, err error) {
	rows, _ := conn.Query(`SELECT * FROM products JOIN users ON products.user_id = users.id`)

	for rows.Next() {
		var product Product
		if err = rows.Scan(&product.ID, &product.UserID, &product.Name, &product.Price, &product.CreatedAt, &product.User.ID, &product.User.Name, &product.User.CreatedAt); err != nil {
			return
		}
		products = append(products, product)
	}

	return
}

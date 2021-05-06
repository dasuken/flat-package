CREATE DATABASE product;
USE product;

CREATE TABLE users (
    id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(name) VALUES ('JOKER');

CREATE TABLE products (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    name varchar(100) NOT NULL,
    price int NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY products_FK (user_id),
    CONSTRAINT products_FK FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
);

INSERT INTO products(user_id, name, price) VALUES (2, 'ball', 300);
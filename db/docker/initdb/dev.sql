DROP SCHEMA IF EXISTS product;
CREATE SCHEMA product;
USE product;

CREATE TABLE users (
    id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(name) VALUES ('JOKER'), ('Jamal');

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

INSERT INTO products(user_id, name, price) VALUES (2, 'ball', 300),  (2, 'food', 200),  (2, 'goal', 400);

CREATE TABLE tags (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100)
);

INSERT INTO tags(name) VALUES ('basket ball');

CREATE TABLE tags_products (
    tag_id int NOT NULL,
    product_id int NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

INSERT INTO tags_products VALUES (1, 1), (1, 2), (1, 3);
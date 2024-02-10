CREATE DATABASE IF NOT EXISTS supermarket;

USE supermarket;

CREATE TABLE IF NOT EXISTS products (
    id    UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name  STRING NOT NULL UNIQUE,
    price DECIMAL(10,2) NOT NULL,

    CONSTRAINT ok_price CHECK (price > 0)
);

CREATE TABLE IF NOT EXISTS orders (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS order_product (
    order_id   UUID NOT NULL REFERENCES orders (id),
    product_id UUID NOT NULL REFERENCES products (id),
    amount     INT  NOT NULL,

    PRIMARY KEY (order_id, product_id),
    CONSTRAINT ok_to_order CHECK (amount > 0)
);
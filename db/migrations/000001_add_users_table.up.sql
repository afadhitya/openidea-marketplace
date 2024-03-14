CREATE TABLE IF NOT EXISTS users 
(
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    product_sold_total INTEGER
);
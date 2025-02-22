CREATE TABLE goods
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(100) NOT NULL,
    price INT          NOT NULL CHECK (price > 0)
);
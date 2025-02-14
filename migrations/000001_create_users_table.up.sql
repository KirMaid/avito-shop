CREATE TABLE users
(
    id               SERIAL PRIMARY KEY,
    username         VARCHAR(50) UNIQUE NOT NULL,
    password         VARCHAR(255)       NOT NULL,
    balance          INT DEFAULT 1000 CHECK (balance >= 0)
--     token            VARCHAR(255)       NOT NULL,
--     token_expires_at TIMESTAMP          NOT NULL
);
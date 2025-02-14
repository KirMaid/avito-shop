CREATE TABLE inventories
(
    id       SERIAL PRIMARY KEY,
    user_id  INT REFERENCES users (id) ON DELETE CASCADE,
    type     VARCHAR(255) NOT NULL,
    quantity INT          NOT NULL
);
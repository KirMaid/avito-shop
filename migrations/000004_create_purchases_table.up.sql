CREATE TABLE purchases
(
    id         SERIAL PRIMARY KEY,
    user_id    INT REFERENCES users (id) ON DELETE CASCADE,
    merch_id   INT REFERENCES goods (id) ON DELETE CASCADE,
    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
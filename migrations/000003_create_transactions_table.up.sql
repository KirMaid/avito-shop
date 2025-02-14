CREATE TABLE transactions
(
    id          SERIAL PRIMARY KEY,
    sender_id   INT REFERENCES users (id) ON DELETE SET NULL,
    receiver_id INT REFERENCES users (id) ON DELETE SET NULL,
    amount      INT NOT NULL CHECK (amount > 0),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
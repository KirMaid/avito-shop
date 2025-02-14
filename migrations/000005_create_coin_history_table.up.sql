CREATE TABLE coin_history
(
    id             SERIAL PRIMARY KEY,
    user_id        INT REFERENCES users (id) ON DELETE CASCADE,
    change_amount  INT          NOT NULL,
    operation_type VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
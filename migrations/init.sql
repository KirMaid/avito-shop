CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    balance INT DEFAULT 1000 CHECK (balance >= 0)
);

CREATE TABLE merch (
   id SERIAL PRIMARY KEY,
   name VARCHAR(100) NOT NULL,
   price INT NOT NULL CHECK (price > 0)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    sender_id INT REFERENCES users(id) ON DELETE SET NULL,
    receiver_id INT REFERENCES users(id) ON DELETE SET NULL,
    amount INT NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- TODO Подумать насчёт количества
CREATE TABLE purchases (
   id SERIAL PRIMARY KEY,
   user_id INT REFERENCES users(id) ON DELETE CASCADE,
   merch_id INT REFERENCES merch(id) ON DELETE CASCADE,
--    quantity INT NOT NULL CHECK (quantity > 0),
   total_price INT NOT NULL CHECK (total_price > 0),
   Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE balance_history (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    change_amount INT NOT NULL,
    operation_type VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO merch (name, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500)



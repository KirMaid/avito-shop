-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     username VARCHAR(50) UNIQUE NOT NULL,
--     password VARCHAR(255) NOT NULL,
--     balance INT DEFAULT 1000 CHECK (balance >= 0)
-- );
--
-- CREATE TABLE goods (
--    id SERIAL PRIMARY KEY,
--    name VARCHAR(100) NOT NULL,
--    price INT NOT NULL CHECK (price > 0)
-- );
--
-- CREATE TABLE transactions (
--     id SERIAL PRIMARY KEY,
--     sender_id INT REFERENCES users(id) ON DELETE SET NULL,
--     receiver_id INT REFERENCES users(id) ON DELETE SET NULL,
--     amount INT NOT NULL CHECK (amount > 0),
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE TABLE purchases (
--    id SERIAL PRIMARY KEY,
--    user_id INT REFERENCES users(id) ON DELETE CASCADE,
--    merch_id INT REFERENCES goods(id) ON DELETE CASCADE,
--    Created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE TABLE coin_history (
--     id SERIAL PRIMARY KEY,
--     user_id INT REFERENCES users(id) ON DELETE CASCADE,
--     change_amount INT NOT NULL,
--     operation_type VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE TABLE inventories (
--     id SERIAL PRIMARY KEY,
--     user_id INT REFERENCES users(id) ON DELETE CASCADE,
--     type VARCHAR(255) NOT NULL,
--     quantity INT NOT NULL
-- );
--
--
-- INSERT INTO goods (name, price) VALUES
--     ('t-shirt', 80),
--     ('cup', 20),
--     ('book', 50),
--     ('pen', 10),
--     ('powerbank', 200),
--     ('hoody', 300),
--     ('umbrella', 200),
--     ('socks', 10),
--     ('wallet', 50),
--     ('pink-hoody', 500);





CREATE TABLE inventories (
     user_id INT NOT NULL,
     good_id INT NOT NULL,
     quantity INT NOT NULL CHECK (quantity >= 0),
     PRIMARY KEY (user_id, good_id),
     FOREIGN KEY (good_id) REFERENCES goods(id),
     FOREIGN KEY (user_id) REFERENCES users(id)
);
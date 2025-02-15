CREATE TABLE inventories (
     user_id INT NOT NULL,
     type VARCHAR(255) NOT NULL,
     quantity INT NOT NULL,
     PRIMARY KEY (user_id, type)
);
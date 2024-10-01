CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password TEXT
);

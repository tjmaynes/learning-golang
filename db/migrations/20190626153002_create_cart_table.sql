-- migrate:up
CREATE TABLE cart (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR (255) NOT NULL,
  price BIGINT NOT NULL,
  manufacturer VARCHAR (255) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS cart;

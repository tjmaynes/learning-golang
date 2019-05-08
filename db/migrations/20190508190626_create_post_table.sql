-- migrate:up
CREATE TABLE post (
  id SERIAL PRIMARY KEY,
  title VARCHAR (100) NOT NULL,
  content TEXT
);

-- migrate:down
DROP TABLE IF EXISTS post;

-- migrate:up
CREATE TABLE cart (
  id INT AUTO_INCREMENT,
  name VARCHAR (255) NOT NULL,
  price BIGINT NOT NULL,
  manufacturer VARCHAR (255) NOT NULL,
  PRIMARY KEY (id)
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

-- migrate:down
DROP TABLE IF EXISTS cart;

CREATE TABLE schema_migrations (version varchar(255) primary key);
CREATE TABLE cart (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR (255) NOT NULL,
  price BIGINT NOT NULL,
  manufacturer VARCHAR (255) NOT NULL
);
CREATE TABLE sqlite_sequence(name,seq);
-- Dbmate schema migrations
INSERT INTO schema_migrations (version) VALUES
  ('20190626153002');

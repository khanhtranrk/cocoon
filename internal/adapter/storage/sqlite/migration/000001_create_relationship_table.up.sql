CREATE TABLE citizens (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(255) NOT NULL,
  contact_gate VARCHAR(32) NOT NULL,
  registration_date INTEGER NOT NULL
);

CREATE UNIQUE INDEX citizens_index_on_id ON citizens (id);

CREATE TABLE contacts (
  citizen_id INTEGER NOT NULL,
  gate VARCHAR(32) NOT NULL,
);

CREATE UNIQUE INDEX contacts_index_on_citizen_id ON contacts (citizen_id);

CREATE TABLE contacts (
  citizen_id INTEGER NOT NULL,
);

CREATE UNIQUE INDEX contacts_index_on_citizen_id ON contacts (citizen_id);

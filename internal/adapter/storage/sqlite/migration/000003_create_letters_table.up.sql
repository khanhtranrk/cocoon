CREATE TABLE letters (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  typ INTEGER NOT NULL,
  code INTEGER NOT NULL,
  foreign_id INTEGER NOT NULL, 
  commit_time INTEGER NOT NULL,
  message BLOB NOT NULL,
  FOREIGN KEY (foreign_id) REFERENCES contacts (citizen_id)
);

CREATE UNIQUE INDEX letters_index_on_id ON letters (id);
CREATE UNIQUE INDEX letters_index_on_code_foreign_id_commit_time ON letters (code, foreign_id, commit_time);


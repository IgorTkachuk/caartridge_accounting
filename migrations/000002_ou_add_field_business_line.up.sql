CREATE TABLE business_line (
  id serial PRIMARY KEY,
  name TEXT
);

ALTER TABLE ou
    ADD COLUMN business_line_id INTEGER REFERENCES business_line(id) ON DELETE RESTRICT
;

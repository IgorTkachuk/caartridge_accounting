ALTER TABLE doc_type
    ADD COLUMN ctr_status_type_id
        INTEGER REFERENCES ctr_status_type(id)
            ON DELETE RESTRICT;
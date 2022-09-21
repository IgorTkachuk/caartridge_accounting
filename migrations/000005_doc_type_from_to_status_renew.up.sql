ALTER TABLE doc_type
    RENAME COLUMN ctr_status_type_id TO ctr_status_type_to;

ALTER TABLE doc_type
    ADD COLUMN ctr_status_type_from integer REFERENCES ctr_status_type(id)
        ON DELETE RESTRICT;
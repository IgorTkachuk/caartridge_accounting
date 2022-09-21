ALTER TABLE doc_type
    RENAME COLUMN ctr_status_type_to TO ctr_status_type_id;

ALTER TABLE doc_type
    DROP COLUMN ctr_status_type_from;
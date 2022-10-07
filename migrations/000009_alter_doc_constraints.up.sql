ALTER TABLE  doc_str RENAME TO doc_ctr;

ALTER TABLE public.doc_ctr DROP CONSTRAINT doc_str_ctr_id_fkey;

ALTER TABLE public.doc_ctr
    ADD CONSTRAINT doc_ctr_ctr_id_fkey FOREIGN KEY (ctr_id)
        REFERENCES public.ctr (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE;

ALTER TABLE public.doc_ctr DROP CONSTRAINT doc_str_doc_id_fkey;

ALTER TABLE public.doc_ctr
    ADD CONSTRAINT doc_ctr_doc_id_fkey FOREIGN KEY (doc_id)
        REFERENCES public.doc (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE;

ALTER TABLE public.ctr_status DROP CONSTRAINT ctr_status_doc_id_fkey;

ALTER TABLE public.ctr_status
    ADD CONSTRAINT ctr_status_doc_id_fkey FOREIGN KEY (doc_id)
        REFERENCES public.doc (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE;

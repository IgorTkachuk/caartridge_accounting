ALTER TABLE  doc_ctr RENAME TO doc_str;

ALTER TABLE public.doc_str DROP CONSTRAINT doc_ctr_ctr_id_fkey;

ALTER TABLE public.doc_str
    ADD CONSTRAINT doc_str_ctr_id_fkey FOREIGN KEY (ctr_id)
        REFERENCES public.ctr (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE RESTRICT ;

ALTER TABLE public.doc_str DROP CONSTRAINT doc_ctr_doc_id_fkey;

ALTER TABLE public.doc_str
    ADD CONSTRAINT doc_ctr_doc_id_fkey FOREIGN KEY (doc_id)
        REFERENCES public.doc (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE RESTRICT;

ALTER TABLE public.str_status DROP CONSTRAINT ctr_status_doc_id_fkey;

ALTER TABLE public.ctr_status
    ADD CONSTRAINT ctr_status_doc_id_fkey FOREIGN KEY (doc_id)
        REFERENCES public.doc (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE RESTRICT;

CREATE VIEW v_ctr (id, vendor, model, sn, status, doc_number, doc_date, employee, ou, business_line) AS
SELECT
    c.id as id,
    v.name as vendor, cm.name as model, c.sn, cst.name as status, d.id as doc_number,
    d.doc_date, empl.name as employee, o.name as ou, bl.name as business_line
FROM ctr c
         JOIN ctr_model cm ON c.model_id = cm.id
         JOIN vendor v ON cm.vendor_id = v.id
         JOIN ctr_status cs ON c.id = cs.ctr_id
         JOIN ctr_status_type cst ON cs.status_id = cst.id
         JOIN doc d ON  cs.doc_id = d.id
         LEFT JOIN employee empl ON d.employee_id = empl.id
         LEFT JOIN ou o ON empl.ou_id = o.id
         LEFT JOIN business_line bl ON o.business_line_id = bl.id
;
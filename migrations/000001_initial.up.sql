CREATE TABLE vendor (
                        id serial PRIMARY KEY,
                        name TEXT,
                        logo_url TEXT
);

CREATE TABLE prn_model (
                           id serial PRIMARY KEY,
                           name TEXT,
                           vendor_id INTEGER REFERENCES vendor(id),
                           image_url TEXT
);

CREATE TABLE ctr_model (
                           id SERIAL PRIMARY KEY,
                           name TEXT,
                           vendor_id INTEGER REFERENCES vendor(id),
                           image_url TEXT
);

CREATE TABLE ctr_supp_prn (
                              ctr_model_id INTEGER REFERENCES ctr_model(id) ON DELETE CASCADE,
                              prn_model_id INTEGER REFERENCES prn_model(id) ON DELETE RESTRICT
);

CREATE TABLE ou (
                    id SERIAL PRIMARY KEY,
                    name TEXT,
                    parent_id INTEGER REFERENCES ou(id) ON DELETE CASCADE
);

CREATE TABLE ctr (
                     id SERIAL PRIMARY KEY,
                     model_id INTEGER REFERENCES ctr_model(id) ON DELETE RESTRICT,
                     sn TEXT,
                     owner_id INTEGER REFERENCES ou(id) ON DELETE RESTRICT
);

CREATE TABLE employee (
                          id SERIAL PRIMARY KEY,
                          name TEXT,
                          ou_id INTEGER REFERENCES ou(id) ON DELETE RESTRICT
);

CREATE TABLE ctr_status_type (
                                 id SERIAL PRIMARY KEY,
                                 name TEXT
);

CREATE TABLE doc_type (
                          id SERIAL PRIMARY KEY,
                          name TEXT
);

CREATE TABLE decomissioning_cause (
                                      id SERIAL PRIMARY KEY,
                                      name TEXT
);

CREATE TABLE regenerate_type (
                                 id SERIAL PRIMARY KEY,
                                 name TEXT
);

CREATE TABLE usr (
                     id SERIAL PRIMARY KEY,
                     name TEXT,
                     pwd_hash TEXT
);

CREATE TABLE doc (
                     id SERIAL PRIMARY KEY,
                     type_id INTEGER REFERENCES doc_type(id) ON DELETE RESTRICT,
                     doc_date DATE,
                     employee_id INTEGER REFERENCES employee(id) on DELETE RESTRICT,
                     doc_owner_id INTEGER REFERENCES usr(id) ON DELETE RESTRICT,
                     decomissioning_couse_id INTEGER REFERENCES decomissioning_cause(id) ON DELETE RESTRICT,
                     ou_id INTEGER REFERENCES ou(id) ON DELETE RESTRICT,
                     sd_claim_number TEXT,
                     regenerate_type_id INTEGER references regenerate_type(id) on DELETE RESTRICT
);

CREATE TABLE doc_str(
                        doc_id INTEGER REFERENCES doc(id) ON DELETE RESTRICT,
                        ctr_id INTEGER REFERENCES ctr(id) ON DELETE RESTRICT
);

CREATE TABLE ctr_status(
                           ctr_id INTEGER REFERENCES ctr(id) ON DELETE RESTRICT,
                           status_from DATE,
                           status_id INTEGER REFERENCES ctr_status_type(id),
                           doc_id INTEGER REFERENCES doc(id) ON DELETE RESTRICT
);
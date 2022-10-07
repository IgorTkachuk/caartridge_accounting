package ctr_showcase

import "time"

type CtrShowcaseDTO struct {
	ID        int        `json:"id,omitempty"`
	Vendor    string     `json:"vendor,omitempty"`
	Model     string     `json:"model,omitempty"`
	Sn        string     `json:"sn,omitempty"`
	Status    string     `json:"status,omitempty"`
	DocNumber int        `json:"doc_number,omitempty"`
	DocDate   *time.Time `json:"doc_date,omitempty"`
	Employee  *string    `json:"employee,omitempty"`
	Ou        *string    `json:"ou,omitempty"`
	BaseLine  *string    `json:"base_line,omitempty"`
}

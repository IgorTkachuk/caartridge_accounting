package doc_type

type DocType struct {
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	CtrStatusTypeFrom *int   `json:"ctr_status_type_from"`
	CtrStatusTypeTo   *int   `json:"ctr_status_type_to"`
}

type CreateDocTypeDTO struct {
	Name              string `json:"name,omitempty"`
	CtrStatusTypeFrom *int   `json:"ctr_status_type_from"`
	CtrStatusTypeTo   *int   `json:"ctr_status_type_to"`
}

type UpdateDocTypeDTO struct {
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	CtrStatusTypeFrom *int   `json:"ctr_status_type_from"`
	CtrStatusTypeTo   *int   `json:"ctr_status_type_to"`
}

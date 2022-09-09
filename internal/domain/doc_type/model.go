package doc_type

type DocType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateDocTypeDTO struct {
	Name string `json:"name,omitempty"`
}

type UpdateDocTypeDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

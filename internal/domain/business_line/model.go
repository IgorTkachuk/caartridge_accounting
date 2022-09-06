package business_line

type BusinessLine struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UpdateBusinessLineDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateBusinessLineDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

package ou

type Ou struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	ParentId       *int   `json:"parent_id"`
	BusinessLineId *int   `json:"businessLineId"`
}

type UpdateOuDTO struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	ParentId       int    `json:"parent_id,omitempty"`
	BusinessLineId int    `json:"businessLineId"`
}

type CreateOuDTO struct {
	Name           string `json:"name,omitempty"`
	ParentId       *int   `json:"parent_id,omitempty"`
	BusinessLineId *int   `json:"businessLineId"`
}

package employee

type Employee struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	OuID int    `json:"ou_id,omitempty"`
}

type CreateEmployeeDTO struct {
	Name string `json:"name,omitempty"`
	OuId int    `json:"ou_id,omitempty"`
}

type UpdateEmployeeDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	OuId int    `json:"ou_id,omitempty"`
}

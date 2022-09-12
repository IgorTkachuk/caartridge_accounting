package decom_cause

type DecomCause struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateDecomCauseDTO struct {
	Name string `json:"name,omitempty"`
}

type UpdateDecomCauseDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

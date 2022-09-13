package cartridge_status_type

type CartridgeStatusType struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateCartridgeStatusTypeDTO struct {
	Name string `json:"name,omitempty"`
}

type UpdateCartridgeStatusTypeDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

package prnt

type Prn struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	VendorID int    `json:"vendor_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type CreatePrnDTO struct {
	Name     string `json:"name,omitempty"`
	VendorID int    `json:"vendor_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type DeletePrnDTO struct {
	ID int `json:"id"`
}

type UpdatePrnDTO struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	VendorID int    `json:"vendor_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

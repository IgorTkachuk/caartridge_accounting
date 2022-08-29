package prnmodel

type Prn struct {
	id       int    `json:"id,omitempty"`
	name     string `json:"name,omitempty"`
	vendorID int    `json:"vendor_id,omitempty"`
	imageUrl string `json:"image_url,omitempty"`
}

type CreatePrnDTO struct {
	name     string `json:"name,omitempty"`
	vendorID int    `json:"vendor_id,omitempty"`
	imageUrl string `json:"image_url,omitempty"`
}

type DeletePrnDTO struct {
	id int `json:"id"`
}

type UpdatePrnDTO struct {
	id       int    `json:"id,omitempty"`
	name     string `json:"name,omitempty"`
	vendorID int    `json:"vendor_id,omitempty"`
	imageUrl string `json:"image_url,omitempty"`
}
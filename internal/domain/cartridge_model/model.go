package cartridge_model

type CartridgeModel struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	VendorId int    `json:"vendor_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	SuppPrns []int  `json:"supp_prns"`
}

type CreateCartridgeModelDTO struct {
	Name     string `json:"name,omitempty"`
	VendorId int    `json:"vendor_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type UpdateCartridgeModelDTO struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	VendorId int    `json:"vendor_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	SuppPrns []int  `json:"supp_prns"`
}

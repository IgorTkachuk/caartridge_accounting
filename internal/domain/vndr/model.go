package vndr

type Vendor struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
}

type CreateVendorDTO struct {
	Name    string `json:"name,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
}

type UpdateVendorDTO struct {
	Name    string `json:"name,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
}

func NewVendor(dto CreateVendorDTO) Vendor {
	return Vendor{
		Name:    dto.Name,
		LogoUrl: dto.LogoUrl,
	}
}

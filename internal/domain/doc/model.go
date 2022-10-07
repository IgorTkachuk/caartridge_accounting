package doc

import "time"

type Doc struct {
	ID                     int        `json:"id,omitempty"`
	DocTypeId              int        `json:"doc_type_id,omitempty"`
	DocDate                *time.Time `json:"doc_date,omitempty"`
	EmployeeId             *int       `json:"employee_id,omitempty"`
	DocOwnerId             int        `json:"doc_owner_id,omitempty"`
	DecommissioningCauseId *int       `json:"decommissioning_cause_id,omitempty"`
	OuId                   *int       `json:"ou_id,omitempty,omitempty"`
	SdClaimNumber          *string    `json:"sd_claim_number,omitempty"`
	RegenerateTypeId       *int       `json:"regenerate_type_id,omitempty"`
	CreatedAt              time.Time  `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}

type CartridgeDTO struct {
	ID           int    `json:"id,omitempty"`
	ModelId      int    `json:"model_id,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
}

type CreateDocDTO struct {
	DocTypeId              int            `json:"doc_type_id,omitempty"`
	DocDate                *string        `json:"doc_date,omitempty"`
	EmployeeId             *int           `json:"employee_id,omitempty"`
	DocOwnerId             int            `json:"doc_owner_id,omitempty"`
	DecommissioningCauseId *int           `json:"decommissioning_cause_id,omitempty"`
	OuId                   *int           `json:"ou_id,omitempty"`
	SdClaimNumber          *string        `json:"sd_claim_number,omitempty"`
	RegenerateTypeId       *int           `json:"regenerate_type_id,omitempty"`
	CreatedAt              string         `json:"created_at,omitempty"`
	UpdatedAt              string         `json:"updated_at,omitempty"`
	Ctrs                   []CartridgeDTO `json:"ctrs"`
}

type UpdateDocDTO struct {
	ID                     int    `json:"id,omitempty"`
	DocTypeId              int    `json:"doc_type_id,omitempty"`
	DocDate                string `json:"doc_date,omitempty"`
	EmployeeId             *int   `json:"employee_id,omitempty"`
	DocOwnerId             int    `json:"doc_owner_id,omitempty"`
	DecommissioningCauseId *int   `json:"decommissioning_cause_id,omitempty"`
	OuId                   int    `json:"ou_id,omitempty"`
	SdClaimNumber          string `json:"sd_claim_number,omitempty"`
	RegenerateTypeId       *int   `json:"regenerate_type_id,omitempty"`
	CreatedAt              string `json:"created_at,omitempty"`
	UpdatedAt              string `json:"updated_at,omitempty"`
}

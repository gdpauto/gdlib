package types 

import (
	"time"
)

type Licence struct {
	Key        string 
	IsValid    bool
	ComputerID string
	StartDate  time.Time
	ExpiresAt  time.Time
}

type VerifyLicenceRequest struct {
	ComputerID string `form:"computer_id"`
	LicenceKey string `form:"licence_key"`
}

type VerifyLicenceResponse struct {
	Key        string    `json:"key"`
	IsValid    bool      `json:"is_valid"`
	ComputerID string	 `json:"computer_id"`
	StartDate  time.Time `json:"start_date"`
	ExpiresAt  time.Time `json:"expires_at"`
}

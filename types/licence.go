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
	Data []byte `form:"data"`
}

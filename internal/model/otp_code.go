package model

import (
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	OTP       string
	UserId    uuid.UUID
	ReadCount int8
	// Add Provider/Type property (phone, email)
	CreatedAt time.Time
	UpdatedAt time.Time
}

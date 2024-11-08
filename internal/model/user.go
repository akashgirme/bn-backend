package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	PhoneNumber      *string   `json:"phoneNumber,omitempty"`
	Email            *string   `json:"email,omitempty"`
	EmailProvider    string    `json:"-"`
	HasPrimaryNumber bool      `json:"hasPrimaryNumber"`
	HasEmail         bool      `josn:"hasEmail"`
	IsVerified       bool      `json:"isVerified"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

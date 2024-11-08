package model

import (
	"github.com/google/uuid"
)

type SignInMethod string

const (
	PhoneSignIn  SignInMethod = "PHONE"
	GoogleSignIn SignInMethod = "GOOGLE"
)

type PhoneSignInRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
}

type PhoneSignInResponse struct {
	UserID      uuid.UUID `json:"userId"`
	PhoneNumber string    `json:"phoneNumber"`
}

type OTPVerificationRequest struct {
	UserID      uuid.UUID `json:"userId" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required,e164"`
	OTP         string    `json:"otp" binding:"required,len=6"`
}

type GoogleSignInRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         *User  `json:"user"`
}

type AddEmailRequest struct {
	UserID uuid.UUID `json:"userId" binding:"required"`
	Email  string    `json:"email" binding:"required,email"`
}

// GoogleProfile represents the user profile information from google Oauth
type GoogleProfile struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	ProfileURL string `json:"profileUrl"`
}

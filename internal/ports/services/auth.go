package services

import (
	"context"

	"github.com/akashgirme/bn-backend/internal/model"
	"github.com/google/uuid"
)

type AuthService interface {
	// Google authentication
	GoogleOauth(ctx context.Context, email, password string) (*model.User, error)
	HandleGoogleCallback(ctx context.Context, code string) (*model.User, string, error)

	// Phone authentication
	RequestPhoneOTP(ctx context.Context, phoneNumber string) error
	VerifyPhoneOTP(ctx context.Context, phoneNumber, code string) (*model.User, string, error)

	// General auth methods
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, userID uuid.UUID) error
}

type OTPService interface {
	GenerateOTP(ctx context.Context, user *model.User) error
	VerifyOTP(ctx context.Context, otp string, user *model.User) (bool, error)
}

type TokenService interface {
	GenerateTokens(u *model.User) (accessToken, refreshToken string, err error)
	ValidateAccess(token string) (userID uuid.UUID, err error)
	ValidateRefresh(token string) (userID uuid.UUID, err error)
}

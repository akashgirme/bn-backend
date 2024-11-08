package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/akashgirme/bn-backend/internal/model"
	services "github.com/akashgirme/bn-backend/internal/ports/services"
	"github.com/google/uuid"
)

// Need Logger, cache,
type AuthService struct {
	userService  services.UserService
	otpService   services.OTPService
	googleAuth   GoogleAuth
	tokenService services.TokenService
}

func NewAuthService(
	userService services.UserService,
	otpService services.OTPService,
	googleAuth GoogleAuth,
	tokenService services.TokenService,
) *AuthService {
	return &AuthService{
		userService:  userService,
		otpService:   otpService,
		googleAuth:   googleAuth,
		tokenService: tokenService,
	}
}
func (as *AuthService) SignInWithPhone(ctx context.Context, phone string) (*model.PhoneSignInResponse, error) {

	var user *model.User

	// Check if user exists
	user, err := as.userService.GetByPhone(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by phone: %w", err)
	}

	// Create user if doesn't exist
	if user == nil {
		user := &model.User{
			ID:          uuid.New(),
			PhoneNumber: &phone,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if _, err := as.userService.CreateWithPhone(ctx, user); err != nil {
			return nil, fmt.Errorf("failed to save new user: %w", err)
		}
	}

	if err := as.otpService.GenerateOTP(ctx, user); err != nil {
		return nil, fmt.Errorf("error while sending otp: %w", err)
	}

	return &model.PhoneSignInResponse{
		UserID:      user.ID,
		PhoneNumber: phone,
	}, nil
}

func (as *AuthService) VerifyOTP(ctx context.Context, req *model.OTPVerificationRequest) (*model.AuthResponse, error) {
	// Get user
	user, err := as.userService.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Verify OTP
	if _, err := as.otpService.VerifyOTP(ctx, req.OTP, user); err != nil {
		return nil, fmt.Errorf("invalid OTP: %w", err)
	}

	return as.authenticate(user)
}

func (as *AuthService) SignInWithGoogle(ctx context.Context, token string) (*model.AuthResponse, error) {
	// Verify Google token and get profile
	// Find user with email if not exist create new
	// authenticate()

	profile, err := s.googleAuth.VerifyToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("invalid Google token: %w", err)
	}

	// Check if user exists by email
	user, err := s.userRepo.GetByEmail(ctx, profile.Email)
	if err != nil && !errors.Is(err, repositories.ErrNotFound) {
		return nil, fmt.Errorf("failed to check user: %w", err)
	}

	if errors.Is(err, repositories.ErrNotFound) {
		// Create new user
		user = &model.User{
			ID:        uuid.New(),
			Email:     &profile.Email,
			Name:      &profile.Name,
			Picture:   &profile.Picture,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := s.userRepo.Create(ctx, user); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		// Update existing user's Google info if needed
		user.Name = &profile.Name
		user.Picture = &profile.Picture
		user.UpdatedAt = time.Now()
		if err := s.userRepo.Update(ctx, user); err != nil {
			return nil, fmt.Errorf("failed to update user: %w", err)
		}
	}

	return as.authenticate(user)
}

func (as *AuthService) authenticate(u *model.User) (*model.AuthResponse, error) {
	// Generate tokens
	accessToken, refreshToken, err := as.tokenService.GenerateTokens(u)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return &model.AuthResponse{
		User:         u,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

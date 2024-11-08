package user

import (
	"context"
	"fmt"
	"time"

	"github.com/akashgirme/bn-backend/internal/domain"
	"github.com/xtgo/uuid"
)

// CreateUserWithPhone()
// createUserWithEmail()

// Create a separate service for google Oauth related stuff.

// store otp in cache for fast retrival don't need long persistance.

// If link email is through google account
// Then validate the request via token and extract the email from response
// Then call LinkEmail() method with provider as `google`.

// Provider is required because have to store the provider of email address for future use.

// If provider is normal email then first verify the email.

// Don't create a separate method for resending the opt, call same methods again with same req.
// `LinkEmail` && `Link Phone`.

// Just need method to verify the user.

// Link phone number is simple and straight forward.

// Link Email ( provider 'Google' | 'Email', email string)
// Link Phone ( phone string)
// Change Number (phone)

// Other stuff in user service

// User profile
// profile image, name/nickname, saved addresses, emergency contacts, secondary contact.
// roles, driver, customer, verification if driver (digilocker or can directly verify the licence + bike)
//

func (as *AuthService) AddEmailWithGoogle(ctx context.Context, userID uuid.UUID, token string) (*domain.User, error) {
	// Verify user exists
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Verify Google token and get profile
	profile, err := s.googleAuth.VerifyToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("invalid Google token: %w", err)
	}

	// Check if email is already used
	if _, err := s.userRepo.GetByEmail(ctx, profile.Email); err == nil {
		return nil, fmt.Errorf("email already in use")
	}

	// Update user with email and Google info
	user.Email = &profile.Email
	user.Name = &profile.Name
	user.Picture = &profile.Picture
	user.UpdatedAt = time.Now()

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

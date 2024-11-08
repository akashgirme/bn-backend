package services

import (
	"context"

	"github.com/akashgirme/bn-backend/internal/model"
	"github.com/google/uuid"
)

type UserService interface {
	// Create User
	CreateWithPhone(ctx context.Context, user *model.User) (*model.User, error)
	// Get user
	GetByID(ctx context.Context, userId uuid.UUID) (*model.User, error)
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)

	// Google
	LinkGoogle()

	//Email
	LinkEmail()
	ResendLinkEmailOTP()
	VerifyLinkEmailOTP()

	//Phone
	LinkPhone()
	ResendLinkPhoneOTP()
	VerifyLinkPhoneOTP()

	ChangePhone()
	ResendChangePhoneOTP()
	VerifyChangePhoneOTP()
}

package store

import (
	"context"

	"github.com/akashgirme/bn-backend/internal/model"
	"github.com/google/uuid"
)

type Store interface {
	User() UserStore
	OTP() OTPStore
}

type UserStore interface {
	CreateWithEmail(ctx context.Context, user *model.User) error
	CreateWithPhone(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, userId string) (*model.User, error)
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type OTPStore interface {
	Save(ctx context.Context, otp *model.OTP) (*model.OTP, error)
	Get(ctx context.Context, userId uuid.UUID) (*model.OTP, error)
}

package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/akashgirme/bn-backend/internal/model"
	"github.com/akashgirme/bn-backend/internal/ports/services"
	"github.com/akashgirme/bn-backend/internal/ports/store"
)

// Need Logger, cache,
type OTPService struct {
	otpRepo    store.OTPStore
	smsService services.SMSService
}

func NewOTPService(
	otpRepo store.OTPStore,
	smsService services.SMSService,
) *OTPService {
	return &OTPService{
		otpRepo:    otpRepo,
		smsService: smsService,
	}
}

func (s *OTPService) GenerateOTP(ctx context.Context, u *model.User) error {

	otp, err := generateOTP()

	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	payload := &model.OTP{
		OTP:       otp,
		UserId:    u.ID,
		ReadCount: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := s.otpRepo.Save(ctx, payload); err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	if err := s.smsService.SendOTP(ctx, *u.PhoneNumber, payload.OTP); err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	return nil

}

func (s *OTPService) VerifyOTP(ctx context.Context, otp string, u *model.User) (bool, error) {

	return true, nil
}

func generateOTP() (string, error) {
	return "123456", nil
}

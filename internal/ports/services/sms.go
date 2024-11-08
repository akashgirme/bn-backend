package services

import "context"

type SMSService interface {
	SendOTP(ctx context.Context, code, phone string) error
}

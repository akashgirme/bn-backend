package config

import (
	"time"

	"github.com/akashgirme/bn-backend/config/env"
)

type Config struct {
	Addr        string
	DB          DBConfig
	Env         string
	APIURL      string
	Mail        MailConfig
	FrontendURL string
	JWTToken    JWTTokenConfig
	Redis       RedisConfig
	Twilio      TwilioConfig
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type RedisConfig struct {
	RedisURL string
	Enabled  bool
}

type JWTTokenConfig struct {
	AuthToken    TokenConfig
	RefreshToken TokenConfig
}

type TokenConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
	Aud    string
}

type MailConfig struct {
	AWSSES    AWSSESConfig
	FromEmail string
	Exp       time.Duration
}

type AWSSESConfig struct {
	AWSRegion string
}

type TwilioConfig struct {
	TwilioAccountSID string
	TwilioAuthToken  string
	VerifyServiceSID string
}

func Load() (*Config, error) {
	return &Config{
		Addr:        env.GetString("ADDR", ":8080"),
		APIURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		FrontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		DB: DBConfig{
			Addr:         env.GetString("DB_URL", "postgres://admin:adminpassword@localhost:5432?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		Redis: RedisConfig{
			RedisURL: env.GetString("REDIS_URL", "redis://localhost:6379"),
			Enabled:  env.GetBool("REDIS_ENABLED", false),
		},
		Env: env.GetString("ENV", "development"),
		Mail: MailConfig{
			Exp:       time.Hour * 24 * 3, // 3 days
			FromEmail: env.GetString("SES_EMAIL_SENDER", ""),
			AWSSES: AWSSESConfig{
				AWSRegion: env.GetString("AWS_REGION", "ap-south-1"),
			},
		},
		Twilio: TwilioConfig{
			TwilioAccountSID: env.GetString("TWILIO_ACCOUNT_SID", ""),
			TwilioAuthToken:  env.GetString("TWILIO_AUTH_TOKEN", ""),
			VerifyServiceSID: env.GetString("TWILIO_VERIFY_SERVICE_SID", ""),
		},
		JWTToken: JWTTokenConfig{
			AuthToken: TokenConfig{
				Secret: env.GetString("AUTH_TOKEN_SECRET", "fallback-jwt-secret"),
				// Exp:    time.Hour * 24 * 3, // 3 days
				Exp: time.Minute * 15, // 15 Minutes
				Iss: "bn-backend",
				Aud: "bn",
			},
			RefreshToken: TokenConfig{
				Secret: env.GetString("REFRESH_TOKEN_SECRET", "fallback-jwt-secret"),
				// Exp:    time.Hour * 24 * 3, // 3 days
				Exp: time.Hour * 24 * 15, // 15 days
				Iss: "bn-backend",
				Aud: "bn",
			},
		},
	}, nil
}

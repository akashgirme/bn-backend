package auth

import (
	"fmt"
	"time"

	"github.com/akashgirme/bn-backend/config"
	"github.com/akashgirme/bn-backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTAuthenticator struct {
	config config.Config
}

func NewJWTAuthenticator(config config.Config) *JWTAuthenticator {
	return &JWTAuthenticator{config: config}
}

func (a *JWTAuthenticator) GenerateTokens(u *model.User) (accessToken, refreshToken string, err error) {
	authclaims := jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(a.config.JWTToken.AuthToken.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.config.JWTToken.AuthToken.Iss,
		"aud": a.config.JWTToken.AuthToken.Aud,
	}

	refreshclaims := jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(a.config.JWTToken.RefreshToken.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": a.config.JWTToken.RefreshToken.Iss,
		"aud": a.config.JWTToken.RefreshToken.Aud,
	}

	accessToken, err = a.generateToken(authclaims, a.config.JWTToken.AuthToken.Secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err = a.generateToken(refreshclaims, a.config.JWTToken.RefreshToken.Secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (a *JWTAuthenticator) generateToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *JWTAuthenticator) ValidateAccess(token string) (userID uuid.UUID, err error) {
	jwtToken, err := a.validateToken(token, a.config.JWTToken.AuthToken.Secret, a.config.JWTToken.AuthToken.Iss, a.config.JWTToken.AuthToken.Aud)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to validate access token: %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("subject claim is not a string")
	}

	userID, err = uuid.Parse(sub) // Convert the subject claim to UUID
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse user ID from claims: %w", err)
	}

	return userID, nil
}

func (a *JWTAuthenticator) ValidateRefresh(token string) (userID uuid.UUID, err error) {
	jwtToken, err := a.validateToken(token, a.config.JWTToken.RefreshToken.Secret, a.config.JWTToken.RefreshToken.Iss, a.config.JWTToken.RefreshToken.Aud)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to validate access token: %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("subject claim is not a string")
	}

	userID, err = uuid.Parse(sub) // Convert the subject claim to UUID
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse user ID from claims: %w", err)
	}

	return userID, nil
}

func (a *JWTAuthenticator) validateToken(token string, secret, issuer, audience string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(audience),
		jwt.WithIssuer(issuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}

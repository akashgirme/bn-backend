package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/akashgirme/bn-backend/internal/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleAuth defines the interface for Google authentication
type GoogleAuth interface {
	GetProfile(token string) (*model.GoogleProfile, error)
}

// GoogleAuthImpl implements GoogleAuth interface
type GoogleAuthImpl struct {
	config *oauth2.Config
}

// NewGoogleAuth creates a new instance of GoogleAuth
func NewGoogleAuth(clientID, clientSecret, redirectURL string) GoogleAuth {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &GoogleAuthImpl{
		config: config,
	}
}

// GetProfile fetches the user profile using the access token
func (g *GoogleAuthImpl) GetProfile(accessToken string) (*Profile, error) {
	// Verify that the token is not empty
	if accessToken == "" {
		return nil, fmt.Errorf("access token is required")
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Google's userinfo endpoint
	userinfoURL := "https://www.googleapis.com/oauth2/v3/userinfo"
	req, err := http.NewRequest("GET", userinfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the access token to the Authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: status code %d", resp.StatusCode)
	}

	// Parse the response
	var result struct {
		Email      string `json:"email"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Picture    string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Create and return the profile
	return &Profile{
		Email:      result.Email,
		FirstName:  result.GivenName,
		LastName:   result.FamilyName,
		ProfileURL: result.Picture,
	}, nil
}

// ValidateToken verifies if the token is valid
func (g *GoogleAuthImpl) ValidateToken(accessToken string) error {
	// Google's token info endpoint
	tokenInfoURL := "https://oauth2.googleapis.com/tokeninfo"

	// Create the request URL with the access token
	reqURL := fmt.Sprintf("%s?access_token=%s", tokenInfoURL, url.QueryEscape(accessToken))

	// Make the request
	resp, err := http.Get(reqURL)
	if err != nil {
		return fmt.Errorf("failed to validate token: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid token: status code %d", resp.StatusCode)
	}

	return nil
}

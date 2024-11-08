package handler

import (
	"net/http"

	"github.com/akashgirme/bn-backend/internal/common/jsonutil"
	"github.com/akashgirme/bn-backend/internal/core/auth"
	"github.com/akashgirme/bn-backend/internal/domain"
)

type AuthHandler struct {
	authService *auth.Service
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// SignInWithPhone handles phone sign-in requests.
func (h *AuthHandler) SignInWithPhone(w http.ResponseWriter, r *http.Request) {
	var payload domain.PhoneSignInRequest

	if err := jsonutil.ReadJSON(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if err := json.Validate.Struct(payload); err != nil {
	// 	errors.BadRequestResponse(w, r, err)
	// 	return
	// }

	resp, err := h.authService.SignInWithPhone(r.Context(), payload.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonutil.JsonResponse(w, http.StatusOK, resp)
}

// VerifyOTP handles OTP verification requests.
func (h *AuthHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var payload domain.OTPVerificationRequest

	if err := jsonutil.ReadJSON(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if err := json.Validate.Struct(payload); err != nil {
	// 	errors.BadRequestResponse(w, r, err)
	// 	return
	// }

	resp, err := h.authService.VerifyOTP(r.Context(), &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	jsonutil.JsonResponse(w, http.StatusOK, resp)

}

// SignInWithGoogle handles Google sign-in requests.
func (h *AuthHandler) SignInWithGoogle(w http.ResponseWriter, r *http.Request) {
	var payload domain.GoogleSignInRequest

	if err := jsonutil.ReadJSON(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if err := json.Validate.Struct(payload); err != nil {
	// 	errors.BadRequestResponse(w, r, err)
	// 	return
	// }

	resp, err := h.authService.SignInWithGoogle(r.Context(), payload.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	jsonutil.JsonResponse(w, http.StatusOK, resp)
}

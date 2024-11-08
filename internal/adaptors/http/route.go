package route

import "github.com/go-chi/chi"

// SetupRoutes sets up the routes for the AuthHandler.
func (h *AuthHandler) SetupRoutes(r chi.Router) {
	r.Post("/auth/signin/phone", h.SignInWithPhone)
	r.Post("/auth/verify-otp", h.VerifyOTP)
	r.Post("/auth/signin/google", h.SignInWithGoogle)
}

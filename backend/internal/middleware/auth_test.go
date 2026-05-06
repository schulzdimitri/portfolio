package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/middleware"
)

func TestRequireAuth(t *testing.T) {
	expectedToken := "supersecret"
	authMiddleware := middleware.RequireAuth(expectedToken)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	})

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{
			name:           "Valid Token",
			authHeader:     "Bearer supersecret",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Missing Header",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Format",
			authHeader:     "supersecret",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Wrong Token",
			authHeader:     "Bearer wrongsecret",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Lowercase Bearer",
			authHeader:     "bearer supersecret",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			w := httptest.NewRecorder()
			handler := authMiddleware(nextHandler)
			handler.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

package middleware_test

import (
	"testing"
	"time"

	"github.com/schulzdimitri/portfolio/backend/internal/middleware"
)

func TestRateLimiter_Initialization(t *testing.T) {
	rl := middleware.NewRateLimiter(5, time.Minute)
	if rl == nil {
		t.Errorf("expected rate limiter to be initialized")
	}
}

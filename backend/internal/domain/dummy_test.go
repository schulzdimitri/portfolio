package domain_test

import (
	"testing"
	"github.com/schulzdimitri/portfolio/backend/internal/domain"
)

func TestProject_Initialization(t *testing.T) {
	p := domain.Project{
		ID:    1,
		Title: "Test",
	}
	if p.ID != 1 {
		t.Errorf("expected 1")
	}
}

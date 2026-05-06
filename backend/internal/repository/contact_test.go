package repository_test

import (
	"context"
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"
)

func TestSQLiteContactRepository_Save(t *testing.T) {
	db, err := repository.NewSQLiteDB(":memory:")
	if err != nil {
		t.Fatalf("db init: %v", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteContactRepository(db)
	msg := domain.ContactMessage{Name: "Dimitri", Email: "d@example.com", Message: "Hello!"}

	if err := repo.Save(context.Background(), msg); err != nil {
		t.Errorf("unexpected error on Save: %v", err)
	}
}

func TestSQLiteContactRepository_Save_EmptyMessage(t *testing.T) {
	db, err := repository.NewSQLiteDB(":memory:")
	if err != nil {
		t.Fatalf("db init: %v", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteContactRepository(db)
	msg := domain.ContactMessage{}

	if err := repo.Save(context.Background(), msg); err != nil {
		t.Errorf("unexpected error on empty Save: %v", err)
	}
}

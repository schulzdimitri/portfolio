package repository_test

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"
)

func setupExperienceDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	query := `
	CREATE TABLE experiences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		company TEXT NOT NULL,
		role TEXT NOT NULL,
		period TEXT NOT NULL,
		duties TEXT NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	return db
}

func TestSQLiteExperienceRepository(t *testing.T) {
	db := setupExperienceDB(t)
	defer db.Close()

	repo := repository.NewSQLiteExperienceRepository(db)

	t.Run("Insert and GetAll", func(t *testing.T) {
		exp := &domain.Experience{
			Company: "Company A",
			Role:    "Backend Engineer",
			Period:  "2022-2024",
			Duties:  []string{"Task 1", "Task 2"},
		}

		err := repo.Insert(exp)
		if err != nil {
			t.Fatalf("expected no error on insert, got: %v", err)
		}
		if exp.ID == 0 {
			t.Error("expected ID to be set after insert")
		}

		experiences, err := repo.GetAll()
		if err != nil {
			t.Fatalf("expected no error on GetAll, got: %v", err)
		}
		if len(experiences) != 1 {
			t.Fatalf("expected 1 experience, got %d", len(experiences))
		}

		if experiences[0].Company != "Company A" {
			t.Errorf("expected Company A, got %s", experiences[0].Company)
		}
		if len(experiences[0].Duties) != 2 {
			t.Errorf("expected 2 duties, got %d", len(experiences[0].Duties))
		}
	})

	t.Run("Count", func(t *testing.T) {
		count, err := repo.Count()
		if err != nil {
			t.Fatalf("expected no error on Count, got: %v", err)
		}
		if count != 1 {
			t.Errorf("expected count 1, got %d", count)
		}
	})
}

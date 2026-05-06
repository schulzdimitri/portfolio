package repository_test

import (
	"testing"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"
)

func TestSQLiteProjectRepository_InsertAndGetAll(t *testing.T) {
	db, err := repository.NewSQLiteDB(":memory:")
	if err != nil {
		t.Fatalf("db init: %v", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteProjectRepository(db)

	p := domain.Project{
		Title:       "Test Project",
		Description: "A great project",
		Github:      "https://github.com/test/repo",
		Tags:        []string{"Go", "Test"},
	}

	err = repo.Insert(&p)
	if err != nil {
		t.Fatalf("unexpected error on Insert: %v", err)
	}

	if p.ID == 0 {
		t.Errorf("expected ID to be set, got 0")
	}

	projects, err := repo.GetAll()
	if err != nil {
		t.Fatalf("unexpected error on GetAll: %v", err)
	}

	if len(projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(projects))
	}

	fetched := projects[0]
	if fetched.Title != p.Title || fetched.Description != p.Description || fetched.Github != p.Github {
		t.Errorf("fetched project data mismatch")
	}

	if len(fetched.Tags) != 2 || fetched.Tags[0] != "Go" || fetched.Tags[1] != "Test" {
		t.Errorf("fetched tags mismatch")
	}
}

func TestSQLiteProjectRepository_Count(t *testing.T) {
	db, err := repository.NewSQLiteDB(":memory:")
	if err != nil {
		t.Fatalf("db init: %v", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteProjectRepository(db)

	count, err := repo.Count()
	if err != nil {
		t.Fatalf("unexpected error on Count: %v", err)
	}

	if count != 0 {
		t.Errorf("expected 0 count initially, got %d", count)
	}

	err = repo.Insert(&domain.Project{Title: "Proj1", Tags: []string{}})
	if err != nil {
		t.Fatalf("failed to insert: %v", err)
	}

	count, err = repo.Count()
	if err != nil {
		t.Fatalf("unexpected error on Count after insert: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 count, got %d", count)
	}
}

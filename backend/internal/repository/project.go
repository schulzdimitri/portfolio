package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
)

type ProjectRepository interface {
	Insert(project *domain.Project) error
	GetAll() ([]domain.Project, error)
	Count() (int, error)
}

type SQLiteProjectRepository struct {
	db *sql.DB
}

func NewSQLiteProjectRepository(db *sql.DB) *SQLiteProjectRepository {
	return &SQLiteProjectRepository{db: db}
}

// Insert adds a new project to the database
func (r *SQLiteProjectRepository) Insert(project *domain.Project) error {
	tagsJSON, err := json.Marshal(project.Tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %w", err)
	}

	query := `INSERT INTO projects (title, description, github, tags) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, project.Title, project.Description, project.Github, string(tagsJSON))
	if err != nil {
		return fmt.Errorf("failed to insert project: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	project.ID = int(id)
	return nil
}

// GetAll returns all projects from the database
func (r *SQLiteProjectRepository) GetAll() ([]domain.Project, error) {
	query := `SELECT id, title, description, github, tags FROM projects ORDER BY id ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []domain.Project
	for rows.Next() {
		var p domain.Project
		var tagsJSON string
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Github, &tagsJSON); err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}

		if err := json.Unmarshal([]byte(tagsJSON), &p.Tags); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tags: %w", err)
		}

		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating projects: %w", err)
	}

	// Return empty slice instead of nil for JSON serialization consistency
	if projects == nil {
		projects = make([]domain.Project, 0)
	}

	return projects, nil
}

// Count returns the total number of projects
func (r *SQLiteProjectRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM projects`).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to count projects: %w", err)
	}
	return count, nil
}

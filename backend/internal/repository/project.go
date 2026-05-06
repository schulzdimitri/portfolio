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
	Delete(id int) error
	Update(id int, project *domain.Project) error
}

func NewSQLiteProjectRepository(db *sql.DB) *SQLiteProjectRepository {
	return &SQLiteProjectRepository{db: db}
}

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

	if projects == nil {
		projects = make([]domain.Project, 0)
	}

	return projects, nil
}

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

func (r *SQLiteProjectRepository) Delete(id int) error {
	query := `DELETE FROM projects WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("project not found")
	}

	return nil
}

func (r *SQLiteProjectRepository) Update(id int, project *domain.Project) error {
	tagsJSON, err := json.Marshal(project.Tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %w", err)
	}

	query := `UPDATE projects SET title = ?, description = ?, github = ?, tags = ? WHERE id = ?`
	result, err := r.db.Exec(query, project.Title, project.Description, project.Github, string(tagsJSON), id)
	if err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("project not found")
	}

	project.ID = id
	return nil
}

type SQLiteProjectRepository struct {
	db *sql.DB
}

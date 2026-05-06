package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
)

type ExperienceRepository interface {
	Insert(exp *domain.Experience) error
	GetAll() ([]domain.Experience, error)
	Count() (int, error)
	Delete(id int) error
	Update(id int, exp *domain.Experience) error
}

type SQLiteExperienceRepository struct {
	db *sql.DB
}

func NewSQLiteExperienceRepository(db *sql.DB) *SQLiteExperienceRepository {
	return &SQLiteExperienceRepository{db: db}
}

func (r *SQLiteExperienceRepository) Insert(exp *domain.Experience) error {
	dutiesJSON, err := json.Marshal(exp.Duties)
	if err != nil {
		return fmt.Errorf("failed to marshal duties: %w", err)
	}

	query := `INSERT INTO experiences (company, role, period, duties) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, exp.Company, exp.Role, exp.Period, string(dutiesJSON))
	if err != nil {
		return fmt.Errorf("failed to insert experience: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	exp.ID = int(id)
	return nil
}

func (r *SQLiteExperienceRepository) GetAll() ([]domain.Experience, error) {
	query := `SELECT id, company, role, period, duties FROM experiences ORDER BY id ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query experiences: %w", err)
	}
	defer rows.Close()

	var experiences []domain.Experience
	for rows.Next() {
		var e domain.Experience
		var dutiesJSON string
		if err := rows.Scan(&e.ID, &e.Company, &e.Role, &e.Period, &dutiesJSON); err != nil {
			return nil, fmt.Errorf("failed to scan experience: %w", err)
		}

		if err := json.Unmarshal([]byte(dutiesJSON), &e.Duties); err != nil {
			return nil, fmt.Errorf("failed to unmarshal duties: %w", err)
		}

		experiences = append(experiences, e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating experiences: %w", err)
	}

	if experiences == nil {
		experiences = make([]domain.Experience, 0)
	}

	return experiences, nil
}

func (r *SQLiteExperienceRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM experiences`).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to count experiences: %w", err)
	}
	return count, nil
}
func (r *SQLiteExperienceRepository) Delete(id int) error {
	query := `DELETE FROM experiences WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete experience: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("experience not found")
	}
	return nil
}

func (r *SQLiteExperienceRepository) Update(id int, exp *domain.Experience) error {
	dutiesJSON, err := json.Marshal(exp.Duties)
	if err != nil {
		return fmt.Errorf("failed to marshal duties: %w", err)
	}

	query := `UPDATE experiences SET company = ?, role = ?, period = ?, duties = ? WHERE id = ?`
	result, err := r.db.Exec(query, exp.Company, exp.Role, exp.Period, string(dutiesJSON), id)
	if err != nil {
		return fmt.Errorf("failed to update experience: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("experience not found")
	}

	exp.ID = id
	return nil
}

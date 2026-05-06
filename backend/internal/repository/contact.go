package repository

import (
	"context"
	"database/sql"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
)

type ContactRepository interface {
	Save(ctx context.Context, msg domain.ContactMessage) error
}

type SQLiteContactRepository struct {
	db *sql.DB
}

func NewSQLiteContactRepository(db *sql.DB) *SQLiteContactRepository {
	return &SQLiteContactRepository{db: db}
}

func (r *SQLiteContactRepository) Save(ctx context.Context, msg domain.ContactMessage) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO contacts (name, email, message) VALUES (?, ?, ?)`,
		msg.Name, msg.Email, msg.Message,
	)
	return err
}

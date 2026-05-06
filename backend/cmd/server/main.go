package main

import (
	_ "embed"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/schulzdimitri/portfolio/backend/internal/domain"
	"github.com/schulzdimitri/portfolio/backend/internal/handler"
	"github.com/schulzdimitri/portfolio/backend/internal/middleware"
	"github.com/schulzdimitri/portfolio/backend/internal/repository"
	"github.com/schulzdimitri/portfolio/backend/internal/sender"
)

func main() {
	port := getenv("PORT", "8080")
	dbPath := getenv("DB_PATH", "portfolio.db")
	allowedOrigin := getenv("ALLOWED_ORIGIN", "*")

	db, err := repository.NewSQLiteDB(dbPath)
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	contactRepo := repository.NewSQLiteContactRepository(db)
	projectRepo := repository.NewSQLiteProjectRepository(db)

	if err := seedProjects(projectRepo); err != nil {
		slog.Warn("could not seed projects", "error", err)
	}

	emailSender := sender.NewSMTP(sender.SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		User:     os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASSWORD"),
		To:       os.Getenv("CONTACT_TO_EMAIL"),
	})

	contactLimiter := middleware.NewRateLimiter(5, time.Minute)
	projectHandler := handler.NewProjectHandler(projectRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handler.Health)
	mux.Handle("/api/contact", contactLimiter.Middleware(handler.ContactHandler(contactRepo, emailSender)))
	mux.HandleFunc("/api/projects", projectHandler.GetProjects)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      middleware.CORS(allowedOrigin, mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info("server starting", "port", port, "db", dbPath)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

var seedData []byte

func seedProjects(repo repository.ProjectRepository) error {
	count, err := repo.Count()
	if err != nil {
		return err
	}
	if count > 0 {
		slog.Info("projects already seeded", "count", count)
		return nil
	}

	var portfolioData struct {
		Projects []domain.Project `json:"projects"`
	}

	if err := json.Unmarshal(seedData, &portfolioData); err != nil {
		return err
	}

	inserted := 0
	for _, p := range portfolioData.Projects {
		if err := repo.Insert(&p); err != nil {
			slog.Error("failed to seed project", "title", p.Title, "error", err)
			continue
		}
		inserted++
	}

	slog.Info("seeded projects", "inserted", inserted)
	return nil
}

package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"


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
	experienceRepo := repository.NewSQLiteExperienceRepository(db)

	emailSender := sender.NewSMTP(sender.SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		User:     os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASSWORD"),
		To:       os.Getenv("CONTACT_TO_EMAIL"),
	})

	contactLimiter := middleware.NewRateLimiter(5, time.Minute)
	projectHandler := handler.NewProjectHandler(projectRepo)
	experienceHandler := handler.NewExperienceHandler(experienceRepo)

	adminToken := getenv("ADMIN_TOKEN", "supersecret123")
	authMiddleware := middleware.RequireAuth(adminToken)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handler.Health)
	mux.Handle("/api/contact", contactLimiter.Middleware(handler.ContactHandler(contactRepo, emailSender)))
	
	mux.HandleFunc("GET /api/projects", projectHandler.GetProjects)
	mux.Handle("POST /api/projects", authMiddleware(http.HandlerFunc(projectHandler.CreateProject)))
        mux.Handle("PUT /api/projects/{id}", authMiddleware(http.HandlerFunc(projectHandler.UpdateProject)))
        mux.Handle("DELETE /api/projects/{id}", authMiddleware(http.HandlerFunc(projectHandler.DeleteProject)))

        mux.HandleFunc("GET /api/experiences", experienceHandler.GetExperiences)
        mux.Handle("POST /api/experiences", authMiddleware(http.HandlerFunc(experienceHandler.CreateExperience)))
        mux.Handle("PUT /api/experiences/{id}", authMiddleware(http.HandlerFunc(experienceHandler.UpdateExperience)))
        mux.Handle("DELETE /api/experiences/{id}", authMiddleware(http.HandlerFunc(experienceHandler.DeleteExperience)))
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

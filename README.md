# Dimitri Schulz Amado — Full-Stack Portfolio

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![DigitalOcean](https://img.shields.io/badge/DigitalOcean-%23008bcf.svg?style=for-the-badge&logo=digitalOcean&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

A modernized, full-stack personal portfolio. What started as a static HTML/JS site has evolved into a robust decoupled application featuring a Vanilla JavaScript frontend and a resilient Golang backend powered by SQLite.

## 🏗️ Architecture & Stack

- **Frontend:** Pure HTML5, CSS3, and Vanilla JavaScript. Deployed statically to GitHub Pages. Focuses on performance and SEO.
- **Backend:** Go 1.26 API following Clean Architecture/DDD patterns. 
- **Database:** SQLite with automated initial seeding (embedded JSON data).
- **Security & Reliability:** Built-in CORS configuration, IP-based Rate Limiting, and graceful failure handling.
- **CI/CD & DevOps:** Automated via GitHub Actions. Tests are enforced (50%+ coverage required). Backend is packaged into a Docker container and deployed via SSH to a DigitalOcean Droplet.

## ✨ Key Features

- **Dynamic Content:** Projects are served from a SQLite database rather than hardcoded static files.
- **Automated Seeding:** If the database is empty, the backend automatically seeds initial data at runtime using Go's `//go:embed` directive.
- **Contact Form Integration:** Includes an SMTP sender to natively dispatch emails when the contact form is submitted.
- **Resilient Infrastructure:** Backend runs containerized with persistent volumes to ensure data survives across restarts.

## 📂 Project Structure

```text
portfolio/
├── frontend/
│   ├── index.html          # Entry point
│   ├── css/                # Styling (CSS Modules)
│   ├── js/                 # Client logic and API fetching
│   └── package.json        # Test dependencies (Vitest)
│
├── backend/
│   ├── cmd/server/         # Entrypoint & Seeder logic
│   ├── internal/
│   │   ├── domain/         # Entities (Project, ContactMessage)
│   │   ├── handler/        # HTTP Controllers
│   │   ├── middleware/     # Security (CORS, Rate Limiter)
│   │   ├── repository/     # SQLite persistence layer
│   │   └── sender/         # External SMTP service
│   ├── Dockerfile          # Multi-stage build definition
│   └── go.mod
│
└── .github/workflows/      # Automated Deployment Pipelines
```

## 🚀 Getting Started (Local Development)

### Prerequisites
- [Go 1.26+](https://golang.org/dl/)
- Node.js & npm (for frontend tests)

### Running the Backend

```bash
cd backend

# The backend will start on port 8080 and create portfolio.db automatically
go run cmd/server/main.go
```

### Running the Frontend

The frontend points to `http://localhost:8080` by default during local development. Simply open `frontend/index.html` in your browser, or use a local static server:

```bash
cd frontend
python3 -m http.server 3000
# Open http://localhost:3000
```

## ⚙️ Environment Variables

To fully run the backend in production (or locally with email capabilities), set the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | The port the HTTP server listens on. | `8080` |
| `DB_PATH` | Path to the SQLite database file. | `portfolio.db` |
| `ALLOWED_ORIGIN` | CORS allowed origin (Frontend URL). | `*` |
| `SMTP_HOST` | Hostname of your email provider (e.g., smtp.gmail.com). | *empty* |
| `SMTP_PORT` | Port for your SMTP server (e.g., 587). | *empty* |
| `SMTP_USER` | SMTP authentication username. | *empty* |
| `SMTP_PASSWORD` | SMTP authentication password/app token. | *empty* |
| `CONTACT_TO_EMAIL`| Destination email to receive contact form submissions. | *empty* |

*(Note: If SMTP variables are missing, the backend will still run but will fallback to logging contact messages to the console).*

## 🔄 CI/CD Pipelines

This project utilizes two independent GitHub Actions workflows:

1. **Frontend Pipeline (`frontend-ci.yml`)**: Triggers on UI changes. Runs Vitest unit tests, injects the production `BACKEND_URL`, and deploys statically to GitHub Pages.
2. **Backend Pipeline (`backend-ci.yml`)**: Triggers on backend changes. Runs Go tests, enforces coverage, builds a Docker image to GitHub Container Registry (GHCR), and triggers a rolling update via SSH to the DigitalOcean Droplet.

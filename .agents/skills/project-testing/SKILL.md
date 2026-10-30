---
name: project-testing
description: >-
  Always use this skill whenever creating, updating, debugging, or running unit
  and integration tests, or verifying code quality for the portfolio project.
---

# Portfolio Project Testing Skill

Follow this workflow whenever adding features, modifying logic, or running tests in this repository.

## 1. Test Architecture & Conventions

All test files MUST be co-located inside their respective Django package's `tests/` directory (`<package>/tests/*_test.py`) matching the module they test:

- **`config/tests/`**:
  - [`config/tests/settings_test.py`](../../../config/tests/settings_test.py): Tests `config/settings.py` (such as `get_database_config()` across SQLite, Vercel `/tmp/schema.db`, custom `DB_PATH`, and PostgreSQL `DATABASE_URL`).
- **`portfolio/tests/`**:
  - [`portfolio/tests/models_test.py`](../../../portfolio/tests/models_test.py): Tests ORM models in `portfolio/models.py` (`Project`, `About`, `Experience`) using `@pytest.mark.django_db`.
  - [`portfolio/tests/serializers_test.py`](../../../portfolio/tests/serializers_test.py): Tests DRF serializers in `portfolio/serializers.py` (`ProjectSerializer` field constraints and URL validation).
  - [`portfolio/tests/views_test.py`](../../../portfolio/tests/views_test.py): Tests API views and routes in `portfolio/views.py` and `portfolio/urls.py` using `rest_framework.test.APIClient`.
- **Future Apps (`<app_name>/tests/`)**:
  - Whenever a new Django app or module is created, place its unit and integration tests inside `<app_name>/tests/<module>_test.py`.

## 2. Running the Test & Quality Suite

Run the full verification script or individual commands using the project's virtual environment (`.venv/bin/pytest`):

1. **Run all checks (recommended)**:
   Execute [run_checks.sh](./scripts/run_checks.sh):
   ```bash
   bash .agents/skills/project-testing/scripts/run_checks.sh
   ```

2. **Run only pytest**:
   ```bash
   ./.venv/bin/pytest -s -v
   ```

3. **Run linters (CI parity with `.github/workflows/ci_cd.yml`)**:
   ```bash
   ./.venv/bin/flake8 .
   ./.venv/bin/pylint config portfolio manage.py
   ```

## 3. Validation Checklist Before Finishing

1. Every new feature or modified logic has corresponding unit/route tests in `<package>/tests/*_test.py`.
2. `./.venv/bin/pytest -s -v` passes with zero failures.
3. `./.venv/bin/flake8 .` and `./.venv/bin/pylint config portfolio manage.py` pass with zero errors.

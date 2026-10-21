import pytest
from fastapi.testclient import TestClient

from app.models.settings.database_connection_handler import (
    get_database_url,
    init_db,
    normalize_postgres_url,
)
from app.models.settings.metadata import metadata
from main import app


def test_main_root_serves_frontend():
    client = TestClient(app)
    response = client.get("/")
    assert response.status_code == 200
    assert "Dimitri" in response.text


def test_main_health_endpoint():
    client = TestClient(app)
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


def test_get_database_url_vercel(monkeypatch):
    monkeypatch.setenv("VERCEL", "1")
    monkeypatch.delenv("DATABASE_URL", raising=False)
    monkeypatch.delenv("DB_PATH", raising=False)
    url = get_database_url()
    assert "/tmp/schema.db" in url


def test_get_database_url_sqlite_custom_path(monkeypatch):
    monkeypatch.setenv("DB_PATH", "custom/path.db")
    monkeypatch.delenv("DATABASE_URL", raising=False)
    url = get_database_url()
    assert "custom/path.db" in url


@pytest.mark.asyncio
async def test_init_db():
    await init_db()
    assert "projects" in metadata.tables


def test_normalize_postgres_url():
    raw = "postgresql://user:pass@ep-xyz.aws.neon.tech/neondb?sslmode=require"
    clean = normalize_postgres_url(raw)
    expected_prefix = (
        "postgresql+asyncpg://user:pass@ep-xyz.aws.neon.tech/neondb"
    )
    assert clean.startswith(expected_prefix)
    assert "ssl=require" in clean


def test_get_database_url_with_postgres_env(monkeypatch):
    raw = "postgres://user:pass@ep-xyz.aws.neon.tech/neondb?sslmode=require"
    monkeypatch.setenv("DATABASE_URL", raw)
    url = get_database_url()
    assert url.startswith("postgresql+asyncpg://")
    assert "ssl=require" in url

import sqlite3

from fastapi.testclient import TestClient

from app.models.settings.database_connection_handler import get_database_url
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


def test_get_database_url_creates_schema_when_missing(tmp_path, monkeypatch):
    test_db = str(tmp_path / "test_init.db")
    monkeypatch.setenv("DB_PATH", test_db)
    monkeypatch.delenv("DATABASE_URL", raising=False)
    url = get_database_url()
    assert test_db in url

    with sqlite3.connect(test_db) as conn:
        cursor = conn.cursor()
        query = (
            "SELECT name FROM sqlite_master "
            "WHERE type='table' AND name='projects';"
        )
        cursor.execute(query)
        table = cursor.fetchone()
        assert table is not None
        assert table[0] == "projects"

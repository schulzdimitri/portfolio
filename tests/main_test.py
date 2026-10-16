from fastapi.testclient import TestClient
from main import app
from app.models.settings.database_connection_handler import get_database_url


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

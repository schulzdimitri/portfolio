from config.settings import get_database_config


def test_get_database_config_sqlite_default(monkeypatch):
    monkeypatch.setenv("TESTING", "1")
    monkeypatch.delenv("DATABASE_URL", raising=False)
    monkeypatch.delenv("VERCEL", raising=False)
    monkeypatch.delenv("DB_PATH", raising=False)
    cfg = get_database_config()
    assert cfg["ENGINE"] == "django.db.backends.sqlite3"
    assert str(cfg["NAME"]).endswith("schema.db")


def test_get_database_config_vercel(monkeypatch):
    monkeypatch.setenv("TESTING", "1")
    monkeypatch.setenv("VERCEL", "1")
    monkeypatch.delenv("DATABASE_URL", raising=False)
    monkeypatch.delenv("DB_PATH", raising=False)
    cfg = get_database_config()
    assert cfg["ENGINE"] == "django.db.backends.sqlite3"
    assert "/tmp/schema.db" in str(cfg["NAME"])


def test_get_database_config_custom_db_path(monkeypatch):
    monkeypatch.setenv("TESTING", "1")
    monkeypatch.setenv("DB_PATH", "custom/path.db")
    monkeypatch.delenv("DATABASE_URL", raising=False)
    cfg = get_database_config()
    assert cfg["NAME"] == "custom/path.db"


def test_get_database_config_postgres_env(monkeypatch):
    raw = "postgres://user:pass@ep-xyz.aws.neon.tech/neondb?sslmode=require"
    monkeypatch.setenv("TESTING", "0")
    monkeypatch.setenv("DATABASE_URL", raw)
    cfg = get_database_config()
    assert cfg["ENGINE"] == "django.db.backends.postgresql"
    assert cfg["NAME"] == "neondb"
    assert cfg["HOST"] == "ep-xyz.aws.neon.tech"

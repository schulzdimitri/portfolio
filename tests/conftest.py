import os
import pytest
from sqlalchemy import create_engine
from app.models.settings.metadata import metadata
import app.models.entities.projects  # noqa: F401

os.environ["TESTING"] = "1"
os.environ.pop("DATABASE_URL", None)


@pytest.fixture(autouse=True, scope="session")
def setup_database():
    sync_engine = create_engine("sqlite:///schema.db")
    metadata.create_all(sync_engine)
    sync_engine.dispose()

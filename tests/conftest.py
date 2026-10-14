import os
import sqlite3
import pytest


@pytest.fixture(autouse=True, scope="session")
def setup_database():
    schema_path = os.path.join(
        os.path.dirname(__file__), "..", "init", "schema.sql"
    )
    if os.path.exists(schema_path):
        with open(schema_path, "r", encoding="utf-8") as schema_file:
            schema_sql = schema_file.read()
        conn = sqlite3.connect("schema.db")
        conn.executescript(schema_sql)
        conn.close()

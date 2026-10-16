import os
import shutil
import sqlite3
from typing import Optional
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from sqlalchemy.orm import sessionmaker
from app.models.settings.metadata import metadata


DEFAULT_SCHEMA_SQL = """
CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
"""


def get_database_url() -> str:
    database_url = os.getenv("DATABASE_URL")
    if database_url:
        return database_url

    is_vercel = os.getenv("VERCEL") == "1" or "VERCEL" in os.environ
    db_dir = "/tmp" if is_vercel else "."
    db_path = os.getenv("DB_PATH", os.path.join(db_dir, "schema.db"))

    if not os.path.exists(db_path):
        base_dir = os.path.abspath(
            os.path.join(os.path.dirname(__file__), "..", "..", "..")
        )
        bundled_db = os.path.join(base_dir, "schema.db")
        schema_file_path = os.path.join(base_dir, "init", "schema.sql")

        if is_vercel and os.path.exists(bundled_db):
            shutil.copyfile(bundled_db, db_path)
        else:
            sql_script = DEFAULT_SCHEMA_SQL
            if os.path.exists(schema_file_path):
                with open(schema_file_path, "r", encoding="utf-8") as f:
                    sql_script = f.read()
            elif os.path.exists("init/schema.sql"):
                with open("init/schema.sql", "r", encoding="utf-8") as f:
                    sql_script = f.read()

            with sqlite3.connect(db_path) as conn:
                conn.executescript(sql_script)

    return f"sqlite+aiosqlite:///{db_path}"


CONNECTION_STRING = get_database_url()

engine = create_async_engine(
    url=CONNECTION_STRING,
    echo=False,
    pool_size=2,
    max_overflow=0,
    pool_timeout=30,
)

async_session = sessionmaker(
    bind=engine, class_=AsyncSession, expire_on_commit=False
)


class DBConnectionHandler:
    def __init__(self) -> None:
        self.session: Optional[AsyncSession] = None

    async def __aenter__(self) -> "DBConnectionHandler":
        self.session = async_session()
        return self

    async def __aexit__(self, exc_type, exc, traceback) -> None:
        if self.session:
            await self.session.close()

    async def create_tables(self, exc_type, exc_val, exc_tb) -> None:
        async with engine.begin() as conn:
            await conn.run_sync(metadata.create_all)

import os
from typing import Optional
from urllib.parse import urlparse, urlencode, urlunparse
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from sqlalchemy.orm import sessionmaker
from app.models.settings.metadata import metadata
import app.models.entities.projects  # noqa: F401

if os.getenv("TESTING") != "1":
    try:
        from dotenv import load_dotenv

        load_dotenv()
    except ImportError:
        pass


def normalize_postgres_url(raw_url: str) -> str:
    parsed = urlparse(raw_url)
    clean_url = urlunparse(
        (
            "postgresql+asyncpg",
            parsed.netloc,
            parsed.path,
            "",
            urlencode({"ssl": "require"}),
            "",
        )
    )
    return clean_url


def get_database_url() -> str:
    database_url = os.getenv("DATABASE_URL")
    if database_url:
        if database_url.startswith("postgres://") or database_url.startswith(
            "postgresql://"
        ):
            return normalize_postgres_url(database_url)
        return database_url

    is_vercel = os.getenv("VERCEL") == "1" or "VERCEL" in os.environ
    db_dir = "/tmp" if is_vercel else "."
    db_path = os.getenv("DB_PATH", os.path.join(db_dir, "schema.db"))

    return f"sqlite+aiosqlite:///{db_path}"


CONNECTION_STRING = get_database_url()
is_postgres = CONNECTION_STRING.startswith("postgresql")

engine_kwargs = {
    "echo": False,
    "pool_size": 2,
    "max_overflow": 0,
    "pool_timeout": 30,
}

if is_postgres:
    engine_kwargs["connect_args"] = {"statement_cache_size": 0}
    engine_kwargs["pool_pre_ping"] = True

engine = create_async_engine(url=CONNECTION_STRING, **engine_kwargs)

async_session = sessionmaker(
    bind=engine, class_=AsyncSession, expire_on_commit=False
)


async def init_db() -> None:
    async with engine.begin() as conn:
        await conn.run_sync(metadata.create_all)


class DBConnectionHandler:
    def __init__(self) -> None:
        self.session: Optional[AsyncSession] = None

    async def __aenter__(self) -> "DBConnectionHandler":
        self.session = async_session()
        return self

    async def __aexit__(self, exc_type, exc, traceback) -> None:
        if self.session:
            await self.session.close()

    async def create_tables(self) -> None:
        await init_db()

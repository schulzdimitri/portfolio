from typing import Optional
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from sqlalchemy.orm import sessionmaker
from app.models.settings.metadata import metadata

CONNECTION_STRING = "sqlite+aiosqlite:///schema.db"

engine = create_async_engine(
    url=CONNECTION_STRING, echo=False, pool_size=2, max_overflow=0, pool_timeout=30
)

async_session = sessionmaker(bind=engine, class_=AsyncSession, expire_on_commit=False)


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

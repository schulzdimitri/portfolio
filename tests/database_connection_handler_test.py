import pytest
from app.models.settings.database_connection_handler import DBConnectionHandler


class TestDBConnectionHandler:
    @pytest.mark.asyncio
    @pytest.mark.skip(reason="Connecting with DB")
    async def test_connection(self) -> None:
        async with DBConnectionHandler() as db:
            assert db.session is not None

from sqlalchemy import insert, select, update, delete
from app.models.entities.projects import projects
from app.models.settings.database_connection_handler import DBConnectionHandler
from app.models.repositories.interfaces.projects import (
    ProjectsRepositoryInterface,
)


class ProjectsRepository(ProjectsRepositoryInterface):
    async def insert_project(self, project_info: dict) -> None:
        async with DBConnectionHandler() as db:
            query = insert(projects).values(**project_info)
            await db.session.execute(query)
            await db.session.commit()

    async def get_project_by_name(self, project_name: str) -> list[dict]:
        async with DBConnectionHandler() as db:
            query = select(projects).where(
                projects.c.name == project_name
            )
            result = await db.session.execute(query)
            rows = result.mappings().fetchall()

            return [dict(row) for row in rows]

    async def get_all_projects(self) -> list[dict]:
        async with DBConnectionHandler() as db:
            query = select(projects)
            result = await db.session.execute(query)
            rows = result.mappings().fetchall()

            return [dict(row) for row in rows]

    async def update_project(self, project_info: dict) -> None:
        async with DBConnectionHandler() as db:
            query = (
                update(projects)
                .where(projects.c.name == project_info["name"])
                .values(**project_info)
            )
            await db.session.execute(query)
            await db.session.commit()

    async def delete_project(self, project_name: dict) -> None:
        async with DBConnectionHandler() as db:
            query = delete(projects).where(
                projects.c.name == project_name
            )
            await db.session.execute(query)
            await db.session.commit()

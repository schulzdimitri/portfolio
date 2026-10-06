# pylint: disable = W012 
from sqlalchemy import insert, select, update, delete
from app.models.entities.projects import projects
from app.models.settings.database_connection_handler import DBConnectionHandler


class ProjectsRepository:
    async def insert_project(self, project_infos: dict) -> None:
        async with DBConnectionHandler() as db:
            query = insert(projects).values(**project_infos)
            await db.session.execute(query)
            await db.session.commit()
        

    async def get_project_by_name(self, project_name: str) -> list[dict]:
        async with DBConnectionHandler() as db:
            query = (
                select(projects)
                .where(projects.c.project_name == project_name)
            )
            result = await db.session.execute(query)
            rows = result.fetchall()

            projects_list = [dict(row._mapping) for row in rows]

            return projects_list

            

    async def get_all_projects(self) -> list[dict]:
        async with DBConnectionHandler() as db:
            query = select(projects)
            result = await db.session.execute(query)
            rows = result.fetchall()

            projects_list = [dict(row._mapping) for row in rows]

            return projects_list

    async def update_project(self, project_infos: dict) -> None:
        async with DBConnectionHandler() as db:
            query = (
                update(projects)
                .where(projects.c.project_name == project_infos["project_name"])
                .values(**project_infos)
            )
            await db.session.execute(query)
            await db.session.commit()

    async def delete_project(self, project_infos: dict) -> None:
        async with DBConnectionHandler() as db:
            query = (
                delete(projects)
                .where(projects.c.project_name == project_infos["project_name"])
            )
            await db.session.execute(query)
            await db.session.commit()

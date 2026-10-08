from app.models.repositories.interfaces.projects import (
    ProjectsRepositoryInterface,
)
from .interfaces.project_finder import ProjectFinderInterface


class ProjectFinder(ProjectFinderInterface):
    def __init__(self, project_repository: ProjectsRepositoryInterface) -> None:
        self.__project_repository = project_repository

    async def find_project_by_name(self, project_name: str) -> dict:
        project = await self.__project_repository.get_project_by_name(project_name)
        return {"type": "PROJECT", "count": len(project), "attributes": project}

    async def find_all_projects(self) -> dict:
        projects = await self.__project_repository.get_all_projects()
        return {"type": "PROJECT", "count": len(projects), "attributes": projects}

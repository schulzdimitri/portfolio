from app.models.repositories.interfaces.projects import (
    ProjectsRepositoryInterface,
)
from .interfaces.project_register import ProjectRegisterInterface


class ProjectRegister(ProjectRegisterInterface):
    def __init__(
        self, project_repository: ProjectsRepositoryInterface
    ) -> None:
        self.__project_repository = project_repository

    async def register_project(self, project_info: dict) -> dict:
        self.__validate_project_data(project_info)
        await self.__registry_project(project_info)
        return self.__format_response(project_info)

    def __validate_project_data(self, project_info: dict) -> None:
        project_name = project_info["name"]
        project_description = project_info["description"]
        project_url = project_info["url"]

        if not project_name or not project_description or not project_url:
            raise ValueError("Project name, description, and URL are required")
        if (
            not isinstance(project_name, str)
            or not isinstance(project_description, str)
            or not isinstance(project_url, str)
        ):
            raise ValueError(
                "Project name, description, and URL must be strings"
            )

        if not project_url.startswith("http"):
            raise ValueError("Project URL must be a valid URL")

    async def __registry_project(self, project_info: dict) -> None:
        await self.__project_repository.insert_project(project_info)

    def __format_response(self, project_info: dict) -> dict:
        return {"type": "PROJECT", "count": 1, "attributes": project_info}

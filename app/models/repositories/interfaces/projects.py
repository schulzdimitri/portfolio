from abc import ABC, abstractmethod


class ProjectsRepositoryInterface(ABC):
    @abstractmethod
    async def insert_project(self, project_info: dict[str, str]) -> None:
        pass

    @abstractmethod
    async def get_project_by_name(self, project_name: str) -> list[dict]:
        pass

    @abstractmethod
    async def get_all_projects(self) -> list[dict]:
        pass

    @abstractmethod
    async def update_project(self, project_info: dict[str, str]) -> None:
        pass

    @abstractmethod
    async def delete_project(self, project_name: str) -> None:
        pass

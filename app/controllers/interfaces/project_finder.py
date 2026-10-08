from abc import ABC, abstractmethod


class ProjectFinderInterface(ABC):

    @abstractmethod
    async def find_project_by_name(self, project_name: str) -> dict:
        pass

    @abstractmethod
    async def find_all_projects(self) -> dict:
        pass

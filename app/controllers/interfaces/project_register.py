from abc import ABC, abstractmethod


class ProjectRegisterInterface(ABC):

    @abstractmethod
    async def register_project(self, project_info: dict) -> dict:
        pass

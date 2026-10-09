# pylint:disable=C1803
import pytest
from app.controllers.project_finder import ProjectFinder


class MockProjectRepository:
    def __init__(self) -> None:
        self.get_projects_att = {}
        self.get_project_by_name_att = {}

    async def get_all_projects(self) -> list[dict]:
        if "projects" in self.get_projects_att:
            return self.get_projects_att["projects"]
        return []

    async def get_project_by_name(self, project_name: str) -> list[dict]:
        self.get_project_by_name_att["project_name"] = project_name
        if "projects" in self.get_projects_att:
            return [
                project
                for project in self.get_projects_att["projects"]
                if project.get("project_name") == project_name
            ]
        return []


@pytest.mark.asyncio
async def test_find_all_projects() -> None:
    project_info_1 = {
        "project_name": "Project 1",
        "description": "Description 1",
        "url": "https://project1.com",
    }
    project_info_2 = {
        "project_name": "Project 2",
        "description": "Description 2",
        "url": "https://project2.com",
    }

    project_repository = MockProjectRepository()
    project_repository.get_projects_att["projects"] = [
        project_info_1,
        project_info_2,
    ]
    project_finder = ProjectFinder(project_repository)

    response = await project_finder.find_all_projects()

    assert response["type"] == "PROJECT"
    assert response["count"] == 2
    assert response["attributes"] == [project_info_1, project_info_2]


@pytest.mark.asyncio
async def test_find_all_projects_empty() -> None:
    project_repository = MockProjectRepository()
    project_finder = ProjectFinder(project_repository)

    response = await project_finder.find_all_projects()

    assert response["type"] == "PROJECT"
    assert response["count"] == 0
    assert response["attributes"] == []


@pytest.mark.asyncio
async def test_find_project_by_name() -> None:
    project_info = {
        "project_name": "Project 1",
        "description": "Description 1",
        "url": "https://project1.com",
    }
    project_repository = MockProjectRepository()
    project_repository.get_projects_att["projects"] = [project_info]
    project_finder = ProjectFinder(project_repository)

    response = await project_finder.find_project_by_name("Project 1")

    assert (
        project_repository.get_project_by_name_att["project_name"]
        == "Project 1"
    )
    assert response["type"] == "PROJECT"
    assert response["count"] == 1
    assert response["attributes"] == [project_info]


@pytest.mark.asyncio
async def test_find_project_by_name_not_found() -> None:
    project_repository = MockProjectRepository()
    project_finder = ProjectFinder(project_repository)

    response = await project_finder.find_project_by_name("Nonexistent")

    assert response["type"] == "PROJECT"
    assert response["count"] == 0
    assert response["attributes"] == []

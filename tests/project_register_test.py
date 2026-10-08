# pylint:disable=C1803
import pytest
from app.controllers.project_register import ProjectRegister


class MockProjectRepository:
    def __init__(self) -> None:
        self.insert_projects_att = {}

    async def insert_project(self, project_info: dict) -> None:
        self.insert_projects_att["project_info"] = project_info


@pytest.mark.asyncio
async def test_register_project() -> None:
    project_repository = MockProjectRepository()
    project_register = ProjectRegister(project_repository)

    project_info = {
        "project_name": "Project Name",
        "project_description": "Project Description",
        "project_url": "https://project-url.com",
    }

    response = await project_register.register_project(project_info)

    assert project_repository.insert_projects_att["project_info"] == project_info

    assert response["type"] == "PROJECT"
    assert response["count"] == 1
    assert response["attributes"] == project_info


@pytest.mark.asyncio
async def test_register_project_invalid_project_name() -> None:
    project_repository = MockProjectRepository()
    project_register = ProjectRegister(project_repository)

    project_info = {
        "project_name": "",
        "project_description": "Project Description",
        "project_url": "https://project-url.com",
    }

    with pytest.raises(ValueError) as excinfo:
        await project_register.register_project(project_info)

    assert str(excinfo.value) == "Project name, description, and URL are required"
    assert project_repository.insert_projects_att == {}


@pytest.mark.asyncio
async def test_register_project_invalid_project_description() -> None:
    project_repository = MockProjectRepository()
    project_register = ProjectRegister(project_repository)

    project_info = {
        "project_name": "Project Name",
        "project_description": "",
        "project_url": "https://project-url.com",
    }

    with pytest.raises(ValueError) as excinfo:
        await project_register.register_project(project_info)

    assert str(excinfo.value) == "Project name, description, and URL are required"
    assert project_repository.insert_projects_att == {}


@pytest.mark.asyncio
async def test_register_project_invalid_project_url() -> None:
    project_repository = MockProjectRepository()
    project_register = ProjectRegister(project_repository)

    project_info = {
        "project_name": "Project Name",
        "project_description": "Project Description",
        "project_url": "",
    }

    with pytest.raises(ValueError) as excinfo:
        await project_register.register_project(project_info)

    assert str(excinfo.value) == "Project name, description, and URL are required"
    assert project_repository.insert_projects_att == {}


@pytest.mark.asyncio
async def test_register_invalid_project_url_not_startwith_http() -> None:
    project_repository = MockProjectRepository()
    project_register = ProjectRegister(project_repository)

    project_info = {
        "project_name": "Project Name",
        "project_description": "Project Description",
        "project_url": "invalid-url",
    }

    with pytest.raises(ValueError) as excinfo:
        await project_register.register_project(project_info)

    assert str(excinfo.value) == "Project URL must be a valid URL"
    assert project_repository.insert_projects_att == {}

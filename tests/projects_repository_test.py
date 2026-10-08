import pytest
from app.models.repositories.projects import ProjectsRepository


@pytest.mark.asyncio
@pytest.mark.skip(reason="Insert in db")
async def test_insert_project():
    project = {
        "project_name": "Test Project",
        "description": "Test Description",
        "url": "https://test.com",
    }

    repository = ProjectsRepository()
    await repository.insert_project(project)


@pytest.mark.asyncio
@pytest.mark.skip(reason="Select in db")
async def test_get_project_by_name():
    repository = ProjectsRepository()
    result = await repository.get_project_by_name("Test Project")
    assert result[0]["project_name"] == "Test Project"


@pytest.mark.asyncio
@pytest.mark.skip(reason="Select in db")
async def test_get_all_projects():
    repository = ProjectsRepository()
    result = await repository.get_all_projects()
    assert len(result) > 0


@pytest.mark.asyncio
@pytest.mark.skip(reason="Update in db")
async def test_update_project():
    project = {
        "project_name": "Test Project",
        "description": "Test Description",
        "url": "https://test.com",
    }

    repository = ProjectsRepository()
    await repository.update_project(project)


@pytest.mark.asyncio
@pytest.mark.skip(reason="Delete in db")
async def test_delete_project():
    project = {
        "project_name": "Test Project",
        "description": "Test Description",
        "url": "https://test.com",
    }

    repository = ProjectsRepository()
    await repository.delete_project(project)

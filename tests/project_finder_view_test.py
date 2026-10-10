import pytest
from app.controllers.interfaces.project_finder import ProjectFinderInterface
from app.views.http_types.http_request import HttpRequest
from app.views.project_finder import ProjectFinderView


class MockProjectFinderController(ProjectFinderInterface):
    def __init__(self, should_raise: bool = False) -> None:
        self.should_raise = should_raise
        self.find_project_by_name_att = {}

    async def find_project_by_name(self, project_name: str) -> dict:
        if self.should_raise:
            raise Exception("Database error")
        self.find_project_by_name_att["name"] = project_name
        return {
            "type": "PROJECT",
            "count": 1,
            "attributes": [{"name": project_name}],
        }

    async def find_all_projects(self) -> dict:
        if self.should_raise:
            raise Exception("Database error")
        return {
            "type": "PROJECT",
            "count": 2,
            "attributes": [{"name": "P1"}, {"name": "P2"}],
        }


@pytest.mark.asyncio
async def test_handle_find_all_projects_success() -> None:
    controller = MockProjectFinderController()
    view = ProjectFinderView(controller)

    response = await view.handle_find_all_projects()

    assert response.status_code == 200
    assert response.body["type"] == "PROJECT"
    assert response.body["count"] == 2


@pytest.mark.asyncio
async def test_handle_find_all_projects_error() -> None:
    controller = MockProjectFinderController(should_raise=True)
    view = ProjectFinderView(controller)

    response = await view.handle_find_all_projects()

    assert response.status_code == 500
    assert response.body == {"error": "Database error"}


@pytest.mark.asyncio
async def test_handle_find_project_by_name_success() -> None:
    controller = MockProjectFinderController()
    view = ProjectFinderView(controller)
    request = HttpRequest(path_params={"name": "Portfolio"})

    response = await view.handle_find_project_by_name(request)

    assert response.status_code == 200
    assert controller.find_project_by_name_att["name"] == "Portfolio"
    assert response.body["count"] == 1


@pytest.mark.asyncio
async def test_handle_find_project_by_name_error() -> None:
    controller = MockProjectFinderController(should_raise=True)
    view = ProjectFinderView(controller)
    request = HttpRequest(path_params={"name": "Portfolio"})

    response = await view.handle_find_project_by_name(request)

    assert response.status_code == 500
    assert response.body == {"error": "Database error"}

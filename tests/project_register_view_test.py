import pytest
from app.controllers.interfaces.project_register import (
    ProjectRegisterInterface,
)
from app.views.http_types.http_request import HttpRequest
from app.views.project_register import ProjectRegisterView


class MockProjectRegisterController(ProjectRegisterInterface):
    def __init__(
        self,
        raise_value_error: bool = False,
        raise_general_error: bool = False,
    ) -> None:
        self.raise_value_error = raise_value_error
        self.raise_general_error = raise_general_error
        self.registered_data = {}

    async def register_project(self, project_info: dict) -> dict:
        if self.raise_value_error:
            raise ValueError("Validation error")
        if self.raise_general_error:
            raise Exception("Internal server error")
        self.registered_data = project_info
        return {
            "type": "PROJECT",
            "count": 1,
            "attributes": project_info,
        }


@pytest.mark.asyncio
async def test_handle_project_register_success() -> None:
    controller = MockProjectRegisterController()
    view = ProjectRegisterView(controller)
    payload = {
        "name": "Project",
        "description": "Desc",
        "url": "https://url.com",
    }
    request = HttpRequest(body=payload)

    response = await view.handle_project_register(request)

    assert response.status_code == 201
    assert response.body["type"] == "PROJECT"
    assert response.body["attributes"] == payload
    assert controller.registered_data == payload


@pytest.mark.asyncio
async def test_handle_project_register_value_error() -> None:
    controller = MockProjectRegisterController(raise_value_error=True)
    view = ProjectRegisterView(controller)
    request = HttpRequest(body={})

    response = await view.handle_project_register(request)

    assert response.status_code == 400
    assert response.body == {"error": "Validation error"}


@pytest.mark.asyncio
async def test_handle_project_register_internal_error() -> None:
    controller = MockProjectRegisterController(raise_general_error=True)
    view = ProjectRegisterView(controller)
    request = HttpRequest(body={})

    response = await view.handle_project_register(request)

    assert response.status_code == 500
    assert response.body == {"error": "Internal server error"}

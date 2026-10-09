from app.views.http_types.http_request import HttpRequest
from app.views.http_types.http_response import HttpResponse
from app.controllers.interfaces.project_register import (
    ProjectRegisterInterface,
)


class ProjectRegisterView:
    def __init__(self, controller: ProjectRegisterInterface):
        self.__controller = controller

    async def handle_project_register(
        self, http_request: HttpRequest
    ) -> HttpResponse:
        try:
            project_data = http_request.body
            response = await self.__controller.register_project(project_data)
            return HttpResponse(status_code=200, body=response)
        except Exception as exception:
            return HttpResponse(
                status_code=500, body={"error": str(exception)}
            )

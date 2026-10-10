from app.views.http_types.http_request import HttpRequest
from app.views.http_types.http_response import HttpResponse
from app.controllers.interfaces.project_finder import ProjectFinderInterface


class ProjectFinderView:
    def __init__(self, controller: ProjectFinderInterface):
        self.__controller = controller

    async def handle_find_project_by_name(
        self, http_request: HttpRequest
    ) -> HttpResponse:
        try:
            path_params = http_request.path_params or {}
            project_name = path_params.get("name")
            response = await self.__controller.find_project_by_name(
                project_name
            )
            return HttpResponse(status_code=200, body=response)
        except Exception as exception:
            return HttpResponse(
                status_code=500, body={"error": str(exception)}
            )

    async def handle_find_all_projects(self) -> HttpResponse:
        try:
            response = await self.__controller.find_all_projects()
            return HttpResponse(status_code=200, body=response)
        except Exception as exception:
            return HttpResponse(
                status_code=500, body={"error": str(exception)}
            )

from fastapi import APIRouter
from fastapi.responses import JSONResponse
from app.composer.project_finder import project_finder_composer
from app.composer.project_register import project_register_composer
from app.views.http_types.http_request import HttpRequest
from app.validators.project_register import ProjectInput

project_routes = APIRouter(tags=["Projects"])


@project_routes.post("/projects")
async def create_project(body: ProjectInput):
    http_request = HttpRequest(body=body.model_dump())
    project_register = project_register_composer()

    http_response = await project_register.handle_project_register(
        http_request
    )

    return JSONResponse(
        content=http_response.body,
        status_code=http_response.status_code
    )


@project_routes.get("/projects")
async def get_projects():
    project_finder = project_finder_composer()
    http_response = await project_finder.handle_find_all_projects()

    return JSONResponse(
        content=http_response.body,
        status_code=http_response.status_code
    )


@project_routes.get("/projects/{project_name}")
async def get_project_by_name(project_name: str):
    http_request = HttpRequest(path_params={"name": project_name})
    project_finder = project_finder_composer()

    http_response = await project_finder.handle_find_project_by_name(
        http_request
    )

    return JSONResponse(
        content=http_response.body,
        status_code=http_response.status_code
    )

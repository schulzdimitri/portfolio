from fastapi import APIRouter
from fastapi.responses import JSONResponse

project_routes = APIRouter(tags=["Projects"])

@project_routes.post("/projects")
async def create_project():
    return JSONResponse(
        content={"message": "Project created"},
        status_code=201
    )

@project_routes.get("/projects")
async def get_projects():
    return JSONResponse(
        content={"message": "This will return all projects"},
        status_code=200
    )


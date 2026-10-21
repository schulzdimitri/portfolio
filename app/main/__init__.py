import os
from contextlib import asynccontextmanager
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.main.routes.projects import project_routes
from app.models.settings.database_connection_handler import init_db


@asynccontextmanager
async def lifespan(_app: FastAPI):
    await init_db()
    yield


app = FastAPI(lifespan=lifespan)

allowed_origins_env = os.getenv("ALLOWED_ORIGINS")
allowed_origins = (
    [
        origin.strip()
        for origin in allowed_origins_env.split(",")
        if origin.strip()
    ]
    if allowed_origins_env
    else ["*"]
)


app.add_middleware(
    CORSMiddleware,
    allow_origins=allowed_origins,
    allow_credentials=allowed_origins != ["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/health")
def health_check():
    return {"status": "ok"}


app.include_router(project_routes)

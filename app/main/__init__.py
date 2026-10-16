import os
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.main.routes.projects import project_routes

app = FastAPI()

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
    allow_credentials=True if allowed_origins != ["*"] else False,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/health")
def health_check():
    return {"status": "ok"}


app.include_router(project_routes)

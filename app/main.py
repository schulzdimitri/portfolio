from fastapi import FastAPI
from app.routes.projects import project_routes


app = FastAPI()


@app.get("/health")
def health_check():
    return {"status": "ok"}


app.include_router(project_routes)

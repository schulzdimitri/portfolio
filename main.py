import os
from fastapi.staticfiles import StaticFiles
from app.main import app

frontend_dir = os.path.join(
    os.path.dirname(os.path.abspath(__file__)), "frontend"
)
if os.path.isdir(frontend_dir):
    app.mount(
        "/", StaticFiles(directory=frontend_dir, html=True), name="frontend"
    )

if __name__ == "__main__":
    import uvicorn

    uvicorn.run("main:app", host="0.0.0.0", port=3001, reload=True)

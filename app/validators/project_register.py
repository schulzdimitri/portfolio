from pydantic import BaseModel, Field


class ProjectInput(BaseModel):
    name: str = Field(..., min_length=1, max_length=100)
    description: str = Field(..., min_length=1, max_length=500)
    url: str = Field(..., min_length=1, max_length=200)

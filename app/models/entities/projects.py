from sqlalchemy import Table, Column, Integer, String
from app.models.settings.metadata import metadata

projects = Table(
    'projects',
    metadata,
    Column('id', Integer, primary_key=True),
    Column('project_name', String(150), nullable=False),
    Column('description', String(255), nullable=False),
    Column('url', String(255), nullable=False),
)
import pytest
from portfolio.models import Experience, Project


@pytest.mark.django_db
def test_models_string_representation():
    project = Project.objects.create(
        name="Portfolio API",
        description="Django REST backend",
        url="https://example.com",
    )
    experience = Experience.objects.create(
        role="Software Engineer",
        company="Acme Corp",
        period="2024 - Present",
        description="Backend services",
        technologies=["Python", "PostgreSQL"],
    )

    assert str(project) == "Portfolio API"
    assert str(experience) == "Software Engineer at Acme Corp"

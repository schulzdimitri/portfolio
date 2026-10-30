import pytest
from portfolio.models import About, Experience, Project


@pytest.mark.django_db
def test_models_string_representation():
    project = Project.objects.create(
        name="Portfolio API",
        description="Django REST backend",
        url="https://example.com",
    )
    about = About.objects.create(
        title="Backend Engineer",
        summary="Building resilient systems",
        pillars=["Clean Code"],
        skills=["Python", "Django"],
    )
    experience = Experience.objects.create(
        role="Software Engineer",
        company="Acme Corp",
        period="2024 - Present",
        description="Backend services",
        technologies=["Python", "PostgreSQL"],
    )

    assert str(project) == "Portfolio API"
    assert str(about) == "Backend Engineer"
    assert str(experience) == "Software Engineer at Acme Corp"

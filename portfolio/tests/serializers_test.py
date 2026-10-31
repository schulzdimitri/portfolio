from portfolio.serializers import (
    AboutSerializer,
    ExperienceSerializer,
    ProjectSerializer,
)


def test_project_serializer_valid_with_long_description():
    long_description = "A" * 450
    serializer = ProjectSerializer(
        data={
            "name": "Scalable Service",
            "description": long_description,
            "url": "https://github.com/schulzdimitri/service",
        }
    )
    assert serializer.is_valid(), serializer.errors


def test_project_serializer_invalid_empty_fields():
    serializer = ProjectSerializer(
        data={"name": "", "description": "", "url": ""}
    )
    assert not serializer.is_valid()
    assert "name" in serializer.errors
    assert "description" in serializer.errors
    assert "url" in serializer.errors


def test_project_serializer_invalid_url_scheme():
    serializer = ProjectSerializer(
        data={
            "name": "Invalid URL Project",
            "description": "Valid description",
            "url": "ftp://invalid-scheme.com",
        }
    )
    assert not serializer.is_valid()
    assert "url" in serializer.errors
    assert "Project URL must be a valid URL" in str(serializer.errors["url"])


def test_about_serializer_valid():
    serializer = AboutSerializer(
        data={
            "title": "About Me",
            "summary": "Backend Software Engineer",
            "pillars": [{"title": "Clean Code", "description": "Testable"}],
            "skills": ["Python", "Django"],
        }
    )
    assert serializer.is_valid(), serializer.errors


def test_about_serializer_invalid_empty_fields():
    serializer = AboutSerializer(data={"title": "", "summary": ""})
    assert not serializer.is_valid()
    assert "title" in serializer.errors
    assert "summary" in serializer.errors
    assert "pillars" in serializer.errors
    assert "skills" in serializer.errors


def test_experience_serializer_valid():
    serializer = ExperienceSerializer(
        data={
            "role": "Backend Engineer",
            "company": "Acme Corp",
            "period": "2024 - Present",
            "description": "Building Django REST APIs",
            "technologies": ["Python", "Django", "PostgreSQL"],
        }
    )
    assert serializer.is_valid(), serializer.errors


def test_experience_serializer_invalid_empty_fields():
    serializer = ExperienceSerializer(data={"role": "", "company": ""})
    assert not serializer.is_valid()
    assert "role" in serializer.errors
    assert "company" in serializer.errors
    assert "period" in serializer.errors
    assert "description" in serializer.errors
    assert "technologies" in serializer.errors

from portfolio.serializers import ProjectSerializer


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

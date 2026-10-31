from unittest.mock import patch
from django.contrib.auth import get_user_model
import pytest
from rest_framework.test import APIClient


@pytest.fixture(name="api_client")
def fixture_api_client():
    return APIClient()


@pytest.fixture(name="admin_client")
def fixture_admin_client():
    user_model = get_user_model()
    admin_user = user_model.objects.create_superuser(
        username="admin",
        email="admin@example.com",
        password="adminpassword123",
    )
    client = APIClient()
    client.force_authenticate(user=admin_user)
    return client


def test_frontend_index_serves_portfolio(api_client):
    response = api_client.get("/")
    assert response.status_code == 200
    content = b"".join(response.streaming_content).decode("utf-8")
    assert "Dimitri" in content
    assert "About Me" in content
    assert "Experience" in content
    assert "Projects" in content


def test_frontend_index_returns_404_when_missing(api_client, monkeypatch):
    monkeypatch.setattr("os.path.isfile", lambda _path: False)
    response = api_client.get("/")
    assert response.status_code == 404


@pytest.mark.django_db
def test_default_django_admin_is_accessible(api_client):
    response = api_client.get("/admin/login/")
    assert response.status_code == 200
    assert "Django administration" in response.content.decode("utf-8")


def test_health_check_endpoint(api_client):
    response = api_client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


@pytest.mark.django_db
def test_unauthenticated_post_requests_are_forbidden(api_client):
    res_about = api_client.post("/about", data={}, format="json")
    res_exp = api_client.post("/experiences", data={}, format="json")
    res_proj = api_client.post("/projects", data={}, format="json")

    assert res_about.status_code == 403
    assert res_exp.status_code == 403
    assert res_proj.status_code == 403


@pytest.mark.django_db
def test_create_and_get_about(api_client, admin_client):
    payload = {
        "title": "About Me",
        "summary": "Backend Engineer focused on clean architecture.",
        "pillars": [{"title": "Clean Code", "description": "Modular design"}],
        "skills": ["Python", "Django", "PostgreSQL"],
    }

    create_res = admin_client.post("/about", data=payload, format="json")
    assert create_res.status_code == 201
    create_body = create_res.json()
    assert create_body["type"] == "ABOUT"
    assert create_body["count"] == 1
    assert create_body["attributes"]["title"] == "About Me"

    list_res = api_client.get("/about")
    assert list_res.status_code == 200
    list_body = list_res.json()
    assert list_body["type"] == "ABOUT"
    assert list_body["count"] == 1
    assert list_body["attributes"][0]["skills"] == [
        "Python",
        "Django",
        "PostgreSQL",
    ]


@pytest.mark.django_db
def test_about_validation_and_error_handling(api_client, admin_client):
    bad_res = admin_client.post("/about", data={"title": ""}, format="json")
    assert bad_res.status_code == 400
    assert "detail" in bad_res.json()

    with patch(
        "portfolio.views.About.objects.all",
        side_effect=RuntimeError("DB failure"),
    ):
        err_get = api_client.get("/about")
        assert err_get.status_code == 500
        assert err_get.json() == {"detail": "Internal server error"}

    with patch(
        "portfolio.views.AboutSerializer.is_valid",
        side_effect=RuntimeError("Unexpected failure"),
    ):
        err_post = admin_client.post("/about", data={}, format="json")
        assert err_post.status_code == 500
        assert err_post.json() == {"detail": "Internal server error"}


@pytest.mark.django_db
def test_create_and_get_experiences(api_client, admin_client):
    payload = {
        "role": "Backend Software Engineer",
        "company": "Cloud Systems",
        "period": "2024 - Present",
        "description": "Designed and maintained Django REST services.",
        "technologies": ["Python", "Django", "PostgreSQL"],
    }

    create_res = admin_client.post("/experiences", data=payload, format="json")
    assert create_res.status_code == 201
    create_body = create_res.json()
    assert create_body["type"] == "EXPERIENCE"
    assert create_body["count"] == 1
    assert create_body["attributes"]["role"] == "Backend Software Engineer"

    list_res = api_client.get("/experiences")
    assert list_res.status_code == 200
    list_body = list_res.json()
    assert list_body["type"] == "EXPERIENCE"
    assert list_body["count"] == 1
    assert list_body["attributes"][0]["company"] == "Cloud Systems"


@pytest.mark.django_db
def test_experiences_validation_and_error_handling(api_client, admin_client):
    bad_res = admin_client.post(
        "/experiences", data={"role": ""}, format="json"
    )
    assert bad_res.status_code == 400
    assert "detail" in bad_res.json()

    with patch(
        "portfolio.views.Experience.objects.all",
        side_effect=RuntimeError("DB failure"),
    ):
        err_get = api_client.get("/experiences")
        assert err_get.status_code == 500
        assert err_get.json() == {"detail": "Internal server error"}

    with patch(
        "portfolio.views.ExperienceSerializer.is_valid",
        side_effect=RuntimeError("Unexpected failure"),
    ):
        err_post = admin_client.post("/experiences", data={}, format="json")
        assert err_post.status_code == 500
        assert err_post.json() == {"detail": "Internal server error"}


@pytest.mark.django_db
def test_create_and_get_projects(api_client, admin_client):
    payload = {
        "name": "Django Portfolio",
        "description": "Migrated portfolio backend using Django and DRF",
        "url": "https://example.com/django-portfolio",
    }

    create_res = admin_client.post("/projects", data=payload, format="json")
    assert create_res.status_code == 201
    create_body = create_res.json()
    assert create_body["type"] == "PROJECT"
    assert create_body["count"] == 1
    assert create_body["attributes"]["name"] == "Django Portfolio"

    list_res = api_client.get("/projects")
    assert list_res.status_code == 200
    list_body = list_res.json()
    assert list_body["type"] == "PROJECT"
    assert list_body["count"] == 1
    assert list_body["attributes"][0]["name"] == "Django Portfolio"

    detail_res = api_client.get("/projects/Django%20Portfolio")
    assert detail_res.status_code == 200
    detail_body = detail_res.json()
    assert detail_body["type"] == "PROJECT"
    assert detail_body["count"] == 1
    assert detail_body["attributes"][0]["name"] == "Django Portfolio"


@pytest.mark.django_db
def test_get_project_by_name_not_found(api_client):
    response = api_client.get("/projects/non_existent_project")
    assert response.status_code == 200
    body = response.json()
    assert body["type"] == "PROJECT"
    assert body["count"] == 0
    assert body["attributes"] == []


@pytest.mark.django_db
def test_create_project_validation_errors(admin_client):
    response = admin_client.post("/projects", data={"name": ""}, format="json")
    assert response.status_code == 400
    assert "detail" in response.json()


@pytest.mark.django_db
def test_create_project_invalid_url(admin_client):
    payload = {
        "name": "Invalid URL Project",
        "description": "Valid description",
        "url": "invalid_url_without_http",
    }
    response = admin_client.post("/projects", data=payload, format="json")
    assert response.status_code == 400
    assert "detail" in response.json()


@pytest.mark.django_db
def test_project_views_error_handling(api_client, admin_client):
    with patch(
        "portfolio.views.Project.objects.all",
        side_effect=RuntimeError("DB failure"),
    ):
        res_get = api_client.get("/projects")
        assert res_get.status_code == 500
        assert res_get.json() == {"detail": "Internal server error"}

    with patch(
        "portfolio.views.ProjectSerializer.is_valid",
        side_effect=RuntimeError("Unexpected error"),
    ):
        res_post = admin_client.post(
            "/projects",
            data={
                "name": "P1",
                "description": "D1",
                "url": "https://example.com",
            },
            format="json",
        )
        assert res_post.status_code == 500
        assert res_post.json() == {"detail": "Internal server error"}

    with patch(
        "portfolio.views.Project.objects.filter",
        side_effect=RuntimeError("DB failure"),
    ):
        res_detail = api_client.get("/projects/P1")
        assert res_detail.status_code == 500
        assert res_detail.json() == {"detail": "Internal server error"}


def test_cors_preflight(api_client):
    response = api_client.options(
        "/projects",
        HTTP_ORIGIN="http://localhost:3000",
        HTTP_ACCESS_CONTROL_REQUEST_METHOD="GET",
    )
    assert response.status_code == 200
    assert response["Access-Control-Allow-Origin"] == "*"

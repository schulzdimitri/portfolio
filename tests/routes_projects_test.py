import uuid
import pytest
from fastapi.testclient import TestClient
from app.main import app


@pytest.fixture
def client():
    return TestClient(app)


def test_health_check(client):
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "ok"}


def test_create_and_get_projects(client):
    unique_name = f"Project {uuid.uuid4().hex[:8]}"
    payload = {
        "name": unique_name,
        "description": "Route description",
        "url": "https://example.com/project",
    }

    create_res = client.post("/projects", json=payload)
    assert create_res.status_code == 201
    create_body = create_res.json()
    assert create_body["type"] == "PROJECT"
    assert create_body["count"] == 1
    assert create_body["attributes"]["name"] == unique_name

    get_all_res = client.get("/projects")
    assert get_all_res.status_code == 200
    get_all_body = get_all_res.json()
    assert get_all_body["type"] == "PROJECT"
    assert any(
        p["name"] == unique_name for p in get_all_body["attributes"]
    )

    get_single_res = client.get(f"/projects/{unique_name}")
    assert get_single_res.status_code == 200
    get_single_body = get_single_res.json()
    assert get_single_body["type"] == "PROJECT"
    assert get_single_body["count"] == 1
    assert get_single_body["attributes"][0]["name"] == unique_name


def test_get_project_not_found(client):
    response = client.get("/projects/non_existent_project_xyz_123")
    assert response.status_code == 200
    body = response.json()
    assert body["type"] == "PROJECT"
    assert body["count"] == 0
    assert body["attributes"] == []


def test_create_project_validation_error(client):
    response = client.post("/projects", json={"name": ""})
    assert response.status_code == 422


def test_create_project_invalid_url(client):
    payload = {
        "name": "Invalid URL Project",
        "description": "Valid description",
        "url": "invalid_url_without_http",
    }
    response = client.post("/projects", json=payload)
    assert response.status_code == 400
    assert "error" in response.json()

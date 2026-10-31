from django.urls import path
from portfolio.views import (
    ExperienceListCreateView,
    FrontendIndexView,
    HealthCheckView,
    ProjectDetailView,
    ProjectListCreateView,
)

urlpatterns = [
    path("", FrontendIndexView.as_view(), name="frontend-index"),
    path("health", HealthCheckView.as_view(), name="health-check"),
    path(
        "experiences",
        ExperienceListCreateView.as_view(),
        name="experience-list",
    ),
    path("projects", ProjectListCreateView.as_view(), name="project-list"),
    path(
        "projects/<str:project_name>",
        ProjectDetailView.as_view(),
        name="project-detail",
    ),
]

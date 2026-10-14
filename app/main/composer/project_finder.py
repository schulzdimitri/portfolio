from app.models.repositories.projects import ProjectsRepository
from app.controllers.project_finder import ProjectFinder
from app.views.project_finder import ProjectFinderView


def project_finder_composer():
    model = ProjectsRepository()
    controller = ProjectFinder(model)
    view = ProjectFinderView(controller)

    return view

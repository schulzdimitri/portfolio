from app.models.repositories.projects import ProjectsRepository
from app.controllers.project_register import ProjectRegister
from app.views.project_register import ProjectRegisterView


def project_register_composer():
    model = ProjectsRepository()
    controller = ProjectRegister(model)
    view = ProjectRegisterView(controller)

    return view

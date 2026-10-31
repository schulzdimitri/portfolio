import { renderExperience } from './components/experience.js';
import { renderProjects } from './components/projects.js';
import { initNavbar } from './components/navbar.js';

function initApp() {
    initNavbar();

    const experienceContainer = document.getElementById('experience-container');
    renderExperience(experienceContainer);

    const projectsContainer = document.getElementById('projects-container');
    renderProjects(projectsContainer);
}

document.addEventListener('DOMContentLoaded', initApp);

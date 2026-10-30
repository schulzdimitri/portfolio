import { portfolioData } from './data/portfolioData.js';
import { renderAbout } from './components/about.js';
import { renderExperience } from './components/experience.js';
import { renderProjects } from './components/projects.js';
import { initNavbar } from './components/navbar.js';

function initApp() {
    initNavbar();

    const aboutContainer = document.getElementById('about-container');
    renderAbout(aboutContainer, portfolioData.about);

    const experienceContainer = document.getElementById('experience-container');
    renderExperience(experienceContainer, portfolioData.experiences);

    const projectsContainer = document.getElementById('projects-container');
    renderProjects(projectsContainer);
}

document.addEventListener('DOMContentLoaded', initApp);

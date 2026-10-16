import { projectApi } from './api.js';

const container = document.getElementById('projects-container');

function createProjectCard(project) {
    const card = document.createElement('article');
    card.className = 'project-card';

    const title = document.createElement('h3');
    title.textContent = project.name;

    const desc = document.createElement('p');
    desc.textContent = project.description;

    const link = document.createElement('a');
    link.href = project.url;
    link.textContent = 'View Project';
    link.target = '_blank';
    link.rel = 'noopener noreferrer';

    card.append(title, desc, link);
    return card;
}

async function renderProjects() {
    try {
        const response = await projectApi.getProjects();
        const projects = response.attributes || response.data || [];


        container.replaceChildren(...projects.map(createProjectCard));
    } catch (error) {
        console.error('Failed to render projects:', error);
    }
}

document.addEventListener('DOMContentLoaded', renderProjects);
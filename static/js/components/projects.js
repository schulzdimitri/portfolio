import { projectApi } from '../api.js';

function sanitizeProjectUrl(rawUrl) {
    try {
        const parsed = new URL(rawUrl);
        if (parsed.protocol === 'http:' || parsed.protocol === 'https:') {
            return parsed.href;
        }
    } catch {
        return '#';
    }
    return '#';
}

function createProjectCard(project) {
    const card = document.createElement('article');
    card.className = 'project-card';

    const title = document.createElement('h3');
    title.textContent = project.name;

    const desc = document.createElement('p');
    desc.textContent = project.description;

    const link = document.createElement('a');
    link.href = sanitizeProjectUrl(project.url);
    link.textContent = 'View Project';
    link.target = '_blank';
    link.rel = 'noopener noreferrer';

    card.append(title, desc, link);
    return card;
}

function showLoadingSkeleton(container) {
    const skeletonGrid = document.createElement('div');
    skeletonGrid.className = 'skeleton-grid';

    for (let i = 0; i < 3; i += 1) {
        const skeleton = document.createElement('div');
        skeleton.className = 'project-card skeleton';
        skeletonGrid.appendChild(skeleton);
    }
    container.replaceChildren(skeletonGrid);
}

function showEmptyOrError(container, message) {
    const notice = document.createElement('div');
    notice.className = 'status-notice';
    notice.textContent = message;
    container.replaceChildren(notice);
}

export async function renderProjects(container) {
    if (!container) return;

    showLoadingSkeleton(container);

    try {
        const response = await projectApi.getProjects();
        const projects = response.attributes || response.data || [];

        if (projects.length === 0) {
            showEmptyOrError(container, 'No projects published yet.');
            return;
        }

        const cards = projects.map(createProjectCard);
        container.replaceChildren(...cards);
    } catch (error) {
        showEmptyOrError(container, 'Failed to load projects. Please try again later.');
    }
}

const FOLDER_ICON = `<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>`;

const GITHUB_ICON = `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path></svg>`;

/**
 * Renders an array of project objects into a given DOM container.
 * @param {Array} projects
 * @param {HTMLElement} container
 */
export function renderProjects(projects, container) {
    container.innerHTML = projects.map(project => `
        <div class="project-card">
            <div class="project-card-header">
                ${FOLDER_ICON}
                <a href="${project.github}" target="_blank" class="project-card-link" aria-label="GitHub">
                    ${GITHUB_ICON}
                </a>
            </div>
            <h3>${project.title}</h3>
            <p>${project.description}</p>
            <ul class="project-tags">
                ${project.tags.map(tag => `<li>${tag}</li>`).join('')}
            </ul>
        </div>
    `).join('');
}

/**
 * Fetches project data from a JSON file and renders it into the grid.
 * @param {string} gridId - ID of the container element
 * @param {string} dataUrl - URL of the JSON data file
 */
export async function loadProjects(
    gridId = 'projects-grid',
    dataUrl = 'data/portfolio.json'
) {
    const grid = document.getElementById(gridId);
    if (!grid) return;

    try {
        const res = await fetch(dataUrl);
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        const data = await res.json();
        renderProjects(data.projects, grid);
    } catch {
        grid.innerHTML = '<p style="color:var(--text-muted)">Projetos indisponíveis no momento.</p>';
    }
}

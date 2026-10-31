const REQUEST_TIMEOUT_MS = 8000;

async function fetchJson(path) {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), REQUEST_TIMEOUT_MS);

    try {
        const response = await fetch(path, {
            method: 'GET',
            credentials: 'same-origin',
            headers: { Accept: 'application/json' },
            signal: controller.signal,
        });

        if (!response.ok) {
            throw new Error(`Error: ${response.status}`);
        }

        return await response.json();
    } finally {
        clearTimeout(timeoutId);
    }
}

export const portfolioApi = Object.freeze({
    getAbout() {
        return fetchJson('/about');
    },

    getExperiences() {
        return fetchJson('/experiences');
    },

    getProjects() {
        return fetchJson('/projects');
    },

    getProjectByName(name) {
        return fetchJson(`/projects/${encodeURIComponent(name)}`);
    },
});

export const projectApi = portfolioApi;

const API_URL = window.API_URL || (window.location.hostname === 'localhost' && window.location.port !== '3001' ? 'http://localhost:3001' : '');

export const projectApi = {
    async getProjects() {
        const response = await fetch(`${API_URL}/projects`, {
            method: 'GET',
            headers: { 'Accept': 'application/json' },
        });
        if (!response.ok) {
            throw new Error(`Error: ${response.status}`);
        }
        return response.json();
    },

    async getProjectByName(name) {
        const response = await fetch(`${API_URL}/projects/${encodeURIComponent(name)}`, {
            method: 'GET',
            headers: { 'Accept': 'application/json' },
        });
        if (!response.ok) {
            throw new Error(`Error: ${response.status}`);
        }
        return response.json();
    },

    async createProject(projectData) {
        const response = await fetch(`${API_URL}/projects`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
            body: JSON.stringify(projectData),
        });

        const data = await response.json().catch(() => ({}));

        if (!response.ok) {
            const errorMessage = Array.isArray(data.detail)
                ? data.detail.map((err) => `${err.loc.join('.')}: ${err.msg}`).join(', ')
                : data.detail || data.message || `Error: ${response.status}`;
            throw new Error(errorMessage);
        }

        return data;
    },
};

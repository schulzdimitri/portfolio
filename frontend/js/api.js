const API_URL = window.API_URL || 'http://localhost:3001';

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
};
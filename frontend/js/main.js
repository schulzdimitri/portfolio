import { initTabs } from './modules/tabs.js';
import { loadProjects } from './modules/projects.js';
import { loadExperiences } from './modules/experiences.js';
import { initContactForm } from './modules/contact.js';

function resolveApiBase() {
    const rawMeta = document.querySelector('meta[name="api-base"]')?.content?.trim() || '';
    const isLocalPage = ['localhost', '127.0.0.1'].includes(window.location.hostname);
    const metaPointsToLocalhost = /^https?:\/\/(localhost|127\.0\.0\.1)(:\d+)?$/i.test(rawMeta);

    if (rawMeta && !(metaPointsToLocalhost && !isLocalPage)) {
        return rawMeta;
    }

    if (isLocalPage) {
        return 'http://localhost:8080';
    }

    return '';
}

const API_BASE = resolveApiBase();

document.addEventListener('DOMContentLoaded', () => {
    initTabs();
    loadProjects('projects-grid', API_BASE);
    loadExperiences(API_BASE);
    initContactForm(API_BASE);
});

import { initTabs } from './modules/tabs.js';
import { loadProjects } from './modules/projects.js';
import { initContactForm } from './modules/contact.js';

const API_BASE = document.querySelector('meta[name="api-base"]')?.content
    || 'http://localhost:8080';

document.addEventListener('DOMContentLoaded', () => {
    initTabs();
    loadProjects();
    initContactForm(API_BASE);
});

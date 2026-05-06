import { initTabs } from './modules/tabs.js';
import { loadProjects } from './modules/projects.js';
import { loadExperiences } from './modules/experiences.js';
import { initContactForm } from './modules/contact.js';

const API_BASE = document.querySelector('meta[name="api-base"]')?.content
    || 'http://localhost:8080';

document.addEventListener('DOMContentLoaded', () => {
    initTabs();
    loadProjects('projects-grid', API_BASE);
    loadExperiences(API_BASE);
    initContactForm(API_BASE);
});

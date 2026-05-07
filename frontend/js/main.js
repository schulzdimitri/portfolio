import { initTabs } from './modules/tabs.js';
import { loadProjects } from './modules/projects.js';
import { loadExperiences } from './modules/experiences.js';
import { initContactForm } from './modules/contact.js';

export function resolveApiBase(runtimeApiBase = '', isLocalPage = ['localhost', '127.0.0.1'].includes(window.location.hostname)) {
    if (runtimeApiBase) {
        return runtimeApiBase;
    }

    if (isLocalPage) {
        return 'http://localhost:8080';
    }

    return '';
}

export async function loadRuntimeConfig() {
    try {
        const res = await fetch('config.json', { cache: 'no-store' });
        if (!res.ok) return { apiBase: '' };
        const json = await res.json();
        const raw = (json && json.apiBase) ? String(json.apiBase).trim() : '';
        if (!raw || /^\*+$/.test(raw) || /^REDACTED$/i.test(raw)) {
            console.warn('runtime config: apiBase is empty or masked, ignoring');
            return { apiBase: '' };
        }
        const isValid = /^(https?:\/\/[^\s/$.?#].[^\s]*)$/.test(raw) || /^\//.test(raw);
        if (!isValid) {
            console.warn('runtime config: apiBase looks invalid, ignoring', raw);
            return { apiBase: '' };
        }
        return { apiBase: raw };
    } catch (e) {
        return { apiBase: '' };
    }
}

export async function initApp() {
    const cfg = await loadRuntimeConfig();
    const API_BASE = resolveApiBase(cfg.apiBase);

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', () => {
            initTabs();
            loadProjects('projects-grid', API_BASE);
            loadExperiences(API_BASE);
            initContactForm(API_BASE);
        });
    } else {
        initTabs();
        loadProjects('projects-grid', API_BASE);
        loadExperiences(API_BASE);
        initContactForm(API_BASE);
    }
}

if (!(typeof import.meta !== 'undefined' && import.meta.vitest)) {
    initApp();
}

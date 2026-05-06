import { describe, it, expect, beforeEach } from 'vitest';
import { renderProjects } from '../modules/projects.js';

const mockProjects = [
    {
        title: 'API Gateway',
        description: 'A Go gateway',
        github: 'https://github.com/schulzdimitri',
        tags: ['Golang', 'Redis'],
    },
    {
        title: 'Data Pipeline',
        description: 'A Java pipeline',
        github: 'https://github.com/schulzdimitri',
        tags: ['Java', 'Kafka'],
    },
];

describe('renderProjects', () => {
    let container;

    beforeEach(() => {
        document.body.innerHTML = '<div id="projects-grid"></div>';
        container = document.getElementById('projects-grid');
    });

    it('renders the correct number of project cards', () => {
        renderProjects(mockProjects, container);
        expect(container.querySelectorAll('.project-card').length).toBe(2);
    });

    it('renders each project title', () => {
        renderProjects(mockProjects, container);
        expect(container.innerHTML).toContain('API Gateway');
        expect(container.innerHTML).toContain('Data Pipeline');
    });

    it('renders all tags for each project', () => {
        renderProjects(mockProjects, container);
        expect(container.innerHTML).toContain('Golang');
        expect(container.innerHTML).toContain('Redis');
        expect(container.innerHTML).toContain('Java');
        expect(container.innerHTML).toContain('Kafka');
    });

    it('renders github links', () => {
        renderProjects(mockProjects, container);
        const links = container.querySelectorAll('a.project-card-link');
        expect(links.length).toBe(2);
        expect(links[0].getAttribute('href')).toBe('https://github.com/schulzdimitri');
    });

    it('renders empty string for empty projects array', () => {
        renderProjects([], container);
        expect(container.innerHTML).toBe('');
    });
});

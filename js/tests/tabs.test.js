import { describe, it, expect, beforeEach } from 'vitest';
import { initTabs } from '../modules/tabs.js';

const HTML = `
    <button class="tab-btn active" data-target="job1">Company A</button>
    <button class="tab-btn" data-target="job2">Company B</button>
    <button class="tab-btn" data-target="job3">Company C</button>
    <div class="tab-pane active" id="job1">Content 1</div>
    <div class="tab-pane" id="job2">Content 2</div>
    <div class="tab-pane" id="job3">Content 3</div>
`;

describe('initTabs', () => {
    beforeEach(() => {
        document.body.innerHTML = HTML;
        initTabs();
    });

    it('activates the clicked pane and deactivates others', () => {
        document.querySelector('[data-target="job2"]').click();

        expect(document.getElementById('job2').classList.contains('active')).toBe(true);
        expect(document.getElementById('job1').classList.contains('active')).toBe(false);
        expect(document.getElementById('job3').classList.contains('active')).toBe(false);
    });

    it('marks the clicked button as active', () => {
        document.querySelector('[data-target="job3"]').click();

        expect(document.querySelector('[data-target="job3"]').classList.contains('active')).toBe(true);
    });

    it('removes active from previously active button', () => {
        document.querySelector('[data-target="job2"]').click();

        expect(document.querySelector('[data-target="job1"]').classList.contains('active')).toBe(false);
    });

    it('clicking the already-active tab keeps it active', () => {
        document.querySelector('[data-target="job1"]').click();

        expect(document.getElementById('job1').classList.contains('active')).toBe(true);
    });
});

import { describe, it, expect, beforeEach, vi } from 'vitest';
import { submitContact } from '../modules/contact.js';

const FORM_HTML = `
    <form id="contact-form">
        <input name="name" value="Dimitri">
        <input name="email" value="d@example.com">
        <textarea name="message">Hello!</textarea>
        <button id="contact-submit">
            <span id="submit-label">Enviar mensagem</span>
        </button>
        <p id="form-feedback" class="form-feedback"></p>
    </form>
`;

describe('submitContact', () => {
    let form;

    beforeEach(() => {
        document.body.innerHTML = FORM_HTML;
        form = document.getElementById('contact-form');
        vi.restoreAllMocks();
    });

    it('shows success message and resets form on 2xx response', async () => {
        global.fetch = vi.fn().mockResolvedValue({ ok: true });

        await submitContact(form, 'http://localhost:8080');

        const feedback = document.getElementById('form-feedback');
        expect(feedback.classList.contains('success')).toBe(true);
        expect(feedback.textContent).toContain('✓');
    });

    it('shows error message on non-2xx response', async () => {
        global.fetch = vi.fn().mockResolvedValue({
            ok: false,
            json: async () => ({ error: 'name is required' }),
        });

        await submitContact(form, 'http://localhost:8080');

        const feedback = document.getElementById('form-feedback');
        expect(feedback.classList.contains('error')).toBe(true);
        expect(feedback.textContent).toContain('name is required');
    });

    it('shows error on network failure', async () => {
        global.fetch = vi.fn().mockRejectedValue(new Error('Network error'));

        await submitContact(form, 'http://localhost:8080');

        const feedback = document.getElementById('form-feedback');
        expect(feedback.classList.contains('error')).toBe(true);
        expect(feedback.textContent).toContain('✗');
    });

    it('disables button during submission and re-enables after', async () => {
        let resolveRequest;
        global.fetch = vi.fn().mockReturnValue(
            new Promise(resolve => { resolveRequest = resolve; })
        );

        const promise = submitContact(form, 'http://localhost:8080');
        expect(document.getElementById('contact-submit').disabled).toBe(true);

        resolveRequest({ ok: true });
        await promise;
        expect(document.getElementById('contact-submit').disabled).toBe(false);
    });

    it('sends payload with trimmed values', async () => {
        global.fetch = vi.fn().mockResolvedValue({ ok: true });
        form.elements['name'].value = '  Dimitri  ';
        form.elements['email'].value = '  d@example.com  ';

        await submitContact(form, 'http://localhost:8080');

        const body = JSON.parse(global.fetch.mock.calls[0][1].body);
        expect(body.name).toBe('Dimitri');
        expect(body.email).toBe('d@example.com');
    });
});

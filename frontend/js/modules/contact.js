/**
 * Initializes the contact form event listener.
 * @param {string} apiBase - Base URL of the backend API
 */
export function initContactForm(apiBase) {
    const form = document.getElementById('contact-form');
    if (!form) return;

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        await submitContact(form, apiBase);
    });
}

/**
 * Submits the contact form data to the API.
 * Exported separately to allow direct testing without DOM event setup.
 * @param {HTMLFormElement} form
 * @param {string} apiBase
 */
export async function submitContact(form, apiBase) {
    const submitBtn = document.getElementById('contact-submit');
    const label = document.getElementById('submit-label');
    const feedback = document.getElementById('form-feedback');

    submitBtn.disabled = true;
    label.textContent = 'Enviando...';
    feedback.textContent = '';
    feedback.className = 'form-feedback';

    const payload = {
        name: form.elements['name'].value.trim(),
        email: form.elements['email'].value.trim(),
        message: form.elements['message'].value.trim(),
    };

    try {
        const res = await fetch(`${apiBase}/api/contact`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload),
        });

        if (res.ok) {
            feedback.textContent = '✓ Mensagem enviada! Em breve entrarei em contato.';
            feedback.classList.add('success');
            form.reset();
        } else {
            const body = await res.json();
            feedback.textContent = `✗ Erro: ${body.error || 'tente novamente.'}`;
            feedback.classList.add('error');
        }
    } catch {
        feedback.textContent = '✗ Não foi possível conectar ao servidor.';
        feedback.classList.add('error');
    } finally {
        submitBtn.disabled = false;
        label.textContent = 'Enviar mensagem';
    }
}

const API_BASE = document.querySelector('meta[name="api-base"]')?.content || 'http://localhost:8080';

const FOLDER_ICON = `<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>`;

const GITHUB_ICON = `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path></svg>`;

document.addEventListener('DOMContentLoaded', () => {
    initTabs();
    loadProjects();
    initContactForm();
});

function initTabs() {
    const tabButtons = document.querySelectorAll('.tab-btn');
    const tabPanes = document.querySelectorAll('.tab-pane');

    tabButtons.forEach(button => {
        button.addEventListener('click', () => {
            tabButtons.forEach(btn => btn.classList.remove('active'));
            tabPanes.forEach(pane => pane.classList.remove('active'));
            button.classList.add('active');
            document.getElementById(button.getAttribute('data-target')).classList.add('active');
        });
    });
}

async function loadProjects() {
    const grid = document.getElementById('projects-grid');
    if (!grid) return;

    try {
        const res = await fetch('data/portfolio.json');
        const data = await res.json();

        grid.innerHTML = data.projects.map(project => `
            <div class="project-card">
                <div class="project-card-header">
                    ${FOLDER_ICON}
                    <a href="${project.github}" target="_blank" class="project-card-link" aria-label="GitHub">
                        ${GITHUB_ICON}
                    </a>
                </div>
                <h3>${project.title}</h3>
                <p>${project.description}</p>
                <ul class="project-tags">
                    ${project.tags.map(tag => `<li>${tag}</li>`).join('')}
                </ul>
            </div>
        `).join('');
    } catch {
        grid.innerHTML = '<p style="color:var(--text-muted)">Projetos indisponíveis no momento.</p>';
    }
}

async function initContactForm() {
    const form = document.getElementById('contact-form');
    if (!form) return;

    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        const submitBtn = document.getElementById('contact-submit');
        const label = document.getElementById('submit-label');
        const feedback = document.getElementById('form-feedback');

        submitBtn.disabled = true;
        label.textContent = 'Enviando...';
        feedback.textContent = '';
        feedback.className = 'form-feedback';

        const payload = {
            name: form.name.value.trim(),
            email: form.email.value.trim(),
            message: form.message.value.trim(),
        };

        try {
            const res = await fetch(`${API_BASE}/api/contact`, {
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
    });
}

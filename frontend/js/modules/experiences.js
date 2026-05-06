export async function loadExperiences(apiBase) {
    const tabsList = document.getElementById('experiences-tabs');
    const tabsContent = document.getElementById('experiences-content');

    if (!tabsList || !tabsContent) return;

    try {
        const response = await fetch(`${apiBase}/api/experiences`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const experiences = await response.json();

        if (!experiences || experiences.length === 0) {
            tabsList.innerHTML = '<p class="form-feedback" style="display:block;">Nenhuma experiência cadastrada ainda.</p>';
            return;
        }

        tabsList.innerHTML = '';
        tabsContent.innerHTML = '';

        experiences.forEach((exp, index) => {
            const isActive = index === 0 ? 'active' : '';
            const tabId = `job-${exp.id}`;

            const btn = document.createElement('button');
            btn.className = `tab-btn ${isActive}`;
            btn.setAttribute('data-target', tabId);
            btn.textContent = exp.company;
            tabsList.appendChild(btn);

            const pane = document.createElement('div');
            pane.className = `tab-pane ${isActive}`;
            pane.id = tabId;

            let dutiesHtml = '';
            if (exp.duties && exp.duties.length > 0) {
                dutiesHtml = '<ul class="job-duties">';
                exp.duties.forEach(duty => {
                    dutiesHtml += `<li>${duty}</li>`;
                });
                dutiesHtml += '</ul>';
            }

            pane.innerHTML = `
                <h3 class="job-title">${exp.role} <span class="highlight">@ ${exp.company}</span></h3>
                <p class="job-dates">${exp.period}</p>
                ${dutiesHtml}
            `;
            tabsContent.appendChild(pane);
        });

        initDynamicTabs();

    } catch (error) {
        console.error('Error fetching experiences:', error);
        tabsList.innerHTML = '<p class="form-feedback" style="display:block; color:var(--error-color);">Erro ao carregar experiências.</p>';
    }
}

function initDynamicTabs() {
    const buttons = document.querySelectorAll('#experiences-tabs .tab-btn');
    const panes = document.querySelectorAll('#experiences-content .tab-pane');

    buttons.forEach(btn => {
        btn.addEventListener('click', () => {
            buttons.forEach(b => b.classList.remove('active'));
            panes.forEach(p => p.classList.remove('active'));

            btn.classList.add('active');

            const targetId = btn.getAttribute('data-target');
            document.getElementById(targetId)?.classList.add('active');
        });
    });
}

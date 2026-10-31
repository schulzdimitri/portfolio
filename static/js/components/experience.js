import { portfolioApi } from '../api.js';
import { portfolioData } from '../data/portfolioData.js';

function buildExperienceTimeline(experiences) {
    const timeline = document.createElement('div');
    timeline.className = 'experience-timeline';

    experiences.forEach((item) => {
        const itemCard = document.createElement('article');
        itemCard.className = 'experience-card';

        const header = document.createElement('header');
        header.className = 'experience-header';

        const roleWrapper = document.createElement('div');
        const roleTitle = document.createElement('h3');
        roleTitle.className = 'experience-role';
        roleTitle.textContent = item.role;

        const company = document.createElement('p');
        company.className = 'experience-company';
        company.textContent = item.company;

        roleWrapper.append(roleTitle, company);

        const periodBadge = document.createElement('span');
        periodBadge.className = 'experience-period';
        periodBadge.textContent = item.period;

        header.append(roleWrapper, periodBadge);

        const desc = document.createElement('p');
        desc.className = 'experience-desc';
        desc.textContent = item.description;

        const techList = document.createElement('ul');
        techList.className = 'experience-tech';

        (item.technologies || []).forEach((tech) => {
            const techTag = document.createElement('li');
            techTag.className = 'tech-pill';
            techTag.textContent = tech;
            techList.appendChild(techTag);
        });

        itemCard.append(header, desc, techList);
        timeline.appendChild(itemCard);
    });

    return timeline;
}

export async function renderExperience(container, fallbackExperiences = portfolioData.experiences) {
    if (!container) return;

    try {
        const response = await portfolioApi.getExperiences();
        const items = response.attributes || response.data || [];
        const experiences = items.length > 0 ? items : fallbackExperiences;
        if (Array.isArray(experiences)) {
            container.replaceChildren(buildExperienceTimeline(experiences));
        }
    } catch (error) {
        if (Array.isArray(fallbackExperiences)) {
            container.replaceChildren(buildExperienceTimeline(fallbackExperiences));
        }
    }
}

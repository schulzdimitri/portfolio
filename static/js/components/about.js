export function renderAbout(container, data) {
    if (!container || !data) return;

    const section = document.createElement('div');
    section.className = 'about-wrapper';

    const summaryPara = document.createElement('p');
    summaryPara.className = 'section-description';
    summaryPara.textContent = data.summary;

    const pillarsContainer = document.createElement('div');
    pillarsContainer.className = 'pillars-grid';

    data.pillars.forEach((pillar) => {
        const pillarCard = document.createElement('article');
        pillarCard.className = 'pillar-card';

        const title = document.createElement('h3');
        title.textContent = pillar.title;

        const desc = document.createElement('p');
        desc.textContent = pillar.description;

        pillarCard.append(title, desc);
        pillarsContainer.appendChild(pillarCard);
    });

    const skillsWrapper = document.createElement('div');
    skillsWrapper.className = 'skills-section';

    const skillsTitle = document.createElement('h3');
    skillsTitle.className = 'skills-heading';
    skillsTitle.textContent = 'Technologies & Skills';

    const skillsList = document.createElement('ul');
    skillsList.className = 'skills-tags';

    data.skills.forEach((skill) => {
        const tag = document.createElement('li');
        tag.className = 'skill-tag';
        tag.textContent = skill;
        skillsList.appendChild(tag);
    });

    skillsWrapper.append(skillsTitle, skillsList);
    section.append(summaryPara, pillarsContainer, skillsWrapper);

    container.replaceChildren(section);
}

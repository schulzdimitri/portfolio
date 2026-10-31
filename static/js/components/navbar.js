export function initNavbar() {
    const scrollTriggers = document.querySelectorAll('[data-target]');

    if (window.location.hash) {
        window.history.replaceState(
            null,
            '',
            window.location.pathname + window.location.search
        );
    }

    scrollTriggers.forEach((trigger) => {
        trigger.addEventListener('click', (event) => {
            const targetId = trigger.getAttribute('data-target');
            if (!targetId) return;

            const targetSection = document.getElementById(targetId);
            if (targetSection) {
                event.preventDefault();
                if (targetId === 'hero') {
                    window.scrollTo({ top: 0, behavior: 'smooth' });
                } else {
                    targetSection.scrollIntoView({ behavior: 'smooth' });
                }
            }
        });
    });
}

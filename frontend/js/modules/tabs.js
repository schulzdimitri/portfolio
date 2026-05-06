/**
 * Initializes the vertical tabs in the Experience section.
 * @param {string} btnSelector - CSS selector for tab buttons
 * @param {string} paneSelector - CSS selector for tab panes
 */
export function initTabs(
    btnSelector = '.tab-btn',
    paneSelector = '.tab-pane'
) {
    const buttons = document.querySelectorAll(btnSelector);
    const panes = document.querySelectorAll(paneSelector);

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

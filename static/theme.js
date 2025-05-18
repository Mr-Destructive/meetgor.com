document.addEventListener('DOMContentLoaded', () => {
    const themeToggle = document.getElementById('theme-toggle');
    const body = document.body;

    const currentTheme = localStorage.getItem('theme') || 'light';

    if (currentTheme === 'secondary') {
        body.classList.add('secondary-theme');
    }

    themeToggle.addEventListener('click', () => {
        if (body.classList.contains('secondary-theme')) {
            body.classList.remove('secondary-theme');
            localStorage.setItem('theme', 'light');
        } else {
            body.classList.add('secondary-theme');
            localStorage.setItem('theme', 'secondary');
        }
    });
});

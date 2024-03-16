window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag('js', new Date());
gtag('config', 'G-EET2ZX4QY1');

function setTheme(theme) {
    document.documentElement.setAttribute("data-theme", theme);
}

function detectColorSchemeOnLoad() {
    if (localStorage.getItem("theme")) {
        if (localStorage.getItem("theme") == "dark") {
            setTheme('dark')
        } else if (localStorage.getItem("theme") == "light") {
            setTheme('light')
        }
    } else if (!window.matchMedia) {
        //matchMedia method not supported
        setTheme('light')
        return false;
    } else if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        //OS theme setting detected as dark
        setTheme('dark')
    } else {
        setTheme('light')
    }
}
detectColorSchemeOnLoad();
document.addEventListener('DOMContentLoaded', function() {
    const toggleSwitch = document.querySelector('#theme-switch input[type="checkbox"]');

    function switchTheme(e) {
        if (e.target.checked) {
            localStorage.setItem('theme', 'dark');
            document.documentElement.setAttribute('data-theme', 'dark');
            toggleSwitch.checked = true;
        } else {
            localStorage.setItem('theme', 'light');
            document.documentElement.setAttribute('data-theme', 'light');
            toggleSwitch.checked = false;
        }
    }

    toggleSwitch.addEventListener('change', switchTheme, false);

    if (document.documentElement.getAttribute("data-theme") == "dark") {
        toggleSwitch.checked = true;
    }

}, false);

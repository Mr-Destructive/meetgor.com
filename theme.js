(function () {
  const STORAGE_KEY = "theme";
  const DARK_THEME = "secondary";

  function normalizeTheme(value) {
    const theme = String(value || "").toLowerCase();
    if (theme === "secondary" || theme === "secondary-theme" || theme === "dark" || theme === "dark_high_contrast") {
      return DARK_THEME;
    }
    return "default";
  }

  function setSyntaxTheme(theme) {
    const stylesheet = document.getElementById("syntax-theme");
    if (!stylesheet) {
      return;
    }

    stylesheet.href = theme === DARK_THEME
      ? "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/github-dark.min.css"
      : "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/github.min.css";
  }

  function applyTheme(theme, toggle) {
    const normalized = normalizeTheme(theme);
    const isDark = normalized === DARK_THEME;

    document.body.classList.toggle("secondary-theme", isDark);
    document.documentElement.classList.toggle("secondary-theme", isDark);

    if (toggle) {
      toggle.checked = isDark;
    }

    setSyntaxTheme(normalized);
    syncGiscusTheme(normalized);
  }

  function syncGiscusTheme(theme) {
    const iframe = document.querySelector("iframe.giscus-frame");
    if (!iframe || !iframe.contentWindow) {
      return;
    }

    const giscusTheme = theme === DARK_THEME ? "dark_high_contrast" : "light";
    iframe.contentWindow.postMessage(
      {
        giscus: {
          setConfig: {
            theme: giscusTheme,
          },
        },
      },
      "https://giscus.app"
    );
  }

  document.addEventListener("DOMContentLoaded", () => {
    const themeToggle = document.getElementById("theme-toggle");
    const savedTheme = normalizeTheme(localStorage.getItem(STORAGE_KEY));

    localStorage.setItem(STORAGE_KEY, savedTheme);
    applyTheme(savedTheme, themeToggle);

    if (!themeToggle) {
      return;
    }

    themeToggle.addEventListener("change", () => {
      const nextTheme = themeToggle.checked ? DARK_THEME : "default";
      localStorage.setItem(STORAGE_KEY, nextTheme);
      applyTheme(nextTheme, themeToggle);
    });
  });
})();

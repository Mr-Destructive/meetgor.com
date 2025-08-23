const runQueryBtn = document.getElementById('run-query');
const resultsContainer = document.getElementById('results-container');

// Initialize CodeMirror
const editor = CodeMirror(document.getElementById('sql-editor'), {
    mode: 'text/x-sql',
    theme: 'material-darker',
    lineNumbers: true,
    value: "SELECT title, date, tags FROM all_posts ORDER BY date DESC LIMIT 10;"
});

async function initDb() {
    const SQL = await initSqlJs({
        locateFile: file => `https://cdnjs.cloudflare.com/ajax/libs/sql.js/1.10.3/${file}`
    });

    const db = new SQL.Database();

    // Create tables
    db.run("CREATE TABLE posts (title TEXT, date TEXT, tags TEXT, content TEXT, description TEXT, series TEXT, status TEXT, slug TEXT);");
    db.run("CREATE TABLE tils (title TEXT, date TEXT, tags TEXT, content TEXT, description TEXT, series TEXT, status TEXT, slug TEXT);");
    db.run("CREATE TABLE newsletter (title TEXT, date TEXT, tags TEXT, content TEXT, description TEXT, series TEXT, status TEXT, slug TEXT);");
    db.run("CREATE TABLE work (title TEXT, date TEXT, tags TEXT, content TEXT, description TEXT, series TEXT, status TEXT, slug TEXT);");
    db.run("CREATE TABLE projects (title TEXT, date TEXT, tags TEXT, content TEXT, description TEXT, series TEXT, status TEXT, slug TEXT);");
    db.run("CREATE TABLE all_posts (title TEXT, date TEXT, tags TEXT, content TEXT, description TEXT, series TEXT, status TEXT, slug TEXT);");

    // Fetch and load data
    const response = await fetch('/posts.json');
    const data = await response.json();

    // Helper to insert posts into a table
    const insertPosts = (tableName, posts) => {
        const stmt = db.prepare(`INSERT INTO ${tableName} VALUES (?, ?, ?, ?, ?, ?, ?, ?)`);
        posts.forEach(post => {
            const tags = Array.isArray(post.tags) ? post.tags.join(',') : '';
            stmt.run([post.title, post.date, tags, post.content, post.description, post.series, post.status, post.slug]);
        });
        stmt.free();
    };

    if (data.posts) insertPosts('posts', data.posts);
    if (data.tils) insertPosts('tils', data.tils);
    if (data.newsletter) insertPosts('newsletter', data.newsletter);
    if (data.work) insertPosts('work', data.work);
    if (data.projects) insertPosts('projects', data.projects);
    if (data.all_posts) insertPosts('all_posts', data.all_posts);

    return db;
}

let dbPromise = initDb();

runQueryBtn.addEventListener('click', async () => {
    const db = await dbPromise;
    const query = editor.getValue();
    try {
        const results = db.exec(query);
        renderResults(results);
    } catch (e) {
        resultsContainer.innerHTML = `<p class="text-red-500">Error: ${e.message}</p>`;
    }
});

function renderResults(results) {
    resultsContainer.innerHTML = '';
    if (results.length === 0) {
        resultsContainer.innerHTML = '<p>No results</p>';
        return;
    }

    results.forEach(result => {
        const table = document.createElement('table');
        table.classList.add('min-w-full', 'bg-white', 'border', 'border-gray-300');
        const thead = document.createElement('thead');
        const tbody = document.createElement('tbody');

        // Table Headers
        const headerRow = document.createElement('tr');
        result.columns.forEach(col => {
            const th = document.createElement('th');
            th.classList.add('py-2', 'px-4', 'border-b', 'bg-gray-100', 'font-bold', 'text-left');
            th.textContent = col;
            headerRow.appendChild(th);
        });
        thead.appendChild(headerRow);

        // Table Rows
        result.values.forEach(row => {
            const tr = document.createElement('tr');
            tr.classList.add('hover:bg-gray-50');
            row.forEach(val => {
                const td = document.createElement('td');
                td.classList.add('py-2', 'px-4', 'border-b', 'border-gray-200');
                td.textContent = val;
                tr.appendChild(td);
            });
            tbody.appendChild(tr);
        });

        table.appendChild(thead);
        table.appendChild(tbody);
        resultsContainer.appendChild(table);
    });
}

// Check for query in URL on page load
window.addEventListener('load', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const query = urlParams.get('query');
    if (query) {
        editor.setValue(query);
        const db = await dbPromise;
        try {
            const results = db.exec(query);
            renderResults(results);
        } catch (e) {
            resultsContainer.innerHTML = `<p class="text-red-500">Error: ${e.message}</p>`;
        }
    }
});
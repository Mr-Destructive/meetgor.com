document.addEventListener('DOMContentLoaded', async () => {
    const playgrounds = document.querySelectorAll('.sql-playground');

    const SQL = await initSqlJs({
        locateFile: file => `https://cdnjs.cloudflare.com/ajax/libs/sql.js/1.6.2/${file}`
    });

    playgrounds.forEach(playground => {
        const runButton = playground.querySelector('.run-button');
        const editButton = playground.querySelector('.edit-button');
        const resetButton = playground.querySelector('.reset-button');
        let codeContainer = playground.querySelector('pre');
        let codeBlock = codeContainer.querySelector('code');
        const outputDiv = playground.querySelector('.output');
        const statusDiv = document.createElement('div');
        statusDiv.className = 'status';
        playground.insertBefore(statusDiv, outputDiv);

        let originalCode = codeBlock.innerText;
        let editor = null;
        let db = new SQL.Database();

        const getCurrentTheme = () => {
            return document.body.classList.contains('secondary-theme') ? 'monokai' : 'default';
        };

        runButton.addEventListener('click', async () => {
            const queries = editor ? editor.getValue().split(';') : codeBlock.innerText.split(';');
            outputDiv.innerHTML = '';
            outputDiv.style.display = 'block';
            statusDiv.innerHTML = '<div class="spinner"></div>';

            let allSucceeded = true;
            let firstError = null;

            for (const query of queries) {
                if (query.trim() === '') continue;

                try {
                    const results = db.exec(query);
                    if (results.length > 0) {
                        const table = document.createElement('table');
                        const thead = document.createElement('thead');
                        const tbody = document.createElement('tbody');
                        const headers = results[0].columns;

                        const headerRow = document.createElement('tr');
                        headers.forEach(header => {
                            const th = document.createElement('th');
                            th.innerText = header;
                            headerRow.appendChild(th);
                        });
                        thead.appendChild(headerRow);

                        results[0].values.forEach(rowData => {
                            const row = document.createElement('tr');
                            rowData.forEach(cellData => {
                                const td = document.createElement('td');
                                td.innerText = cellData;
                                row.appendChild(td);
                            });
                            tbody.appendChild(row);
                        });

                        table.appendChild(thead);
                        table.appendChild(tbody);
                        outputDiv.appendChild(table);
                    }
                } catch (error) {
                    allSucceeded = false;
                    if (!firstError) {
                        firstError = error.message;
                    }
                }
            }

            if (allSucceeded) {
                statusDiv.innerHTML = '<span class="success">&#10004;</span> All queries executed successfully.';
            } else {
                statusDiv.innerHTML = `<span class="error">&#10006;</span> Error: ${firstError}`;
            }
        });

        const openEditor = () => {
            if (editor) return;
            editor = CodeMirror(elt => {
                playground.replaceChild(elt, codeContainer);
            }, {
                value: codeBlock.innerText,
                mode: 'sql',
                theme: getCurrentTheme(),
                lineNumbers: true,
                viewportMargin: Infinity
            });
            editButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x-circle"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>';
        };

        const closeEditor = () => {
            if (!editor) return;
            const newPre = document.createElement('pre');
            const newCodeBlock = document.createElement('code');
            newCodeBlock.className = 'language-sql';
            newCodeBlock.innerText = editor.getValue();
            newPre.appendChild(newCodeBlock);
            playground.replaceChild(newPre, playground.querySelector('.CodeMirror'));
            codeContainer = newPre;
            codeBlock = newCodeBlock;
            hljs.highlightBlock(codeBlock);
            editButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-edit"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path></svg>';
            editor = null;
        };

        codeContainer.addEventListener('click', openEditor);

        editButton.addEventListener('click', () => {
            if (editor) {
                closeEditor();
            } else {
                openEditor();
            }
        });

        resetButton.addEventListener('click', () => {
            if (editor) {
                editor.setValue(originalCode);
            } else {
                codeBlock.innerText = originalCode;
            }
            db = new SQL.Database();
            outputDiv.style.display = 'none';
        });

        const themeToggle = document.getElementById('theme-toggle');
        if (themeToggle) {
            themeToggle.addEventListener('change', () => {
                if (editor) {
                    editor.setOption('theme', getCurrentTheme());
                }
            });
        }
    });

    const codeBlocks = document.querySelectorAll('.code-block');
    codeBlocks.forEach(codeBlock => {
        const copyButton = codeBlock.querySelector('.copy-button');
        const code = codeBlock.querySelector('code');
        copyButton.addEventListener('click', () => {
            navigator.clipboard.writeText(code.innerText);
            copyButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-check"><polyline points="20 6 9 17 4 12"></polyline></svg>';
            setTimeout(() => {
                copyButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>';
            }, 2000);
        });
    });

    const headings = document.querySelectorAll('h1, h2, h3, h4, h5, h6');
    headings.forEach(heading => {
        const id = heading.getAttribute('id');
        if (id) {
            const anchor = document.createElement('a');
            anchor.className = 'anchor-link';
            anchor.href = `#${id}`;
            anchor.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.72"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.72-1.72"></path></svg>';
            heading.appendChild(anchor);
        }
    });
});
document.addEventListener('DOMContentLoaded', async () => {
    const playgrounds = document.querySelectorAll('.sql-playground');

    // Update to sql.js 1.13.0
    const SQL = await initSqlJs({
        locateFile: file => `https://cdnjs.cloudflare.com/ajax/libs/sql.js/1.13.0/${file}`
    });

    // Create a single shared database instance for all playgrounds on the page
    let sharedDB = new SQL.Database();

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

        // Create copy button
        const copyButton = document.createElement('button');
        copyButton.className = 'copy-button';
        copyButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg> Copy';
        
        // Create info div for SQL dialect
        const infoDiv = document.createElement('div');
        infoDiv.className = 'info';
        infoDiv.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg> SQLite 3.49';

        // Add copy button and info to controls
        const controls = playground.querySelector('.controls');
        controls.appendChild(copyButton);
        controls.appendChild(infoDiv);

        // Updated theme function to be consistent with the site's dynamic theme
        const getCurrentTheme = () => {
            return document.body.classList.contains('secondary-theme') ? 'monokai' : 'default';
        };

        // Updated theme function for consistent UI with the site
        const updateUITheme = () => {
            const isDarkTheme = document.body.classList.contains('secondary-theme');
            if (editor) {
                editor.setOption('theme', isDarkTheme ? 'monokai' : 'default');
            }
            
            // Update playground UI elements to match theme
            playground.style.borderColor = isDarkTheme ? 'var(--border-color)' : 'var(--border-color)';
            statusDiv.style.color = isDarkTheme ? 'var(--secondary-text-color)' : 'var(--text-color)';
            outputDiv.style.borderColor = isDarkTheme ? 'var(--border-color)' : 'var(--border-color)';
        };

        // Copy button functionality
        copyButton.addEventListener('click', () => {
            const codeToCopy = editor ? editor.getValue() : codeBlock.innerText;
            navigator.clipboard.writeText(codeToCopy).then(() => {
                const originalHTML = copyButton.innerHTML;
                copyButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg> Copied!';
                setTimeout(() => {
                    copyButton.innerHTML = originalHTML;
                }, 2000);
            });
        });

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
                    // Use the shared database instance
                    const results = sharedDB.exec(query);
                    if (results.length > 0) {
                        const table = document.createElement('table');
                        table.className = 'sql-result-table';
                        const thead = document.createElement('thead');
                        const tbody = document.createElement('tbody');
                        const headers = results[0].columns;

                        const headerRow = document.createElement('tr');
                        headers.forEach(header => {
                            const th = document.createElement('th');
                            th.innerText = header;
                            th.className = 'sql-result-header';
                            headerRow.appendChild(th);
                        });
                        thead.appendChild(headerRow);

                        results[0].values.forEach(rowData => {
                            const row = document.createElement('tr');
                            rowData.forEach(cellData => {
                                const td = document.createElement('td');
                                td.innerText = cellData;
                                td.className = 'sql-result-cell';
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
                statusDiv.innerHTML = '<span class="success">✓</span> All queries executed successfully.';
            } else {
                statusDiv.innerHTML = `<span class="error">✖</span> Error: ${firstError}`;
            }
            
            updateUITheme();
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
            editButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg> Close';
            updateUITheme();
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
            editButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path></svg> Edit';
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
            // Reset the shared database to a fresh instance
            sharedDB = new SQL.Database();
            outputDiv.style.display = 'none';
            statusDiv.innerHTML = '<span class="success">✓</span> Database reset to initial state';
            setTimeout(() => {
                statusDiv.innerHTML = '';
            }, 3000);
            updateUITheme();
        });

        const themeToggle = document.getElementById('theme-toggle');
        if (themeToggle) {
            themeToggle.addEventListener('change', () => {
                updateUITheme();
            });
        }
        
        // Initial theme update
        updateUITheme();
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

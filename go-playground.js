document.addEventListener('DOMContentLoaded', async () => {
    const playgrounds = document.querySelectorAll('.go-playground');

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

        runButton.addEventListener('click', async () => {
            const code = editor ? editor.getValue() : codeBlock.innerText;
            outputDiv.innerHTML = '';
            outputDiv.style.display = 'block';
            statusDiv.innerHTML = '<div class="spinner"></div>';
            
            try {
                // Execute Go code simulation
                const result = await executeGoCode(code);
                outputDiv.innerHTML = `<pre>${escapeHtml(result.output)}</pre>`;
                statusDiv.innerHTML = `<span class="success">✓</span> Code executed in ${result.duration}ms.`;
            } catch (error) {
                outputDiv.innerHTML = `<pre class="error">${escapeHtml(error.message)}</pre>`;
                statusDiv.innerHTML = '<span class="error">✖</span> Error during execution.';
            }
            
            updateUITheme();
        });

        const openEditor = () => {
            if (editor) return;
            editor = CodeMirror(elt => {
                playground.replaceChild(elt, codeContainer);
            }, {
                value: codeBlock.innerText,
                mode: 'go',
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
            newCodeBlock.className = 'language-go';
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
            outputDiv.style.display = 'none';
            statusDiv.innerHTML = '<span class="success">✓</span> Code reset to initial state';
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

    // Function to execute Go code (simulation)
    async function executeGoCode(code) {
        return new Promise((resolve, reject) => {
            try {
                // Simulate compilation and execution time
                const startTime = performance.now();
                
                // Simple Go code interpreter simulation
                let output = '';
                let hasError = false;
                
                // Check for basic syntax errors
                if (!code.includes('package main')) {
                    throw new Error('Missing package declaration: package main');
                }
                
                if (!code.includes('func main()')) {
                    throw new Error('Missing main function: func main()');
                }
                
                // Simulate execution based on code patterns
                if (code.includes('fmt.Println')) {
                    const printMatches = code.match(/fmt\.Println\s*\(([^)]+)\)/g);
                    if (printMatches) {
                        printMatches.forEach(match => {
                            const content = match.match(/fmt\.Println\s*\(([^)]+)\)/)[1];
                            // Handle string literals
                            if (content.startsWith('"') && content.endsWith('"')) {
                                output += content.slice(1, -1) + '\n';
                            } else {
                                // Handle variables and expressions
                                output += content + '\n';
                            }
                        });
                    }
                }
                
                if (code.includes('fmt.Printf')) {
                    const printfMatches = code.match(/fmt\.Printf\s*\(([^)]+)\)/g);
                    if (printfMatches) {
                        printfMatches.forEach(match => {
                            const content = match.match(/fmt\.Printf\s*\(([^)]+)\)/)[1];
                            // Simplified handling
                            if (content.startsWith('"') && content.endsWith('"')) {
                                output += content.slice(1, -1) + '\n';
                            }
                        });
                    }
                }
                
                // Simulate goroutines
                if (code.includes('go ')) {
                    output += "// Goroutines would run concurrently in a real implementation\n";
                }
                
                // Simulate time operations
                if (code.includes('time.Sleep')) {
                    output += "// Sleep operation would pause execution in a real implementation\n";
                }
                
                const endTime = performance.now();
                const duration = Math.round(endTime - startTime);
                
                resolve({
                    output: output || "// Program executed successfully\n// No output generated",
                    duration: duration
                });
            } catch (error) {
                reject(error);
            }
        });
    }

    // Utility function to escape HTML
    function escapeHtml(unsafe) {
        return unsafe
            .replace(/&/g, "&amp;")
            .replace(/</g, "&lt;")
            .replace(/>/g, "&gt;")
            .replace(/"/g, "&quot;")
            .replace(/'/g, "&#039;");
    }

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
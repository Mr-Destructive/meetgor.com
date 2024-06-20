document.addEventListener("DOMContentLoaded", function() {
    // Get all pre elements
    var preElements = document.querySelectorAll("pre");

    // Iterate over each pre element
    preElements.forEach(function(preElement) {
        // Create a div element for the highlight bar
        var highlightBar = document.createElement("div");
        highlightBar.classList.add("highlight-bar");

        // Create a select element for language selection
        var language = preElement.classList[1];

        // Create a span element for the language
        var languageSpan = document.createElement("span");
        languageSpan.classList.add("language");
        languageSpan.textContent = language;

        // Append the language span to the highlight bar
        highlightBar.appendChild(languageSpan);

        // Append the highlight bar before the pre element
        preElement.parentNode.insertBefore(highlightBar, preElement);

        // Move the pre element inside the copy wrapper
        var copyWrapper = document.createElement("div");
        copyWrapper.classList.add("copy-wrapper");
        copyWrapper.appendChild(preElement.parentNode.removeChild(preElement));

        // Append the copy wrapper before the highlight bar
        highlightBar.parentNode.insertBefore(copyWrapper, highlightBar.nextSibling);

        // Append the copy wrapper before the pre element

        // Create the copy button
        var copyButton = document.createElement("button");
        copyButton.classList.add("copy");
        copyButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>`;

        // Append the copy button to the copy wrapper
        copyWrapper.appendChild(copyButton);

        // Add event listener to the copy button
        copyButton.addEventListener("click", function() {
            // Get the text content of the pre element
            var codeText = preElement.textContent;

            // Create a textarea element
            var tempTextarea = document.createElement("textarea");
            tempTextarea.value = codeText;

            // Append the textarea to the document body
            document.body.appendChild(tempTextarea);

            // Select the textarea content
            tempTextarea.select();
            tempTextarea.setSelectionRange(0, 99999); /* For mobile devices */

            // Copy the selected text to the clipboard
            document.execCommand("copy");

            // Remove the textarea
            document.body.removeChild(tempTextarea);

            // Change copy button text
            copyButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-check"><polyline points="20 6 9 17 4 12"></polyline></svg>`;

            // Reset button text after 2 seconds
            setTimeout(function() {
                copyButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>`;
            }, 2000);
        });
    });
});

# Techstructive Blog

- [Blog](https://meetgor.com/)
- Author: [Meet Gor](https://github.com/mr-destructive)

## A personal blog made with the help of [Markata](https://markata.dev) and GitHub pages.

I write about the things I learn here. Some of the topics for most of the Articles/Blogs are:

- [Django](https://meetgor.com/tag/django)
- [Golang](https://meetgor.com/tag/go)
- [Python](https://meetgor.com/tag/python)
- [Vim](https://meetgor.com/tag/vim)
- And Programming/Development in general

This is a personal blog made with html/css/markdown, generated using Markata(a Python-plugin based SSG) and hosted on GitHub pages.

### Techstructive Blog is created with a aim to learn programming and have fun in the process.

You'll also find some of my articles written at [GeeksforGeeks](https://auth.geeksforgeeks.org/user/meetgor/articles).
They mainly contain shell scripting(bash), Java, Vim and some random articles as well.

I hope you'll have fun reading my articles, if not please give feedback on the handles from [Contact](https://meetgor.com/blog/) section.

## Credits

- [Markata](https://github.com/waylonwalker/markata) - A Python Plugin based Static Site Generator
- [Waylon Walker](https://waylonwalker.com)

## Triggering Manual Builds via Netlify Function

A Netlify function is available to programmatically trigger the GitHub Actions build and deployment process (`.github/workflows/deploy.yml`).

**Endpoint:** `/.netlify/functions/trigger-build`

**Method:** `POST`

**Headers (if `NETLIFY_TRIGGER_SECRET` is configured):**
*   `X-Trigger-Secret`: `<Your_Secret_Value>`

**Request Body:**
No request body is strictly required by the function itself for its primary operation of triggering the build. The GitHub API call will use the `GH_BRANCH` environment variable (or 'main' if not set) to specify the branch.

**Netlify Environment Variables Required:**

*   `GH_OWNER`: Your GitHub username or organization name (e.g., `your-github-username`).
*   `GH_REPO`: Your GitHub repository name (e.g., `your-repo-name`).
*   `GITHUB_PAT_FOR_TRIGGER`: A GitHub Personal Access Token (PAT) with `repo` scope (or at least `workflow` scope if sufficient) to authorize the triggering of GitHub Actions. This token is used by the Netlify function.
*   `NETLIFY_TRIGGER_SECRET` (Optional, but Recommended for Security): A secret string that must be passed in the `X-Trigger-Secret` header of the request to authorize the function call. If this variable is set in Netlify, the function will require the header for authentication.
*   `GH_BRANCH` (Optional): The name of the branch to build. Defaults to `main` if not set.

**GitHub Repository Secrets Required:**

*   `GITHUB_TOKEN`: This is the standard token provided by GitHub Actions, used by the `deploy.yml` workflow itself for operations like deploying to GitHub Pages. Ensure it's available for your actions.

**Example Usage (using cURL):**

If `NETLIFY_TRIGGER_SECRET` is set:
```bash
curl -X POST \
  -H "X-Trigger-Secret: YOUR_SHARED_SECRET" \
  https://your-netlify-site.netlify.app/.netlify/functions/trigger-build
```

If `NETLIFY_TRIGGER_SECRET` is NOT set (less secure):
```bash
curl -X POST \
  https://your-netlify-site.netlify.app/.netlify/functions/trigger-build
```

**Success Response (200 OK):**
```json
{
  "message": "Build triggered successfully."
}
```

**Error Responses:**
*   `401 Unauthorized`: If `NETLIFY_TRIGGER_SECRET` is set and the `X-Trigger-Secret` header is missing or incorrect.
*   `405 Method Not Allowed`: If a method other than POST (or OPTIONS for preflight) is used.
*   `500 Internal Server Error`: If required environment variables are missing on Netlify.
*   `502 Bad Gateway`: If the GitHub API call fails or returns an error.

---
type: static
title: Colophon
slug: colophon
description: Notes on how this site is built and a live map of the content.
---

This site is a hand-rolled static setup focused on speed, clarity, and long-term ownership. I keep the stack lightweight, keep the writing first, and let the structure evolve as the content grows.

## Stack and Build
- Content lives in Markdown with YAML frontmatter, grouped by type inside `posts/` (posts, newsletter, links, projects, thoughts, til, work, and more).
- A custom SSG written in Go (in this repo) reads the Markdown, parses frontmatter, and renders HTML using Go templates.
- The SSG generates per-type feeds and pages, tag/series indexes, RSS feeds, and a sitemap.
- Static pages like About/Contact/Now/Colophon live in `static/` and are rendered through the same templating pipeline.

## Rendering Details
- Go templates in `templates/` define shared layout and per-post/per-page views.
- Content is turned into HTML with Goldmark and a few custom extensions (SQL playground, embeds).
- The build outputs to `public/`, which is the deployable static site.

## Deploy and Hosting
- The site is built as static HTML/CSS/JS and served from `public/`.
- GitHub Pages is used for hosting.
- Netlify functions are used for lightweight automation (CMS actions, sync triggers, build hooks).
- A scheduled sync script can import newsletter posts from Substack into `posts/newsletter/`.

## Notes on Design
- Minimal layout and typography to keep reading fast and distraction‑free.
- A light/dark theme toggle is available, with syntax themes matched to the current mode.
- Seasonal effects are subtle and opt‑in (only after clicking outside the article area).

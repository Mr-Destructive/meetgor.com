# Newsletter Sync Function

This Netlify function fetches the Substack RSS feed and syncs newsletter posts to the blog.

## Features

- Fetches Substack RSS feed with proper User-Agent headers
- Parses RSS/Atom feed items into blog posts
- Generates markdown files with frontmatter
- Sets canonical URL to original Substack link
- Stores posts in `posts/newsletter/` directory with date-based slugs

## Configuration

Set the following environment variables in Netlify:

| Variable | Default | Description |
|----------|---------|-------------|
| `NEWSLETTER_FEED_URL` | `https://techstructively.substack.com/feed` | RSS feed URL to sync from |

## API Endpoint

```
POST /.netlify/functions/newsletter
```

## Response

```json
{
  "message": "Newsletter sync completed",
  "saved": 5,
  "skipped": 0,
  "total": 5
}
```

## Generated Post Structure

Each newsletter post generates a markdown file in `posts/newsletter/` with:

- **Filename**: `YYYY-MM-DD-slug.md` (e.g., `2024-01-15-how-to-build-apis.md`)
- **Frontmatter**:
  - `type`: `newsletter`
  - `title`: From RSS item title
  - `date`: From RSS item published date
  - `canonical_url`: Link to original Substack article
  - `tags`: `["newsletter", "substack"]`
  - `description`: Truncated from RSS description
  - `status`: `published`

## Usage

### Manual Trigger

```bash
curl -X POST https://yourblog.netlify.app/.netlify/functions/newsletter
```

### GitHub Actions Integration

Add to your workflow:

```yaml
- name: Sync Newsletter
  run: |
    curl -X POST https://yourblog.netlify.app/.netlify/functions/newsletter
```

### Scheduled via Netlify

Use Netlify's scheduled functions (if available) or external cron service.

## Notes

- Uses User-Agent headers to bypass Substack's RSS restrictions
- Only fetches feed on demand (not cached)
- Skips items with no title
- Creates files in `/tmp/posts/newsletter` during function execution
- For persistent storage, integrate with GitHub API to commit files back to repo

## Future Enhancements

- [ ] Save files directly to GitHub via API
- [ ] Fallback RSS aggregators if direct feed fails
- [ ] Deduplication based on canonical URL
- [ ] Automatic database sync to Turso

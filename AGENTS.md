# System Architecture & Guidelines

## Overview
Techstructive Blog: Static site (Go SSG) + Turso database + Netlify serverless functions + GitHub workflows. Three data flows sync Git markdown ↔ Turso bidirectionally.

## Tech Stack
- **SSG**: Go (main.go) → renders `/posts/*.md` → `public/`
- **Database**: Turso (libsql) with posts, authors tables
- **Functions**: Netlify (Go Lambda) at `netlify/functions/`
- **Hosting**: GitHub Pages (`gh-pages` branch) + Netlify for functions
- **CI/CD**: GitHub Actions (workflows in `.github/workflows/`)

## Key Components

### 1. Hash Sync (`cmd/hash_sync/`)
- Scans `posts/*.md`, computes SHA-256 hash of title+body
- Rewrites files with metadata (slug, content_hash)
- Insert/update Turso only if hash differs
- **Run**: `go run ./cmd/hash_sync --root posts --author 1`
- **Test**: `go test ./cmd/hash_sync -v`

### 2. Netlify Functions (`netlify/functions/`)
- **posts**: CRUD endpoint (`GET/POST/PUT/DELETE /posts?slug=...`)
- **newsletter**: Weekly Substack RSS → Turso sync
- **cms**: Legacy CMS endpoint
- Auth via bcrypt, env vars: `TURSO_DATABASE_*`

### 3. GitHub Workflows (`.github/workflows/`)
- **hash-sync.yml**: Triggers on `posts/**/*.md` push or daily 2 AM UTC
  - Runs hash_sync, commits rewritten files back to GitHub
  - **Requires secrets**: `TURSO_DATABASE_NAME`, `TURSO_DATABASE_AUTH_TOKEN`
- **deploy.yml**: Manual/webhook trigger for full build & GitHub Pages deploy

## Data Flows (Bidirectional Sync)

1. **GitHub → Turso** (hash-sync workflow)
   - Direct markdown edits on GitHub → hash_sync detects → updates Turso

2. **Editor UI → Turso** (posts function)
   - Web form `/editor` → `/.netlify/functions/posts` → Turso & triggers rebuild

3. **Substack RSS → Turso + GitHub** (newsletter function)
   - Weekly fetch → `posts/newsletter/YYYY-MM-DD-slug.md` + Turso insert

## Database Schema
```
posts: (id, title, slug, body, metadata JSON, author_id, deleted, created_at, updated_at)
authors: (id, username, name, password BCRYPT)
```
DDL in `plugins/db.go` (embedded `db/schema.sql`)

## Testing
- **Local**: `go test ./...` (all packages use in-memory SQLite, no Turso calls)
- **TDD approach**: Integration tests in `cmd/hash_sync/*_test.go`
- **Netlify functions**: Tests mock DB with in-memory sqlite3

## Important Configuration

### netlify.toml
```toml
[build]
  command = "go run main.go"  # Run SSG, outputs to public/
  publish = "public"
[functions]
  directory = "netlify/functions"
[[redirects]]
  /cms → /.netlify/functions/cms
  /posts → /.netlify/functions/posts
```

### Environment Variables (Netlify Secrets)
- `TURSO_DATABASE_NAME`: libsql://[account]-[org].turso.io
- `TURSO_DATABASE_AUTH_TOKEN`: Turso auth token
- `GITHUB_PAT_FOR_TRIGGER`: GitHub Personal Access Token (for webhook dispatch)
- `NEWSLETTER_FEED_URL`: Substack RSS feed URL (default: techstructively.substack.com/feed)

## Critical Paths

1. **Adding a post via editor**: `POST /.netlify/functions/posts` → Turso + triggers GitHub rebuild
2. **GitHub-only edits**: hash-sync detects → rewrites + commits metadata → syncs to Turso
3. **Build on Netlify**: `go run main.go` → `public/` → functions bundled + deployed

## Common Tasks

- **Run full test suite**: `go test ./...`
- **Local SSG build**: `go run main.go` (outputs `public/`)
- **Sync GitHub to Turso**: `TURSO_*=... go run ./cmd/hash_sync --root posts --author 1`
- **Dry-run**: Add `--dry-run` flag to hash_sync
- **Deploy**: Push to `main` → GitHub Actions → builds, syncs, publishes to `gh-pages`

## Gotchas

- **Netlify cache**: Can cause stale builds. If functions fail mysteriously, clear cache in Netlify UI
- **Deploy branch**: Must be `main` (not `gh-pages`). `gh-pages` is read-only output
- **Turso credentials**: Missing secrets → hash_sync + functions fail silently
- **Hash collisions**: Unlikely but SHA-256 ensures posts don't duplicate on re-runs
- **Metadata JSON**: Stored in posts.metadata column. Must stay valid JSON for editor sync

## Monitoring

- GitHub Actions logs: `.github/workflows/` runs visible in repo Actions tab
- Netlify logs: Dashboard → Deploys → click build for function compile errors
- Hash sync: Check git commits for "hash-sync: added=X updated=Y skipped=Z" pattern

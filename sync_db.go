package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
	timestampFile = ".last_build_timestamp"
	timeFormat    = "2006-01-02 15:04:05" // Used for formatting time for SQL query
	rfc3339Format = time.RFC3339          // Used for timestamp file and frontmatter
)

// Post represents a database post record
type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Slug     string `json:"slug"` // Slug from the database
	Body     string `json:"body"`
	Metadata string `json:"metadata"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	AuthorID int64  `json:"author_id"`
	Deleted  bool   `json:"deleted"`
}

// PostMetadata represents the metadata structure
type PostMetadata struct {
	Title       string   `json:"title,omitempty"`
	PostDir     string   `json:"post_dir,omitempty"`
	Type        string   `json:"type,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Draft       bool     `json:"draft,omitempty"`
	Status      string   `json:"status,omitempty"`
	Description string   `json:"description,omitempty"`
	Slug        string   `json:"slug,omitempty"` // Slug from metadata JSON, used as a fallback
	Date        string   `json:"date,omitempty"`
}

// Config holds application configuration
type Config struct {
	DBUrl      string
	SyncAll    bool
	TimeCutoff time.Time
	OutputDir  string
	DryRun     bool
	Verbose    bool
}

func main() {
	logger := log.New(os.Stdout, "[SYNC] ", log.LstdFlags|log.Lshortfile)

	config, err := parseConfig()
	if err != nil {
		logger.Fatalf("Configuration error: %v", err)
	}

	if config.Verbose {
		logger.Printf("Starting content sync with config: SyncAll=%t, DryRun=%t, OutputDir='%s'", config.SyncAll, config.DryRun, config.OutputDir)
		if !config.SyncAll && !config.TimeCutoff.IsZero() {
			logger.Printf("Time cutoff for sync: %s", config.TimeCutoff.Format(rfc3339Format))
		}
	}

	// Get existing slugs from the filesystem if not syncing all
	existingFileSlugs := make(map[string]bool)
	existingFileTitles := make(map[string]bool)
	if !config.SyncAll {
		var errSlugs error
		// Pass config.Verbose to control logging within getExistingFileSlugs
		existingFileSlugs, existingFileTitles, errSlugs = getExistingFileSlugs(config.OutputDir, logger, config.Verbose)
		if errSlugs != nil {
			logger.Printf("Warning: Could not read existing slugs from '%s': %v. Proceeding without slug-based check for missing files, relying on time-based sync only for delta.", config.OutputDir, errSlugs)
			// existingFileSlugs will be empty, so the NOT IN clause might not be added or might behave as if no files exist.
		} else if config.Verbose {
			logger.Printf("Found %d existing file slugs in '%s' for differential sync.", len(existingFileSlugs), config.OutputDir)
		}
	}

	db, err := connectDB(config.DBUrl)
	if err != nil {
		logger.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Pass existingFileSlugs to fetchPosts
	posts, err := fetchPosts(db, config, existingFileSlugs, existingFileTitles, logger)
	if err != nil {
		logger.Fatalf("Failed to fetch posts: %v", err)
	}

	if len(posts) == 0 {
		logger.Println("No posts to process based on current criteria.")
		// Update timestamp if this was a differential sync check and nothing was found
		if !config.DryRun && !config.SyncAll {
			if err := updateTimestamp(); err != nil {
				logger.Printf("Warning: Failed to update timestamp: %v", err)
			} else if config.Verbose {
				logger.Println("Updated build timestamp after checking for posts (none found).")
			}
		}
		return
	}

	logger.Printf("Found %d posts to process", len(posts))

	processedCount := 0
	errorCount := 0
	for _, post := range posts {
		if config.Verbose {
			logger.Printf("Processing post: ID=%d, Title='%s', DB Slug='%s'", post.ID, post.Title, post.Slug)
		}
		if err := processPost(post, config, logger); err != nil {
			logger.Printf("Error processing post ID %d ('%s'): %v", post.ID, post.Title, err)
			errorCount++
			continue
		}
		processedCount++
	}

	logger.Printf("Processing complete: %d processed, %d errors", processedCount, errorCount)

	// Update timestamp if not dry run and we successfully processed posts
	if !config.DryRun && processedCount > 0 {
		if err := updateTimestamp(); err != nil {
			logger.Printf("Warning: Failed to update timestamp: %v", err)
		} else if config.Verbose {
			logger.Println("Updated build timestamp after processing posts.")
		}
	}
}

// getEffectiveSlugFromFile attempts to parse frontmatter to find a 'slug' or 'title'.
// It returns the "effective slug" (frontmatter slug, or slugified frontmatter title)
// and a boolean indicating if such an identifier was successfully derived from frontmatter.
// It relies on the global 'slugify' function.
func getEffectiveSlugFromFile(filePath string, logger *log.Logger, verbose bool) (string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		if verbose {
			logger.Printf("Warning: Could not open file %s to read frontmatter: %v", filePath, err)
		}
		return "", false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inFrontmatter := false
	var frontmatterLines []string

	// Basic frontmatter detection (lines between "---")
	if scanner.Scan() && strings.TrimSpace(scanner.Text()) == "---" {
		// Check if the file is potentially empty or too small to have valid frontmatter
		// This simple check can be expanded if needed.
		if stat, _ := file.Stat(); stat != nil && stat.Size() < 8 { // "---" \n "---" \n
			if verbose {
				logger.Printf("File %s is too small for valid frontmatter, skipping frontmatter parse.", filePath)
			}
			return "", false
		}

		inFrontmatter = true
		for scanner.Scan() {
			line := scanner.Text()
			if strings.TrimSpace(line) == "---" {
				break // End of frontmatter
			}
			frontmatterLines = append(frontmatterLines, line)
		}
	}

	if !inFrontmatter || len(frontmatterLines) == 0 {
		if verbose && inFrontmatter { // inFrontmatter true but no lines means "--- \n ---"
			logger.Printf("File %s: Detected frontmatter delimiters but no content within. No effective slug from frontmatter.", filePath)
		}
		return "", false
	}

	var fmSlug, fmTitle string

	// Regex to find 'slug: value' or 'title: value' lines.
	// Captures unquoted, single-quoted, or double-quoted values.
	// Example: slug: my-value / slug: "my value" / slug: 'my value'
	// Limitations: Does not handle complex YAML structures, multiline strings, comments within values perfectly.
	slugRegex := regexp.MustCompile(`^\s*slug:\s*(?:"([^"]*)"|'([^']*)'|([^'"\s#][^#]*?))\s*(?:#.*)?$`)
	titleRegex := regexp.MustCompile(`^\s*title:\s*(?:"([^"]*)"|'([^']*)'|([^'"\s#][^#]*?))\s*(?:#.*)?$`)

	extractValue := func(match []string) string {
		if len(match) == 4 { // Entire match + 3 capturing groups
			if match[1] != "" {
				return strings.TrimSpace(match[1])
			} // double-quoted
			if match[2] != "" {
				return strings.TrimSpace(match[2])
			} // single-quoted
			if match[3] != "" {
				return strings.TrimSpace(match[3])
			} // unquoted
		}
		return ""
	}

	for _, line := range frontmatterLines {
		if fmSlug == "" {
			match := slugRegex.FindStringSubmatch(line)
			val := extractValue(match)
			if val != "" {
				fmSlug = val
			}
		}
		if fmTitle == "" {
			match := titleRegex.FindStringSubmatch(line)
			val := extractValue(match)
			if val != "" {
				fmTitle = val
			}
		}
	}

	if fmSlug != "" {
		// Use slugify to ensure the extracted slug is in the canonical format,
		// in case the frontmatter slug itself has characters that slugify would clean.
		// However, typically frontmatter slugs are already clean.
		// If you trust frontmatter slugs to be 100% as-is, you can just return fmSlug.
		// For consistency with slugify(title), let's slugify fmSlug too.
		// Or, decide if fmSlug should be used raw if present.
		// Let's assume fmSlug is canonical and use it directly if present.
		return fmSlug, true
	}
	if fmTitle != "" {
		s := slugify(fmTitle) // slugify is the function defined in your main code
		if s != "" {
			return s, true
		}
	}

	return "", false
}

func getEffectiveTitleFromFile(filePath string, logger *log.Logger, verbose bool) (string, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		if verbose {
			logger.Printf("Warning: Could not open file %s to read frontmatter: %v", filePath, err)
		}
		return "", false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inYAMLFrontmatter := false
	var frontmatterLines []string

	// Check if file starts with "---" (YAML frontmatter)
	if scanner.Scan() {
		firstLine := strings.TrimSpace(scanner.Text())
		if firstLine == "---" {
			inYAMLFrontmatter = true
			for scanner.Scan() {
				line := scanner.Text()
				if strings.TrimSpace(line) == "---" {
					break // End of frontmatter
				}
				frontmatterLines = append(frontmatterLines, line)
			}
		} else if strings.HasPrefix(firstLine, "{") {
			// Handle fallback JSON frontmatter
			jsonContent := firstLine
			for scanner.Scan() {
				line := scanner.Text()
				jsonContent += line
				if strings.Contains(line, "}") {
					break
				}
			}

			var frontJSON map[string]interface{}
			if err := json.Unmarshal([]byte(jsonContent), &frontJSON); err != nil {
				if verbose {
					logger.Printf("Warning: Invalid JSON frontmatter in %s: %v", filePath, err)
				}
				return "", false
			}

			if titleVal, ok := frontJSON["title"]; ok {
				if titleStr, ok := titleVal.(string); ok && titleStr != "" {
					return strings.TrimSpace(titleStr), true
				}
			}
			return "", false
		}
	}

	// No YAML frontmatter or JSON fallback
	if !inYAMLFrontmatter || len(frontmatterLines) == 0 {
		return "", false
	}

	// Parse YAML frontmatter lines
	titleRegex := regexp.MustCompile(`^\s*title:\s*(?:"([^"]*)"|'([^']*)'|([^'"\s#][^#]*?))\s*(?:#.*)?$`)
	extractValue := func(match []string) string {
		if len(match) == 4 {
			if match[1] != "" {
				return strings.TrimSpace(match[1])
			}
			if match[2] != "" {
				return strings.TrimSpace(match[2])
			}
			if match[3] != "" {
				return strings.TrimSpace(match[3])
			}
		}
		return ""
	}

	for _, line := range frontmatterLines {
		match := titleRegex.FindStringSubmatch(line)
		val := extractValue(match)
		if val != "" {
			return val, true
		}
	}

	return "", false
}

// getExistingFileSlugs scans the output directory for existing markdown files.
// It attempts to determine an "effective slug" for each file by checking its frontmatter
// for a 'slug' field, then a 'title' field (which is then slugified),
// and finally falls back to the filename if frontmatter doesn't yield a usable identifier.
func getExistingFileSlugs(outputDir string, logger *log.Logger, verbose bool) (map[string]bool, map[string]bool, error) {
	existingEffectiveSlugs := make(map[string]bool)
	existingTitles := make(map[string]bool)

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if verbose {
			logger.Printf("Output directory '%s' does not exist. Assuming no existing file slugs.", outputDir)
		}
		return existingEffectiveSlugs, existingTitles, nil
	}

	err := filepath.WalkDir(outputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			logger.Printf("Warning: Error accessing path '%s' during slug scan: %v", path, err)
			return nil // Continue walking other files/directories
		}

		if d.IsDir() {
			return nil // Skip directories
		}

		fileNameLower := strings.ToLower(d.Name())
		if !(strings.HasSuffix(fileNameLower, ".md") || strings.HasSuffix(fileNameLower, ".markdown")) {
			return nil // Skip non-markdown files
		}

		var finalIdentifierForFile string

		// Try to get an effective slug from the file's frontmatter (slug field or slugified title)
		effectiveFmSlug, foundInFrontmatter := getEffectiveSlugFromFile(path, logger, verbose)
		effectiveFmTitle, foundInFrontmatter := getEffectiveTitleFromFile(path, logger, verbose)

		existingTitles[effectiveFmTitle] = true

		if foundInFrontmatter && effectiveFmSlug != "" {
			finalIdentifierForFile = effectiveFmSlug
			if verbose {
				logger.Printf("File '%s': Using effective slug '%s' from its frontmatter.", d.Name(), finalIdentifierForFile)
			}
		} else {
			// Fallback: use filename-derived slug if no suitable frontmatter identifier was found
			baseName := d.Name()
			var filenameSlug string
			if strings.HasSuffix(fileNameLower, ".markdown") {
				filenameSlug = baseName[:len(baseName)-len(".markdown")]
			} else { // .md
				filenameSlug = baseName[:len(baseName)-len(".md")]
			}

			// Ensure the extracted filename slug is itself "clean" or "slugified".
			// If filenames are already expected to be valid slugs, this is fine.
			// If they can contain spaces or other characters, slugify it:
			// filenameSlug = slugify(filenameSlug)

			finalIdentifierForFile = filenameSlug // Assuming filename is already a slug
			if verbose {
				if foundInFrontmatter { // but effectiveFmSlug was empty
					logger.Printf("File '%s': Frontmatter slug/title resulted in empty slug. Using filename-derived slug '%s'.", d.Name(), finalIdentifierForFile)
				} else {
					logger.Printf("File '%s': No usable slug/title in frontmatter. Using filename-derived slug '%s'.", d.Name(), finalIdentifierForFile)
				}
			}
		}

		if finalIdentifierForFile != "" {
			existingEffectiveSlugs[finalIdentifierForFile] = true
		} else if verbose {
			logger.Printf("File '%s': Could not determine a valid identifier (frontmatter or filename). Skipping.", d.Name())
		}

		return nil
	})

	if err != nil { // This err is from filepath.WalkDir itself if the initial call failed.
		return nil, nil, fmt.Errorf("error walking directory %s: %w", outputDir, err)
	}
	return existingEffectiveSlugs, existingTitles, nil
}

func parseConfig() (*Config, error) {
	config := &Config{
		OutputDir: "posts", // Default output directory
		DryRun:    false,
		Verbose:   false,
	}

	dbName := os.Getenv("TURSO_DATABASE_NAME")
	dbToken := os.Getenv("TURSO_DATABASE_AUTH_TOKEN")

	if dbName == "" || dbToken == "" {
		return nil, fmt.Errorf("missing required environment variables: TURSO_DATABASE_NAME and TURSO_DATABASE_AUTH_TOKEN")
	}
	config.DBUrl = fmt.Sprintf("%s?authToken=%s", dbName, dbToken)

	args := os.Args[1:]
	customHoursSet := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "sync_all", "--sync-all", "-a":
			config.SyncAll = true
		case "sync_post", "--sync-posts", "-p": // Assuming this implies differential sync
			config.SyncAll = false
		case "--dry-run", "-d":
			config.DryRun = true
		case "--verbose", "-v":
			config.Verbose = true
		case "--output-dir", "-o":
			if i+1 < len(args) {
				config.OutputDir = args[i+1]
				i++ // consume next arg
			} else {
				return nil, fmt.Errorf("argument %s requires a value", arg)
			}
		case "--hours":
			if i+1 < len(args) {
				hours := parseHours(args[i+1])
				config.TimeCutoff = time.Now().Add(time.Duration(-hours) * time.Hour)
				customHoursSet = true
				i++ // consume next arg
			} else {
				return nil, fmt.Errorf("argument %s requires a value", arg)
			}
		}
	}

	if !config.SyncAll && !customHoursSet { // Only use last build time if not sync_all and --hours not set
		config.TimeCutoff = getLastBuildTime()
	}
	// If config.SyncAll is true, TimeCutoff is not used by fetchPosts logic for query conditions.
	// If customHoursSet is true, it overrides getLastBuildTime().

	return config, nil
}

func parseHours(s string) int {
	var hours int
	_, err := fmt.Sscanf(s, "%d", &hours)
	if err != nil || hours <= 0 {
		log.Printf("Warning: Invalid hours value '%s', defaulting to 1 hour.", s)
		return 1 // Default to 1 hour if parsing fails or value is non-positive
	}
	return hours
}

func getLastBuildTime() time.Time {
	data, err := os.ReadFile(timestampFile)
	if err == nil {
		t, err := time.Parse(rfc3339Format, strings.TrimSpace(string(data)))
		if err == nil {
			return t
		}
		log.Printf("Warning: Could not parse timestamp from %s ('%s'): %v. Defaulting.", timestampFile, strings.TrimSpace(string(data)), err)
	} else if !os.IsNotExist(err) {
		log.Printf("Warning: Could not read timestamp file %s: %v. Defaulting.", timestampFile, err)
	}
	// Default to 1 hour ago if file doesn't exist or is invalid
	return time.Now().Add(-1 * time.Hour)
}

func updateTimestamp() error {
	timestamp := time.Now().UTC().Format(rfc3339Format)
	return os.WriteFile(timestampFile, []byte(timestamp), 0644)
}

func connectDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

// fetchPosts retrieves posts from database based on configuration and existing file slugs
func fetchPosts(db *sql.DB, config *Config, existingFileSlugs, existingFileTitles map[string]bool, logger *log.Logger) ([]Post, error) {
	var queryBuilder strings.Builder
	var args []interface{}

	queryBuilder.WriteString("SELECT id, title, slug, body, metadata, created_at, updated_at, author_id, deleted FROM posts WHERE deleted = false")

	if config.SyncAll {
		queryBuilder.WriteString(" ORDER BY updated_at DESC")
	} else {
		// Differential sync: (time-based) OR (slug-based for missing files)
		queryBuilder.WriteString(" AND (") // Start of compound OR condition

		// Part 1: Time-based condition
		queryBuilder.WriteString(" (created_at > ? OR updated_at > ?)")
		timeStr := config.TimeCutoff.Format(timeFormat)
		args = append(args, timeStr, timeStr)

		// Part 2: Slug-based condition
		if len(existingFileSlugs) > 0 {
			queryBuilder.WriteString(" OR (slug IS NOT NULL AND slug != '' AND slug NOT IN (")
			placeholders := make([]string, 0, len(existingFileSlugs))
			for slug := range existingFileSlugs {
				placeholders = append(placeholders, "?")
				args = append(args, slug)
			}
			queryBuilder.WriteString(strings.Join(placeholders, ","))
			queryBuilder.WriteString("))")
		} else {
			queryBuilder.WriteString(" OR (slug IS NOT NULL AND slug != '')")
		}

		// Part 3: Title-based condition
		titlePlaceholders := make([]string, 0, len(existingFileTitles))
		for title := range existingFileTitles {
			if title != "" {
				titlePlaceholders = append(titlePlaceholders, "?")
				args = append(args, title)
			}
		}
		if len(titlePlaceholders) > 0 {
			queryBuilder.WriteString(" OR (title IS NOT NULL AND title != '' AND title NOT IN (")
			queryBuilder.WriteString(strings.Join(titlePlaceholders, ","))
			queryBuilder.WriteString("))")
		}

		queryBuilder.WriteString(") ORDER BY updated_at DESC")
	}

	finalQuery := queryBuilder.String()
	if config.Verbose {
		logger.Printf("Executing query: %s", finalQuery)
		logger.Printf("Query args (%d): %v", len(args), args)
	}

	rows, err := db.Query(finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("query execution failed ('%s'): %w", finalQuery, err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var dbSlug sql.NullString
		var metadata sql.NullString

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&dbSlug,
			&post.Body,
			&metadata,
			&post.Created,
			&post.Updated,
			&post.AuthorID,
			&post.Deleted,
		)
		if err != nil {
			logger.Printf("Error scanning row: %v", err)
			continue
		}

		post.Slug = ""
		if dbSlug.Valid {
			post.Slug = dbSlug.String
		}

		post.Metadata = "{}"
		if metadata.Valid {
			post.Metadata = metadata.String
		}

		if post.Deleted || strings.TrimSpace(post.Body) == "" || strings.TrimSpace(post.Title) == "" {
			if config.Verbose {
				logger.Printf("Skipping post ID %d ('%s'): empty title/body or marked deleted.", post.ID, post.Title)
			}
			continue
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	var newPosts []Post
	for _, post := range posts {
		if existingFileTitles[post.Title] {
			if config.Verbose {
				logger.Printf("Skipping post ID %d ('%s'): title already exists in '%s'.", post.ID, post.Title, config.OutputDir)
			}
			continue
		}
		if existingFileSlugs[post.Slug] {
			if config.Verbose {
				logger.Printf("Skipping post ID %d ('%s'): slug already exists in '%s'.", post.ID, post.Slug, config.OutputDir)
			}
			continue
		}
		newPosts = append(newPosts, post)
	}
	return newPosts, nil
}

func processPost(post Post, config *Config, logger *log.Logger) error {
	metadata, err := parseMetadata(post.Metadata) // post.Metadata is already handled for NULL by fetchPosts
	if err != nil {
		// parseMetadata now logs internally and returns empty metadata on JSON error, so this might not be fatal
		logger.Printf("Warning: Non-fatal issue parsing metadata for post ID %d: %v. Proceeding with defaults.", post.ID, err)
	}

	bodyMd := post.Body
	//bodyMd, err := htmltomarkdown.ConvertString(post.Body)
	//if err != nil {
	//	return fmt.Errorf("failed to convert HTML to Markdown for post ID %d: %w", post.ID, err)
	//}

	// Slug generation/determination logic
	finalSlug := post.Slug // Slug from DB takes precedence if not empty
	if finalSlug == "" && metadata.Slug != "" {
		finalSlug = metadata.Slug // Fallback to slug from metadata JSON
		if config.Verbose {
			logger.Printf("Post ID %d: Using slug '%s' from metadata JSON.", post.ID, finalSlug)
		}
	}
	if finalSlug == "" { // If still no slug, generate from title
		originalTitle := post.Title
		finalSlug = slugify(originalTitle)
		if config.Verbose {
			logger.Printf("Post ID %d: Generated slug '%s' from title '%s'.", post.ID, finalSlug, originalTitle)
		}
	}
	if finalSlug == "" {
		return fmt.Errorf("could not determine or generate a slug for post ID %d, title '%s'", post.ID, post.Title)
	}

	filePath, err := determineFilePath(metadata, finalSlug, config.OutputDir)
	if err != nil {
		return fmt.Errorf("failed to determine file path for post ID %d (slug '%s'): %w", post.ID, finalSlug, err)
	}

	content, err := createPostContent(post, metadata, finalSlug, bodyMd, config.Verbose, logger)
	if err != nil {
		return fmt.Errorf("failed to create post content for post ID %d (slug '%s'): %w", post.ID, finalSlug, err)
	}

	if config.DryRun {
		logger.Printf("DRY RUN: Would write to %s (%d bytes)", filePath, len(content))
		// Optionally print content in dry run for debugging
		// if config.Verbose { fmt.Println(content) }
		return nil
	}

	if err := writeFile(filePath, content); err != nil {
		return fmt.Errorf("failed to write file %s for post ID %d: %w", filePath, post.ID, err)
	}

	if config.Verbose {
		logger.Printf("Successfully written post ID %d to: %s", post.ID, filePath)
	}
	return nil
}

func parseMetadata(metadataStr string) (*PostMetadata, error) {
	metadata := &PostMetadata{} // Initialize with defaults
	if strings.TrimSpace(metadataStr) == "" {
		return metadata, nil // No metadata string, return default empty struct
	}

	if err := json.Unmarshal([]byte(metadataStr), metadata); err != nil {
		// Log as a warning and return an empty PostMetadata struct, allowing processing to continue.
		log.Printf("Warning: Failed to parse metadata JSON ('%s'). Using default/empty metadata. Error: %v", metadataStr, err)
		return &PostMetadata{}, nil // Return empty, non-nil metadata
	}
	return metadata, nil
}

// Corrected determineFilePath
func determineFilePath(metadata *PostMetadata, slug, outputDir string) (string, error) {
	var resolvedBaseDir string // This will be the full path to the directory where the file should be saved.

	if metadata.PostDir != "" {
		// metadata.PostDir is treated as a path relative to outputDir (e.g., "articles/technical")
		resolvedBaseDir = filepath.Join(outputDir, metadata.PostDir)
	} else {
		// Default directory logic: use metadata.Type or "posts" as a subdirectory within outputDir.
		subDir := "posts" // Default subdirectory name
		if metadata.Type != "" {
			// Sanitize or validate metadata.Type if it can contain problematic characters for a directory name.
			// For now, assume it's a simple, clean string like "article", "note", "project".
			typeNameDir := strings.ToLower(strings.TrimSpace(metadata.Type))
			// Ensure typeNameDir is not empty and is a valid path segment.
			// A simple slugify might be good here if type can be arbitrary.
			// For now, just use it if non-empty.
			if typeNameDir != "" {
				// Basic check: disallow path separators in typeNameDir to prevent breaking out.
				// More robust sanitization might be needed depending on source of metadata.Type.
				if !strings.Contains(typeNameDir, string(filepath.Separator)) && typeNameDir != "." && typeNameDir != ".." {
					subDir = typeNameDir
				} else {
					log.Printf("Warning: Invalid characters or path segments in metadata.Type ('%s'), using default subdir '%s'.", metadata.Type, subDir)
				}
			}
		}
		resolvedBaseDir = filepath.Join(outputDir, subDir)
	}

	if slug == "" {
		return "", fmt.Errorf("cannot determine file path with an empty slug")
	}
	// Ensure slug itself is clean for a filename (slugify function should handle this).
	// For example, slugify should prevent slugs like "../evil-slug"
	// but an extra check here for path traversal in the slug itself might be good if slugify isn't perfectly secure.
	// However, slugify typically produces safe results.

	filename := fmt.Sprintf("%s.md", slug)
	return filepath.Join(resolvedBaseDir, filename), nil
}

func createPostContent(post Post, metadata *PostMetadata, slug, bodyMd string, verbose bool, logger *log.Logger) (string, error) {
	createdTime, err := parseTime(post.Created)
	if err != nil {
		if verbose {
			logger.Printf("Warning: Invalid created_at time ('%s') for post ID %d (slug '%s'). Using current time. Error: %v", post.Created, post.ID, slug, err)
		}
		createdTime = time.Now() // Fallback to current time
	}

	updatedTime, err := parseTime(post.Updated)
	if err != nil {
		if verbose {
			logger.Printf("Warning: Invalid updated_at time ('%s') for post ID %d (slug '%s'). Using created time. Error: %v", post.Updated, post.ID, slug, err)
		}
		updatedTime = createdTime // Fallback to createdTime if updatedTime is invalid
	}

	// Ensure updatedTime is not before createdTime
	if updatedTime.Before(createdTime) {
		if verbose {
			logger.Printf("Warning: updated_at ('%s') is before created_at ('%s') for post ID %d (slug '%s'). Setting updated_at to created_at.", updatedTime.Format(rfc3339Format), createdTime.Format(rfc3339Format), post.ID, slug)
		}
		updatedTime = createdTime
	}
	if metadata.Title == "" {
		metadata.Title = post.Title
	}
	if !metadata.Draft || metadata.Status == "published" {
		metadata.Status = "published"
		metadata.Draft = false
	} else {
		metadata.Status = "draft"
		metadata.Draft = true
	}
	if metadata.Date == "" {
		metadata.Date = createdTime.Format(rfc3339Format)
	}
	if metadata.Slug == "" {
		metadata.Slug = slug
	}
	frontmatterStr, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		frontmatter := map[string]interface{}{
			"title":   post.Title,
			"slug":    slug, // Use the final, determined slug
			"date":    createdTime.Format(rfc3339Format),
			"lastmod": updatedTime.Format(rfc3339Format),
			// "id":      post.ID, // Optionally include the original DB ID
		}

		if metadata.Type != "" {
			frontmatter["type"] = metadata.Type
		}
		if len(metadata.Tags) > 0 {
			frontmatter["tags"] = metadata.Tags
		}
		if len(metadata.Categories) > 0 {
			frontmatter["categories"] = metadata.Categories
		}
		if metadata.Description != "" {
			frontmatter["description"] = metadata.Description
		}
		if metadata.Draft { // Only add 'draft: true' if it's explicitly true
			frontmatter["draft"] = true
			frontmatter["published"] = false
		}
		// Add any other custom fields from metadata if needed, by iterating through a map representation of metadata.

		var frontmatterBuilder strings.Builder
		frontmatterBuilder.WriteString("---\n")
		// Marshal to YAML-like string. For more complex cases, a YAML library would be better.
		for key, value := range frontmatter {
			switch v := value.(type) {
			case string:
				// Basic YAML string quoting: escape double quotes within the string.
				escapedValue := strings.ReplaceAll(v, "\"", "\\\"")
				frontmatterBuilder.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, escapedValue))
			case bool:
				frontmatterBuilder.WriteString(fmt.Sprintf("%s: %t\n", key, v))
			case []string:
				if len(v) > 0 {
					frontmatterBuilder.WriteString(fmt.Sprintf("%s:\n", key))
					for _, item := range v {
						escapedItem := strings.ReplaceAll(item, "\"", "\\\"")
						frontmatterBuilder.WriteString(fmt.Sprintf("  - \"%s\"\n", escapedItem))
					}
				}
			// Add other types as needed (int, float, etc.)
			default: // Fallback for other types (e.g., int for ID)
				frontmatterBuilder.WriteString(fmt.Sprintf("%s: %v\n", key, v))
			}
		}
		frontmatterBuilder.WriteString("---\n\n")

		return frontmatterBuilder.String() + bodyMd, nil
	}
	return string(frontmatterStr) + "\n\n" + bodyMd, nil
}

func parseTime(timeStr string) (time.Time, error) {
	timeStr = strings.TrimSpace(timeStr)
	if timeStr == "" {
		return time.Time{}, fmt.Errorf("empty time string")
	}

	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05Z07:00", // ISO8601 with timezone
		"2006-01-02 15:04:05Z07",    // ISO8601 with short timezone
		"2006-01-02 15:04:05",       // Common DB format (like SQLite DATETIME)
		"2006-01-02T15:04:05",       // ISO8601 without timezone
		"2006-01-02 15:04",          // Without seconds
		"2006-01-02T15:04",          // ISO8601 without seconds
		"2006-01-02",                // Date only
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}
	// Attempt to parse with location if it's a common local time format from DB
	// For instance, if DB stores in UTC but doesn't specify 'Z'
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.UTC); err == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("could not parse time string '%s' with known formats", timeStr)
}

func writeFile(filePath, content string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

// slugify creates a URL-friendly slug from a string
func slugify(input string) string {
	if input == "" {
		return ""
	}

	// Normalize to lowercase
	s := strings.ToLower(input)

	// Replace common separators and symbols with hyphens or remove
	s = strings.ReplaceAll(s, "&", "and")
	s = strings.ReplaceAll(s, "@", "at")
	s = strings.ReplaceAll(s, "_", "-") // replace underscores with hyphens

	// Use a regex to remove unwanted characters, keep letters, numbers, and hyphens
	// This regex allows a broader range of characters initially to preserve language specific chars if desired before further transliteration
	// For basic ASCII slugging:
	reg := regexp.MustCompile(`[^a-z0-9\-]+`)
	s = reg.ReplaceAllString(s, "-")

	// Replace multiple hyphens with a single hyphen
	reg = regexp.MustCompile(`-{2,}`)
	s = reg.ReplaceAllString(s, "-")

	// Trim leading and trailing hyphens
	s = strings.Trim(s, "-")

	// Optional: Truncate slug to a max length
	maxLength := 75 // Common slug length limit
	if len(s) > maxLength {
		s = s[:maxLength]
		// Ensure not to cut in the middle of a multi-byte character if dealing with unicode
		// and try to cut at a word boundary (hyphen)
		if lastHyphen := strings.LastIndex(s, "-"); lastHyphen > maxLength/2 && lastHyphen < len(s)-1 { // Avoid tiny stubs
			s = s[:lastHyphen]
		}
		s = strings.Trim(s, "-") // Trim again if truncation resulted in trailing hyphen
	}

	return s
}

package plugins

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"regexp"
	"strings"

	models "github.com/mr-destructive/mr-destructive.github.io/models"
)

// GenerateTOC extracts headings from a Goldmark AST and returns a nested TOC structure.
func GenerateTOC(node ast.Node, source []byte) []*models.TOCItem {
	var tocItems []*models.TOCItem
	var stack []*models.TOCItem // Stack to keep track of parent items

	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() == ast.KindHeading {
			heading := n.(*ast.Heading)
			if entering {
				// Extract text content of the heading
				var buf bytes.Buffer
				for c := heading.FirstChild(); c != nil; c = c.NextSibling() {
					if c.Kind() == ast.KindText {
						buf.Write(c.(*ast.Text).Segment.Value(source))
					} else if c.Kind() == ast.KindLink {
						// Handle links within headings, extract text
						link := c.(*ast.Link)
						for lc := link.FirstChild(); lc != nil; lc = lc.NextSibling() {
							if lc.Kind() == ast.KindText {
								buf.Write(lc.(*ast.Text).Segment.Value(source))
							}
						}
					}
				}
				text := buf.String()
				id := generateID(text)

				// Ensure the heading has an ID attribute
				if heading.Attributes() == nil {
					heading.SetAttributeString("id", []byte(id))
				} else {
					foundID := false
					for _, attr := range heading.Attributes() {
						if string(attr.Name) == "id" {
							id = string(attr.Value.([]byte))
							foundID = true
							break
						}
					}
					if !foundID {
						heading.SetAttributeString("id", []byte(id))
					}
				}

				item := &models.TOCItem{
					Text:  text,
					ID:    id,
					Level: heading.Level,
				}

				for len(stack) > 0 && stack[len(stack)-1].Level >= item.Level {
					stack = stack[:len(stack)-1]
				}

				if len(stack) > 0 {
				} else {
					tocItems = append(tocItems, item)
				}
				stack = append(stack, item)
			}
		}
		return ast.WalkContinue, nil
	})

	return tocItems
}

var idCleaner = regexp.MustCompile(`[^a-zA-Z0-9-_]+`)

func generateID(text string) string {
	id := strings.ToLower(text)
	id = strings.ReplaceAll(id, " ", "-")
	id = idCleaner.ReplaceAllString(id, "")
	return id
}


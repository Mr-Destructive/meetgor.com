package plugins

import (
	"fmt"
	"reflect"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"github.com/yuin/goldmark"
)

// goPlaygroundRenderer is a custom renderer for fenced code blocks.
type goPlaygroundRenderer struct {
	html.Config
}

// newGoPlaygroundRenderer creates a new instance of the custom renderer.
func newGoPlaygroundRenderer() renderer.NodeRenderer {
	return &goPlaygroundRenderer{
		Config: html.NewConfig(),
	}
}

// RegisterFuncs registers the rendering function for the FencedCodeBlock node type.
func (r *goPlaygroundRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

// renderFencedCodeBlock is the function that renders the fenced code block.
func (r *goPlaygroundRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	if entering {
		language := string(n.Language(source))
		if language == "go" || language == "golang" {
			// Opening tag with a wrapper div for styling and functionality
			_, _ = w.WriteString("<div class=\"go-playground\">")

			// Add the run, edit, and reset buttons
			_, _ = w.WriteString("<div class=\"controls\">")
			_, _ = w.WriteString("<button class=\"run-button\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"feather feather-play\"><polygon points=\"5 3 19 12 5 21 5 3\"></polygon></svg> Run</button>")
			_, _ = w.WriteString("<button class=\"edit-button\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"feather feather-edit\"><path d=\"M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7\"></path><path d=\"M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z\"></path></svg> Edit</button>")
			_, _ = w.WriteString("<button class=\"reset-button\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"feather feather-refresh-cw\"><polyline points=\"23 4 23 10 17 10\"></polyline><polyline points=\"1 20 1 14 7 14\"></polyline><path d=\"M20.49 9A9 9 0 0 0 5.64 5.64M3.51 15A9 9 0 0 0 18.36 18.36\"></path></svg> Reset</button>")
			_, _ = w.WriteString("<div class=\"info\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><circle cx=\"12\" cy=\"12\" r=\"10\"></circle><line x1=\"12\" y1=\"16\" x2=\"12\" y2=\"12\"></line><line x1=\"12\" y1=\"8\" x2=\"12.01\" y2=\"8\"></line></svg> Go 1.23</div>")
			_, _ = w.WriteString("</div>")

			// Default <pre> and <code> tags
			_, _ = w.WriteString("<pre><code")
			if n.Info != nil {
				_, _ = w.WriteString(fmt.Sprintf(" class=\"language-%s\"", language))
			}
			_, _ = w.WriteString(">")

			// Write the code content
			lines := n.Lines()
			for i := 0; i < lines.Len(); i++ {
				line := lines.At(i)
				_, _ = w.Write(line.Value(source))
			}
		} else {
			// Default rendering for other languages
			_, _ = w.WriteString("<div class=\"code-block\">")
			_, _ = w.WriteString("<div class=\"code-header\"><span class=\"language-name\">" + language + "</span><button class=\"copy-button\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"feather feather-copy\"><rect x=\"9\" y=\"9\" width=\"13\" height=\"13\" rx=\"2\" ry=\"2\"></rect><path d=\"M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1\"></path></svg></button></div>")
			_, _ = w.WriteString("<pre><code")
			if n.Info != nil {
				_, _ = w.WriteString(fmt.Sprintf(" class=\"language-%s\"", language))
			}
			_, _ = w.WriteString(">")
			lines := n.Lines()
			for i := 0; i < lines.Len(); i++ {
				line := lines.At(i)
				_, _ = w.Write(line.Value(source))
			}
		}
	} else {
		language := string(n.Language(source))
		if language == "go" || language == "golang" {
			// Closing tags
			_, _ = w.WriteString("</code></pre>")
			_, _ = w.WriteString("<div class=\"output\"></div>")
			_, _ = w.WriteString("</div>")
		} else {
			_, _ = w.WriteString("</code></pre></div>")
		}
	}
	return ast.WalkContinue, nil
}

// GoPlaygroundExtender is the extender that adds the custom renderer.
type GoPlaygroundExtender struct{}

// Extend adds the custom renderer to the Goldmark instance.
func (e *GoPlaygroundExtender) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(newGoPlaygroundRenderer(), 200),
		),
	)
}

func init() {
	RegisterPlugin("GoPlayground", reflect.TypeOf(GoPlaygroundExtender{}))
}
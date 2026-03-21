package plugins

import "github.com/mr-destructive/meetgor.com/models"

type Phase string

const (
	PhaseRead        Phase = "read"
	PhaseTransform   Phase = "transform"
	PhaseRender      Phase = "render"
	PhasePostProcess Phase = "postprocess"
	PhaseServe       Phase = "serve"
	PhaseOther       Phase = "other"
)

type AdminPolicy int

const (
	AdminRun AdminPolicy = iota
	AdminSkip
	AdminOnly
)

type Plugin interface {
	Name() string
	Execute(ssg *models.SSG) error
}

type PluginMetadata interface {
	Phase() Phase
	Requires() []string
	AdminPolicy() AdminPolicy
}

var pluginRegistry = make(map[string]func() Plugin)

func RegisterPlugin(name string, factory func() Plugin) {
	pluginRegistry[name] = factory
}

func NewPlugin(name string) (Plugin, bool) {
	factory, exists := pluginRegistry[name]
	if !exists {
		return nil, false
	}
	return factory(), true
}

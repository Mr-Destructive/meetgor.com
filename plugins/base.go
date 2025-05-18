package plugins

import (
	//"reflect"

	"github.com/mr-destructive/mr-destructive.github.io/models"
)

type BasePlugin struct {
	PluginName string
}

func (bp *BasePlugin) Name() string {
	return bp.PluginName
}

func (bp *BasePlugin) Execute(ssg *models.SSG) {
	// write the plugin here
}

// uncomment this
//func init() {
//	RegisterPlugin("Base", reflect.TypeOf(BasePlugin{
//		PluginName: "Base",
//	}))
//}

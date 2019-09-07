package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"

// Node - wraps LilvNode
type Node struct {
	node    *C.LilvNode
	managed bool
}

// Nodes - wraps LilvNodes
type Nodes struct {
	nodes   *C.LilvNodes
	managed bool
}

// World - wraps LilvWorld
type World struct {
	world *C.LilvWorld
}

// Plugin - wraps LilvPlugin
type Plugin struct {
	plugin *C.LilvPlugin
}

// Plugins - wraps LilvPlugins
type Plugins struct {
	plugins *C.LilvPlugins
}

// PluginClass - wraps LilvPluginClass
type PluginClass struct {
	pluginClass *C.LilvPluginClass
}

// PluginClasses - wraps LilvPluginClasses
type PluginClasses struct {
	pluginClasses *C.LilvPluginClasses
	shared        bool
}

// Port - wraps LilvPort
type Port struct {
	port *C.LilvPort
}

// UI - wraps LilvUI
type UI struct {
	ui *C.LilvUI
}

// UIs - wraps LilvUIs
type UIs struct {
	uis    *C.LilvUIs
	shared bool
}

// Iter - alias LilvIter
type Iter C.LilvIter

// Instance - wraps LilvInstance
type Instance struct {
	instance *C.LilvInstance
}

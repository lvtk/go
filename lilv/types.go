package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// Plugin - wraps LilvPlugin
type Plugin struct {
	plugin *C.LilvPlugin
}

// PluginClass - wraps LilvPluginClass
type PluginClass struct {
	pluginClass *C.LilvPluginClass
}

// Port - wraps LilvPort
type Port struct {
	plugin *C.LilvPlugin
	port   *C.LilvPort
}

// ScalePoint - wraps LilvScalePoint
type ScalePoint struct {
	scalePoint *C.LilvScalePoint
}

// UI - wraps LilvUI
type UI struct {
	ui *C.LilvUI
}

// Node - wraps LilvNode
type Node struct {
	node    *C.LilvNode
	managed bool
}

// World - wraps LilvWorld
type World struct {
	world *C.LilvWorld
}

// Instance - wraps LilvInstance
type Instance struct {
	instance *C.LilvInstance
}

// State - wraps LilvState
type State struct {
	instance *C.LilvState
}

// Iter - alias LilvIter
type Iter C.LilvIter

// PluginClasses - wraps LilvPluginClasses
type PluginClasses struct {
	pluginClasses unsafe.Pointer
	managed       bool
}

// Plugins - wraps LilvPlugins
type Plugins struct {
	plugins unsafe.Pointer
}

// ScalePoints - wraps LilvScalePoints
type ScalePoints struct {
	scalePoints unsafe.Pointer
	managed     bool
}

// UIs - wraps LilvUIs
type UIs struct {
	uis     unsafe.Pointer
	managed bool
}

// Nodes - wraps LilvNodes
type Nodes struct {
	nodes   unsafe.Pointer
	managed bool
}

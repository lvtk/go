package ui

/*
#cgo pkg-config: lv2
#include "goui.h"
#include <stdlib.h>
#include <memory.h>
*/
import "C"
import (
	"unsafe"

	"github.com/lvtk/go/lv2"
)

// NativeWidget opaque widget type (alias of LV2UI_Widget)
type NativeWidget C.LV2UI_Widget

// GoUI alias to C-type GoUI
type GoUI C.GoUI

// Instance interface
type Instance interface {
	PortEvent()
	Cleanup()
	GetWidget() NativeWidget
}

// IdleInterface idle interface support
type IdleInterface interface {
	Idle() int
}

// URI of your UI
var URI string

// Instantiate - set this to create your Instance type
var Instantiate = func(features *lv2.FeatureList) Instance {
	return nil
}

// ExtensionData - set this to get extension data
var ExtensionData = func(uri string) unsafe.Pointer {
	return nil
}

var instances = make(map[*GoUI]Instance)

//export lvtk_go_ui_instantiate
func lvtk_go_ui_instantiate(ui *GoUI, p, b, w, c, f unsafe.Pointer) C.int {
	if Instantiate == nil {
		return 1
	}

	instance := Instantiate(new(lv2.FeatureList))
	if instance == nil {
		return 2
	}

	instances[ui] = instance
	return 0
}

//export lvtk_go_ui_port_event
func lvtk_go_ui_port_event(ui *GoUI, i, s, p, v unsafe.Pointer) {
	instances[ui].PortEvent()
}

//export lvtk_go_ui_cleanup
func lvtk_go_ui_cleanup(ui *GoUI) {
	i := instances[ui]
	if i != nil {
		i.Cleanup()
	}

	delete(instances, ui)
}

//export lvtk_go_ui_get_widget
func lvtk_go_ui_get_widget(ui *GoUI) NativeWidget {
	return instances[ui].GetWidget()
}

//export lvtk_go_ui_get_uri
func lvtk_go_ui_get_uri() *C.char {
	return C.CString(URI)
}

//export lvtk_go_ui_extension_data
func lvtk_go_ui_extension_data(u unsafe.Pointer) unsafe.Pointer {
	if ExtensionData == nil {
		return nil
	}
	uri := C.GoString((*C.char)(u))
	return ExtensionData(uri)
}

//export lvtk_go_ui_idle
func lvtk_go_ui_idle(ui *GoUI) C.int {
	i := instances[ui]
	if idle, ok := i.(IdleInterface); ok {
		return C.int(idle.Idle())
	}
	return 1
}

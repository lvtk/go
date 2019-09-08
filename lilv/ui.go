package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
*/
import "C"
import "unsafe"

/* UI */

/*
GetURI - Get the URI of a Plugin UI.
@param ui The Plugin UI
@return a shared value which must not be modified or freed.
*/
func (ui *UI) GetURI() *Node {
	if ui == nil || ui.ui == nil {
		return nil
	}
	return createSharedNode(C.lilv_ui_get_uri(ui.ui))
}

/*
GetClasses - Get the types (URIs of RDF classes) of a Plugin UI.
Returns a shared value which must not be modified.
*/
func (ui *UI) GetClasses() *Nodes {
	if ui == nil || ui.ui == nil {
		return nil
	}
	return createNodes(false, C.lilv_ui_get_classes(ui.ui))
}

/*
IsA - Check whether a plugin UI has a given type.
*/
func (ui *UI) IsA(classURI *Node) bool {
	if ui == nil || ui.ui == nil || classURI == nil || classURI.node == nil {
		return false
	}
	return bool(C.lilv_ui_is_a(ui.ui, classURI.node))
}

/*
GetBundleURI - Get the URI for a Plugin UI's bundle.
@param ui The Plugin UI
@return a shared value which must not be modified or freed.
*/
func (ui *UI) GetBundleURI() *Node {
	if ui == nil || ui.ui == nil {
		return nil
	}
	return createSharedNode(C.lilv_ui_get_bundle_uri(ui.ui))
}

/*
GetBinaryURI - Get the URI for a Plugin UI's shared library.
@param ui The Plugin UI
@return a shared value which must not be modified or freed.
*/
func (ui *UI) GetBinaryURI() *Node {
	if ui == nil || ui.ui == nil {
		return nil
	}
	return createSharedNode(C.lilv_ui_get_binary_uri(ui.ui))
}

/* UIs */

// Free - free this UIs collection
func (uis *UIs) Free() {
	if uis == nil || uis.uis == nil {
		return
	}
	if uis.managed {
		C.lilv_uis_free(uis.uis)
	}
	uis = nil
}

// Size get the number of uis in the collection
func (uis *UIs) Size() uint32 {
	if uis == nil || uis.uis == nil {
		return 0
	}
	return uint32(C.lilv_uis_size(uis.uis))
}

// Begin the start iter
func (uis *UIs) Begin() *Iter {
	if uis == nil || uis.uis == nil {
		return nil
	}
	return (*Iter)(C.lilv_uis_begin(uis.uis))
}

// Get the ui for this iter
func (uis *UIs) Get(iter *Iter) *UI {
	if uis == nil || uis.uis == nil {
		return nil
	}
	ui := new(UI)
	ui.ui = C.lilv_uis_get(uis.uis, unsafe.Pointer(iter))
	return ui
}

// Next iterator
func (uis *UIs) Next(iter *Iter) *Iter {
	if uis == nil || uis.uis == nil {
		return nil
	}
	return (*Iter)(C.lilv_uis_next(uis.uis,
		unsafe.Pointer(iter)))
}

// IsEnd true if end iter
func (uis *UIs) IsEnd(iter *Iter) bool {
	if uis == nil || uis.uis == nil {
		return false
	}
	return bool(C.lilv_uis_is_end(uis.uis, unsafe.Pointer(iter)))
}

// GetByURI find UI by uri
func (uis *UIs) GetByURI(uri *Node) *UI {
	if uis == nil || uis.uis == nil || uri == nil || uri.node == nil {
		return nil
	}
	ui := new(UI)
	ui.ui = C.lilv_uis_get_by_uri(uis.uis, uri.node)
	return ui
}

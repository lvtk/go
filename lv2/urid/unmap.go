package urid

/*
#cgo pkg-config: lv2
#include <lv2/lv2plug.in/ns/ext/urid/urid.h>

static const char* call_unmap(LV2_URID_Unmap* unmap, uint32_t urid) {
	return unmap->unmap(unmap->handle, urid);
}
*/
import "C"
import (
	"unsafe"
)

// UnmapRef pointer to an LV2_URID_Unmap
type UnmapRef unsafe.Pointer

// Unmap go version of LV2_URID_Unmap
type Unmap struct {
	unmap *C.LV2_URID_Unmap
}

// NewUnmapRef wraps a LV2_URID_Unmap pointer
func NewUnmapRef(ref UnmapRef) *Unmap {
	m := new(Unmap)
	m.unmap = (*C.LV2_URID_Unmap)(ref)
	return m
}

// Ref returns C pointer to LV2_URID_Unmap
func (x *Unmap) Ref() UnmapRef {
	return UnmapRef(x.unmap)
}

// Unmap a URI
func (x *Unmap) Unmap(urid uint32) string {
	if x == nil || x.unmap == nil {
		return ""
	}
	return C.GoString(C.call_unmap(x.unmap, C.uint32_t(urid)))
}

package urid

/*
#cgo pkg-config: lv2
#include <lv2/lv2plug.in/ns/ext/urid/urid.h>
#include <stdlib.h>

static uint32_t call_map(LV2_URID_Map* map, const char* uri) {
	return map->map(map->handle, uri);
}
*/
import "C"
import (
	"unsafe"
)

// MapRef pointer to an LV2_URID_Map
type MapRef unsafe.Pointer

// Map go version of LV2_URID_Map
type Map struct {
	cmap *C.LV2_URID_Map
}

// NewMapRef wraps a LV2_URID_Map pointer
func NewMapRef(ref MapRef) *Map {
	m := new(Map)
	m.cmap = (*C.LV2_URID_Map)(ref)
	return m
}

// Ref returns C pointer to LV2_URID_Map
func (x *Map) Ref() MapRef {
	return MapRef(x.cmap)
}

// Map a URI
func (x *Map) Map(uri string) uint32 {
	if x == nil || x.cmap == nil {
		return 0
	}
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	return uint32(C.call_map(x.cmap, cstr))
}

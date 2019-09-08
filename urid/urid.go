package urid

/*
#cgo pkg-config: lv2
#include <stdlib.h>
#include <lv2/lv2plug.in/ns/ext/urid/urid.h>
#include "directory.h"
*/
import "C"
import (
	"unsafe"

	"github.com/lvtk/go/lv2"
)

// Directory of mapped URIs
type Directory struct {
	directory unsafe.Pointer
}

func (d *Directory) ensureDirectory() {
	if d.directory == nil {
		d.directory = C.lvtk_uri_directory_new()
	}
}

// Free allocated data
func (d *Directory) Free() {
	if d == nil {
		return
	}
	if d.directory != nil {
		C.lvtk_uri_directory_free(d.directory)
		d.directory = nil
	}
}

// MapFeature returns the underlying map feature
func (d *Directory) MapFeature() *lv2.Feature {
	if d == nil {
		return nil
	}
	d.ensureDirectory()
	return lv2.NewFeatureRef(unsafe.Pointer(
		C.lvtk_uri_directory_get_map_feature(d.directory)))
}

// MapRef - return pointer to ctype LV2_URID_Map
func (d *Directory) MapRef() unsafe.Pointer {
	if d == nil || d.directory == nil {
		return nil
	}
	return unsafe.Pointer(C.lvtk_uri_directory_get_map(d.directory))
}

// UnmapFeature returns the underlying unmap feature
func (d *Directory) UnmapFeature() *lv2.Feature {
	if d == nil {
		return nil
	}

	d.ensureDirectory()
	return lv2.NewFeatureRef(unsafe.Pointer(
		C.lvtk_uri_directory_get_unmap_feature(d.directory)))
}

// UnmapRef - return pointer to ctype LV2_URID_Unmap
func (d *Directory) UnmapRef() unsafe.Pointer {
	if d == nil || d.directory == nil {
		return nil
	}
	return unsafe.Pointer(C.lvtk_uri_directory_get_unmap(d.directory))
}

// Map a URI
func (d *Directory) Map(uri string) uint32 {
	if d == nil {
		return 0
	}

	d.ensureDirectory()
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	return uint32(C.lvtk_uri_directory_map(d.directory, cstr))
}

// Unmap a URID
func (d *Directory) Unmap(urid uint32) string {
	if d == nil {
		return ""
	}
	d.ensureDirectory()
	return C.GoString(C.lvtk_uri_directory_unmap(d.directory, C.uint32_t(urid)))
}

package urid

/*
#cgo pkg-config: lv2
#cgo pkg-config: lvtk-2

#include <lv2/lv2plug.in/ns/ext/urid/urid.h>
#include <stdlib.h>
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

func (d *Directory) createDirectoryIfNeeded() {
	if d.directory == nil {
		d.directory = C.lvtk_uri_directory_new()
	}
}

// Free allocated data
func (d *Directory) Free() {
	if d.directory != nil {
		C.lvtk_uri_directory_free(d.directory)
		d.directory = nil
	}
}

// GetMapFeature get the underlying map feature
func (d *Directory) GetMapFeature() *lv2.Feature {
	if d == nil {
		return nil
	}
	d.createDirectoryIfNeeded()
	return lv2.NewFeatureRef(unsafe.Pointer(
		C.lvtk_uri_directory_get_map_feature(d.directory)))
}

// GetUnmapFeature get the underlying unmap feature
func (d *Directory) GetUnmapFeature() *lv2.Feature {
	if d == nil {
		return nil
	}

	d.createDirectoryIfNeeded()
	return lv2.NewFeatureRef(unsafe.Pointer(
		C.lvtk_uri_directory_get_unmap_feature(d.directory)))
}

// Map a URI
func (d *Directory) Map(uri string) uint32 {
	if d == nil {
		return 0
	}

	d.createDirectoryIfNeeded()
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))
	return uint32(C.lvtk_uri_directory_map(d.directory, cstr))
}

// Unmap a URID
func (d *Directory) Unmap(urid uint32) string {
	if d == nil {
		return ""
	}
	d.createDirectoryIfNeeded()
	return C.GoString(C.lvtk_uri_directory_unmap(d.directory, C.uint32_t(urid)))
}

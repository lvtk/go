package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <memory.h>
*/
import "C"
import (
	"unsafe"

	"github.com/lvtk/go/urid"
)

// Free - destroy the state
func (s *State) Free() {
	if s == nil || s.state == nil {
		return
	}
	if s.state != nil {
		C.lilv_state_free(s.state)
	}
	s.world = nil
	s.state = nil
}

// Equals - true if this state equals `b`
func (s *State) Equals(b *State) bool {
	if s == nil || s.state == nil || b == nil || b.state == nil {
		return false
	}
	return bool(C.lilv_state_equals(s.state, b.state))
}

// NumProperties - Return the number of properties in `state`.
func (s *State) NumProperties() uint32 {
	if s == nil || s.state == nil {
		return 0
	}
	return uint32(C.lilv_state_get_num_properties(s.state))
}

// PluginURI - Return the number of properties in `state`.
func (s *State) PluginURI() *Node {
	if s == nil || s.state == nil {
		return nil
	}
	return createSharedNode(C.lilv_state_get_plugin_uri(s.state))
}

// URI - Get the URI of `state`.
// This may return NULL if the state has not been saved and has no URI.
func (s *State) URI() *Node {
	if s == nil || s.state == nil {
		return nil
	}
	cnode := C.lilv_state_get_uri(s.state)
	if cnode == nil {
		return nil
	}
	return createSharedNode(cnode)
}

// Label - Get the label of `state`.
func (s *State) Label() string {
	if s == nil || s.state == nil {
		return ""
	}
	return C.GoString(C.lilv_state_get_label(s.state))
}

// SetLabel - Set the label of `state`.
func (s *State) SetLabel(label string) {
	if s == nil || s.state == nil {
		return
	}
	cstr := C.CString(label)
	defer Free(unsafe.Pointer(cstr))
	C.lilv_state_set_label(s.state, cstr)
}

/*
SetMetadata - Set a metadata property on `state`.
@param state The state to set the metadata for.
@param key The key to store `value` under (URID).
@param value Pointer to the value to be stored.
@param size The size of `value` in bytes.
@param tp The type of `value` (URID).
@param flags LV2_State_Flags for `value`.
@return 0 on success.

This is a generic version of lilv_state_set_label(), which sets metadata
properties visible to hosts, but not plugins.  This allows storing useful
information such as comments or preset banks.
*/
func (s *State) SetMetadata(key uint32, value []byte, size, tp, flags uint32) bool {
	if s == nil || s.state == nil {
		return false
	}

	return 0 == C.lilv_state_set_metadata(s.state,
		C.uint32_t(key), unsafe.Pointer(&value),
		C.size_t(size), C.uint32_t(tp),
		C.uint32_t(flags))
}

/*
Save - saves state to a file.
@param world The world.
@param map URID mapper.
@param unmap URID unmapper.
@param state State to save.
@param uri URI of state, may be NULL.
@param dir Path of the bundle directory to save into.
@param filename Path of the state file relative to `dir`.

The format of state on disk is compatible with that defined in the LV2
preset extension, i.e. this function may be used to save presets which can
be loaded by any host.

If `uri` is NULL, the preset URI will be a file URI, but the bundle
can safely be moved (i.e. the state file will use "<>" as the subject).
*/
func (s *State) Save(_map urid.Map, unmap urid.Unmap, uri, dir, filename string) bool {
	curi := C.CString(uri)
	cdir := C.CString(dir)
	cfilename := C.CString(filename)
	defer Free(unsafe.Pointer(curi))
	defer Free(unsafe.Pointer(cdir))
	defer Free(unsafe.Pointer(cfilename))
	return 0 == C.lilv_state_save(s.world,
		(*C.LV2_URID_Map)(_map),
		(*C.LV2_URID_Unmap)(unmap),
		s.state, curi, cdir, cfilename)
}

/*
ToString - Save state to a string.  This function does not use the filesystem.

@param world The world.
@param map URID mapper.
@param unmap URID unmapper.
@param state The state to serialize.
@param uri URI for the state description (mandatory).
@param baseURI Base URI for serialisation.  Unless you know what you are
doing, pass NULL for this, otherwise the state may not be restorable via
lilv_state_new_from_string().
*/
func (s *State) ToString(_map urid.Map, unmap urid.Unmap, uri, baseURI string) string {
	if _map == nil || unmap == nil || s == nil || s.world == nil || s.state == nil {
		return ""
	}

	curi := C.CString(uri)
	defer Free(unsafe.Pointer(curi))

	var cbaseuri *C.char
	if len(baseURI) > 0 {
		cbaseuri = C.CString(baseURI)
		defer Free(unsafe.Pointer(cbaseuri))
	}

	cstatestr := C.lilv_state_to_string(s.world,
		(*C.LV2_URID_Map)(_map),
		(*C.LV2_URID_Unmap)(unmap),
		s.state, curi, cbaseuri)

	if cstatestr != nil {
		defer Free(unsafe.Pointer(cstatestr))
		return C.GoString(cstatestr)
	}

	return ""
}

/*
Delete - Unload a state from the world and delete all associated files.
@param world The world.
@param state State to remove from the system.

This function DELETES FILES/DIRECTORIES FROM THE FILESYSTEM!  It is intended
for removing user-saved presets, but can delete any state the user has
permission to delete, including presets shipped with plugins.

The rdfs:seeAlso file for the state will be removed.  The entry in the
bundle's manifest.ttl is removed, and if this results in an empty manifest,
then the manifest file is removed.  If this results in an empty bundle, then
the bundle directory is removed as well.
*/
func (s *State) Delete() bool {
	if s == nil || s.state == nil || s.world == nil {
		return false
	}
	return 0 == int(C.lilv_state_delete(s.world, s.state))
}

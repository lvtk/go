package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// See World.SetOption
const (
	// OptionFilterLang - Enable/disable language filtering.
	// Language filtering applies to any functions that return (a) value(s).
	// With filtering enabled, Lilv will automatically return the best value(s)
	// for the current LANG.  With filtering disabled, all matching values will
	// be returned regardless of language tag.  Filtering is enabled by default.
	OptionFilterLang = "http://drobilla.net/ns/lilv#filter-lang"

	// OptionDynManifest - Enable/disable dynamic manifest support.
	// Dynamic manifest data will only be loaded if this option is true.
	OptionDynManifest = "http://drobilla.net/ns/lilv#dyn-manifest"

	// OptionLv2Path - Set application-specific LV2_PATH.  This overrides the
	// LV2_PATH from the environment, so that lilv will only look inside the
	// given path.  This can be used to make self-contained applications.
	OptionLv2Path = "http://drobilla.net/ns/lilv#lv2-path"
)

/*
Free - Free memory allocated by Lilv.

This function exists because some systems require memory allocated by a
library to be freed by code in the same library.  It is otherwise equivalent
to the standard C free() function.
*/
func Free(ptr unsafe.Pointer) {
	C.lilv_free(ptr)
}

/*
ParseFileURI - wraps lilv_file_uri_parse
Convert a file URI string to a local path string.
For example, "file://foo/bar%20one/baz.ttl" returns "/foo/bar one/baz.ttl".
Return value must be freed by caller with lilv_free().
@param uri The file URI to parse.
*/
func ParseFileURI(uri string) (string, string) {
	curi := C.CString(uri)
	defer Free(unsafe.Pointer(curi))
	var chostname **C.char
	res := C.lilv_file_uri_parse(curi, chostname)

	var path string
	var host string

	if res != nil {
		path = C.GoString(res)
		if *chostname != nil {
			host = C.GoString(*chostname)
		}
		defer Free(unsafe.Pointer(res))
	}

	return path, host
}

/* World */

// NewWorld - Initialize a new, empty world. If initialization fails, nil is
// returned.
func NewWorld() *World {
	r := new(World)
	r.world = C.lilv_world_new()
	return r
}

// SetOption - lilv_world_set_option
func (w *World) SetOption(uri string, value *Node) {
	if w == nil || w.world == nil || value == nil || value.node == nil {
		return
	}
	cstr := C.CString(uri)
	defer C.lilv_free(unsafe.Pointer(cstr))
	C.lilv_world_set_option(w.world, cstr, value.node)
}

// Free - wraps lilv_world_free
func (w *World) Free() {
	if w.world != nil {
		C.lilv_world_free(w.world)
		w.world = nil
	}
}

// LoadAll - lilv_world_load_all
func (w *World) LoadAll() {
	C.lilv_world_load_all(w.world)
}

// LoadBundle - lilv_world_load_bundle
func (w *World) LoadBundle(uri *Node) {
	C.lilv_world_load_bundle(w.world, uri.node)
}

// LoadSpecifications - lilv_world_load_specifications
func (w *World) LoadSpecifications() {
	C.lilv_world_load_specifications(w.world)
}

// LoadPluginClasses - lilv_world_load_plugin_classes
func (w *World) LoadPluginClasses() {
	C.lilv_world_load_plugin_classes(w.world)
}

// UnloadBundle - lilv_world_unload_bundle
func (w *World) UnloadBundle(uri *Node) int {
	return int(C.lilv_world_unload_bundle(w.world, uri.node))
}

// LoadResource - lilv_world_load_resource
func (w *World) LoadResource(uri *Node) {
	C.lilv_world_load_resource(w.world, uri.node)
}

// UnloadResource - lilv_world_unload_resource
func (w *World) UnloadResource(uri *Node) int {
	return int(C.lilv_world_unload_resource(w.world, uri.node))
}

// PluginClass - Get the parent of all other plugin classes, lv2:Plugin.
func (w *World) PluginClass() *PluginClass {
	pc := new(PluginClass)
	pc.pluginClass = (*C.LilvPluginClass)(C.lilv_world_get_plugin_class(w.world))
	return pc
}

/*
AllPlugins - Return a list of all found plugins.
The returned list contains just enough references to query
or instantiate plugins.  The data for a particular plugin will not be
loaded into memory until a call to an lilv_plugin_* function results in
a query (at which time the data is cached with the LilvPlugin so future
queries are very fast).

The returned list and the plugins it contains are owned by `World`
and must not be modifed
*/
func (w *World) AllPlugins() *Plugins {
	if w == nil || w.world == nil {
		return nil
	}
	plugs := new(Plugins)
	plugs.plugins = C.lilv_world_get_all_plugins(w.world)
	return plugs
}

/*
PluginClasses - Return a list of all found plugin classes.
Returned list is owned by world and must not be modified by the caller.
*/
func (w *World) PluginClasses() *PluginClasses {
	if w == nil || w.world == nil {
		return nil
	}

	return createPluginClasses(false, C.lilv_world_get_plugin_classes(w.world))
}

/*
FindNodes - Find nodes matching a triple pattern.
Either `subject` or `object` may be nil (i.e. a wildcard), but not both.
All matches for the wildcard field, or nil.
Retured value MUST be freed with Nodes.Free()
*/
func (w *World) FindNodes(subject *Node, predicate *Node, object *Node) *Nodes {
	var s *C.LilvNode
	if subject != nil {
		s = subject.node
	}
	var p *C.LilvNode
	if predicate != nil {
		p = predicate.node
	}
	var o *C.LilvNode
	if object != nil {
		o = object.node
	}

	return createNodes(true, C.lilv_world_find_nodes(w.world, s, p, o))
}

/*
Get - Find a single node that matches a pattern.
Exactly one of `subject`, `predicate`, `object` must be NULL.
This function is equivalent to
lilv_nodes_get_first(lilv_world_find_nodes(...)) but simplifies the common
case of only wanting a single value.
@return the first matching node, or NULL if no matches are found.
*/
func (w *World) Get(subject *Node, predicate *Node, object *Node) *Node {
	if w == nil || w.world == nil {
		return nil
	}

	var s *C.LilvNode
	if subject != nil {
		s = subject.node
	}
	var p *C.LilvNode
	if predicate != nil {
		p = predicate.node
	}
	var o *C.LilvNode
	if object != nil {
		o = object.node
	}

	return createManagedNode(C.lilv_world_get(w.world, s, p, o))
}

/*
Ask - Return true if a statement matching a certain pattern exists.

This is useful for checking if particular statement exists without having to
bother with collections and memory management.

@param world The world.
@param subject Subject of statement, or NULL for anything.
@param predicate Predicate (key) of statement, or NULL for anything.
@param object Object (value) of statement, or NULL for anything.
*/
func (w *World) Ask(subject *Node, predicate *Node, object *Node) bool {
	if w == nil || w.world == nil {
		return false
	}

	var s *C.LilvNode
	if subject != nil {
		s = subject.node
	}
	var p *C.LilvNode
	if predicate != nil {
		p = predicate.node
	}
	var o *C.LilvNode
	if object != nil {
		o = object.node
	}

	return bool(C.lilv_world_ask(w.world, s, p, o))
}

/*
Symbol - Get an LV2 symbol for some subject.

This will return the lv2:symbol property of the subject if it is given
explicitly, and otherwise will attempt to derive a symbol from the URI.
Returns A string node that is a valid LV2 symbol, or nil on error.
Return value must be freed with Node.Free()
*/
func (w *World) Symbol(subject *Node) *Node {
	if w == nil || w.world == nil || subject == nil || subject.node == nil {
		return nil
	}
	return createManagedNode(C.lilv_world_get_symbol(w.world, subject.node))
}

// NewURI - wraps lilv_new_uri
// Return value must be freed with Node.Free()
func (w *World) NewURI(uri string) *Node {
	if w == nil || w.world == nil {
		return nil
	}
	cstr := C.CString(uri)
	defer C.lilv_free(unsafe.Pointer(cstr))
	return createManagedNode(C.lilv_new_uri(w.world, cstr))
}

// NewFileURI - Create a new File URI node
// Return value must be freed with Node.Free()
func (w *World) NewFileURI(host string, path string) *Node {
	if w == nil || w.world == nil {
		return nil
	}
	chost := C.CString(host)
	cpath := C.CString(path)
	defer Free(unsafe.Pointer(chost))
	defer Free(unsafe.Pointer(cpath))
	return createManagedNode(C.lilv_new_file_uri(w.world, chost, cpath))
}

// NewString - wraps lilv_new_string
// Return value must be freed with Node.Free()
func (w *World) NewString(str string) *Node {
	if w == nil || w.world == nil {
		return nil
	}

	cstr := C.CString(str)
	defer Free(unsafe.Pointer(cstr))
	return createManagedNode(C.lilv_new_string(w.world, cstr))
}

// NewInt - create a new Int node
// Return value must be freed with Node.Free()
func (w *World) NewInt(val int32) *Node {
	if w == nil || w.world == nil {
		return nil
	}

	return createManagedNode(C.lilv_new_int(w.world, C.int(val)))
}

// NewFloat - creates a new Float node
// Return value must be freed with Node.Free()
func (w *World) NewFloat(val float32) *Node {
	if w == nil || w.world == nil {
		return nil
	}

	return createManagedNode(C.lilv_new_float(w.world, C.float(val)))
}

// NewBool - creates a new Bool node
// Return value must be freed with Node.Free()
func (w *World) NewBool(val bool) *Node {
	if w == nil || w.world == nil {
		return nil
	}
	return createManagedNode(C.lilv_new_bool(w.world, C.bool(val)))
}

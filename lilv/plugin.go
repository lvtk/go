package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/lvtk/go/lv2"
)

/* Plugins */

// Size - returns total plugins in this collection
func (ps *Plugins) Size() uint32 {
	if ps == nil || ps.plugins == nil {
		return 0
	}
	return uint32(C.lilv_plugins_size(unsafe.Pointer(ps.plugins)))
}

// Begin - returns begin iterator
func (ps *Plugins) Begin() *Iter {
	return (*Iter)(C.lilv_plugins_begin(unsafe.Pointer(ps.plugins)))
}

// Get - convert iter to `Plugin`
func (ps *Plugins) Get(iter *Iter) *Plugin {
	p := new(Plugin)
	p.plugin = C.lilv_plugins_get(unsafe.Pointer(ps.plugins), unsafe.Pointer(iter))
	return p
}

// Next - returns the next plugin iter
func (ps *Plugins) Next(iter *Iter) *Iter {
	return (*Iter)(C.lilv_plugins_next(unsafe.Pointer(ps.plugins), unsafe.Pointer(iter)))
}

// IsEnd - true if is the end iterator
func (ps *Plugins) IsEnd(iter *Iter) bool {
	return bool((C.lilv_plugins_is_end(unsafe.Pointer(ps.plugins), unsafe.Pointer(iter))))
}

// GetByURI - Get a plugin from `plugins` by URI.
// Return value is shared (stored in `plugins`) and must not be freed or
// modified by the caller in any way.
// @return NULL if no plugin with `uri` is found in `plugins`.
func (ps *Plugins) GetByURI(uri *Node) *Plugin {
	cplug := C.lilv_plugins_get_by_uri(unsafe.Pointer(ps.plugins), uri.node)
	if cplug == nil {
		return nil
	}
	plug := new(Plugin)
	plug.plugin = cplug
	return plug
}

/* Plugin */

/*
Instantiate - Instantiate a plugin.
The returned value is a lightweight handle for an LV2 plugin instance,
it does not refer to `plugin`, or any other Lilv state.  The caller must
eventually free it with lilv_instance_free().
`features` is a nil-terminated array of features the host supports.
nil may be passed if the host supports no additional features.
Returns nil if instantiation failed.
*/
func (p *Plugin) Instantiate(sampleRate float64, features *lv2.FeatureList) *Instance {
	if p == nil || p.plugin == nil {
		return nil
	}

	var fs **C.LV2_Feature
	if features != nil {
		fs = (**C.LV2_Feature)(features.Ref())
	}

	ci := C.lilv_plugin_instantiate(p.plugin, (C.double)(sampleRate), fs)
	if ci == nil {
		return nil
	}

	i := new(Instance)
	i.instance = ci
	return i
}

// Verify - Check if `plugin` is valid.
// This is not a rigorous validator, but can be used to reject some malformed
// plugins that could cause bugs (e.g. plugins with missing required fields).
//
// Note that normal hosts do NOT need to use this - lilv does not
// load invalid plugins into plugin lists.  This is included for plugin
// testing utilities, etc.
//
// return true if valid
func (p *Plugin) Verify() bool {
	return p != nil && p.plugin != nil && bool(C.lilv_plugin_verify(p.plugin))
}

// URI - Get the URI of `plugin`.
// Any serialization that refers to plugins should refer to them by this.
// Hosts SHOULD NOT save any filesystem paths, plugin indexes, etc. in saved
// files; save only the URI.
//
// The URI is a globally unique identifier for one specific plugin.  Two
// plugins with the same URI are compatible in port signature, and should
// be guaranteed to work in a compatible and consistent way.  If a plugin
// is upgraded in an incompatible way (eg if it has different ports), it
// MUST have a different URI than it's predecessor.
//
// @return A shared URI value which must NOT be modified
func (p *Plugin) URI() *Node {
	return createSharedNode(C.lilv_plugin_get_uri(p.plugin))
}

// BundleURI - Get the (resolvable) URI of the plugin's "main" bundle.
// This returns the URI of the bundle where the plugin itself was found.  Note
// that the data for a plugin may be spread over many bundles, that is,
// lilv_plugin_get_data_uris() may return URIs which are not within this
// bundle.
//
// Typical hosts should not need to use this function.
// Note this always returns a fully qualified URI.  If you want a local
// filesystem path, use lilv_file_uri_parse().
// @return a shared string which must not be modified or freed.
func (p *Plugin) BundleURI() *Node {
	return createSharedNode(C.lilv_plugin_get_bundle_uri(p.plugin))
}

// DataURIs - Get the (resolvable) URIs of the RDF data files that define a plugin.
// Typical hosts should not need to use this function.
// Note this always returns fully qualified URIs.  If you want local
// filesystem paths, use lilv_file_uri_parse().
// @return a list of complete URLs eg. "file:///foo/ABundle.lv2/aplug.ttl",
// which is shared and must not be modified or freed.
func (p *Plugin) DataURIs() *Nodes {
	return createNodes(true, C.lilv_plugin_get_data_uris(p.plugin))
}

// LibraryURI - Get the (resolvable) URI of the shared library for `plugin`.
// Note this always returns a fully qualified URI.  If you want a local
// filesystem path, use lilv_file_uri_parse().
// @return a shared string which must not be modified or freed.
func (p *Plugin) LibraryURI() *Node {
	return createSharedNode(C.lilv_plugin_get_library_uri(p.plugin))
}

// Name - Get the name of `plugin`.
// This returns the name (doap:name) of the plugin.  The name may be
// translated according to the current locale, this value MUST NOT be used
// as a plugin identifier (use the URI for that).
// Returned value MUST be freed by the caller with node.Free()
func (p *Plugin) Name() *Node {
	return createManagedNode(C.lilv_plugin_get_name(p.plugin))
}

// Class - Get the class this plugin belongs to (e.g. Filters).
func (p *Plugin) Class() *PluginClass {
	pc := new(PluginClass)
	pc.pluginClass = C.lilv_plugin_get_class(p.plugin)
	return pc
}

// Value - Get a value associated with the plugin in a plugin's data files.
// `predicate` must be either a URI or a QName.
//
// Returns the ?object of all triples found of the form:
//
// <code>&lt;plugin-uri&gt; predicate ?object</code>
//
// May return NULL if the property was not found, or if object(s) is not
// sensibly represented as a LilvNodes (e.g. blank nodes).
// Return value must be freed by caller with Nodes.Free().
func (p *Plugin) Value(predicate *Node) *Nodes {
	return createNodes(true, C.lilv_plugin_get_value(p.plugin, predicate.node))
}

/*
HasFeature - Return whether a feature is supported by a plugin.
This will return true if the feature is an optional or required feature
of the plugin.
*/
func (p *Plugin) HasFeature(feature *Node) bool {
	return bool(C.lilv_plugin_has_feature(p.plugin, feature.node))
}

/*
SupportedFeatures - Get the LV2 Features supported (required or optionally) by a plugin.
A feature is "supported" by a plugin if it is required OR optional.

Since required features have special rules the host must obey, this function
probably shouldn't be used by normal hosts.  Using lilv_plugin_get_optional_features()
and lilv_plugin_get_required_features() separately is best in most cases.

Returned value must be freed by caller with Nodes.Free().
*/
func (p *Plugin) SupportedFeatures() *Nodes {
	return createNodes(true, C.lilv_plugin_get_supported_features(p.plugin))
}

/*
RequiredFeatures - Get the LV2 Features required by a plugin.
If a feature is required by a plugin, hosts MUST NOT use the plugin if they do not
understand (or are unable to support) that feature.

All values returned here MUST be passed to the plugin's instantiate method
(along with data, if necessary, as defined by the feature specification)
or plugin instantiation will fail.

Return value must be freed by caller with Nodes.Free().
*/
func (p *Plugin) RequiredFeatures() *Nodes {
	return createNodes(true, C.lilv_plugin_get_required_features(p.plugin))
}

/*
OptionalFeatures - Get the LV2 Features optionally supported by a plugin.
Hosts MAY ignore optional plugin features for whatever reasons.  Plugins
MUST operate (at least somewhat) if they are instantiated without being
passed optional features.

Return value must be freed by caller with Nodes.Free().
*/
func (p *Plugin) OptionalFeatures() *Nodes {
	return createNodes(true, C.lilv_plugin_get_required_features(p.plugin))
}

/*
HasExtensionData - Return whether or not a plugin provides a specific extension data.
*/
func (p *Plugin) HasExtensionData(uri *Node) bool {
	return bool(C.lilv_plugin_has_extension_data(p.plugin, uri.node))
}

/*
ExtensionData - Get a sequence of all extension data provided by a plugin.
This can be used to find which URIs lilv_instance_get_extension_data()
will return a value for without instantiating the plugin.

Return value must be freed by caller with Nodes.Free().
*/
func (p *Plugin) ExtensionData() *Nodes {
	return createNodes(true, C.lilv_plugin_get_extension_data(p.plugin))
}

// NumPorts - Get the number of ports on this plugin
func (p *Plugin) NumPorts() uint32 {
	return uint32(C.lilv_plugin_get_num_ports(p.plugin))
}

// PortRangesFloat - Get the port ranges (minimum, maximum and default values) for all ports.
// `min_values`, `max_values` and `def_values` must either point to an array
// of N floats, where N is the value returned by lilv_plugin_get_num_ports()
// for this plugin, or NULL.  The elements of the array will be set to the
// the minimum, maximum and default values of the ports on this plugin,
// with array index corresponding to port index.  If a port doesn't have a
// minimum, maximum or default value, or the port's type is not float, the
// corresponding array element will be set to NAN.
//
// This is a convenience method for the common case of getting the range of
// all float ports on a plugin, and may be significantly faster than
// repeated calls to lilv_port_get_range().
func (p *Plugin) PortRangesFloat() ([]float32, []float32, []float32) {
	nports := p.NumPorts()
	cmin := make([]C.float, nports)
	cmax := make([]C.float, nports)
	cdef := make([]C.float, nports)
	C.lilv_plugin_get_port_ranges_float(p.plugin, &cmin[0], &cmax[0], &cdef[0])
	min := make([]float32, nports)
	max := make([]float32, nports)
	def := make([]float32, nports)
	for i := 0; i < int(nports); i++ {
		min[i] = float32(cmin[i])
		max[i] = float32(cmin[i])
		def[i] = float32(cmin[i])
	}
	return min, max, def
}

/*
HasLatency - Return whether or not the plugin introduces (and reports) latency.
The index of the latency port can be found with Plugin.GetLatencyPort() ONLY
if this function returns true.
*/
func (p *Plugin) HasLatency() bool {
	return bool(C.lilv_plugin_has_latency(p.plugin))
}

/*
LatencyPortIndex - Return the index of the plugin's latency port.
It is a fatal error to call this on a plugin without checking if the port
exists by first calling lilv_plugin_has_latency().

Any plugin that introduces unwanted latency that should be compensated for
(by hosts with the ability/need) MUST provide this port, which is a control
rate output port that reports the latency for each cycle in frames.
*/
func (p *Plugin) LatencyPortIndex() uint32 {
	return uint32(C.lilv_plugin_get_latency_port_index(p.plugin))
}

/*
Project - Get the project the plugin is a part of.

More information about the project can be read via lilv_world_find_nodes(),
typically using properties from DOAP (e.g. doap:name).
*/
func (p *Plugin) Project() *Node {
	return createManagedNode(C.lilv_plugin_get_project(p.plugin))
}

/*
AuthorName - Get the full name of the plugin's author.
Returns NULL if author name is not present.
Returned value must be freed by caller with Node.Free()
*/
func (p *Plugin) AuthorName() *Node {
	return createManagedNode(C.lilv_plugin_get_author_name(p.plugin))
}

/*
AuthorEmail - Get the email address of the plugin's author.
Returns NULL if author email address is not present.
Returned value must be freed by caller with Node.Free()
*/
func (p *Plugin) AuthorEmail() *Node {
	return createManagedNode(C.lilv_plugin_get_author_email(p.plugin))
}

/*
AuthorHomepage - Get the address of the plugin author's home page.
Returns NULL if author homepage is not present.
Returned value must be freed by caller with Node.Free()
*/
func (p *Plugin) AuthorHomepage() *Node {
	return createManagedNode(C.lilv_plugin_get_author_homepage(p.plugin))
}

/*
IsReplaced - Return true iff `plugin` has been replaced by another plugin.

The plugin will still be usable, but hosts should hide them from their
user interfaces to prevent users from using deprecated plugins.
*/
func (p *Plugin) IsReplaced() bool {
	return bool(C.lilv_plugin_is_replaced(p.plugin))
}

/*
Related - Get the resources related to `plugin` with lv2:appliesTo.

Some plugin-related resources are not linked directly to the plugin with
rdfs:seeAlso and thus will not be automatically loaded along with the plugin
data (usually for performance reasons).  All such resources of the given @c
type related to `plugin` can be accessed with this function.

If `type` is NULL, all such resources will be returned, regardless of type.

To actually load the data for each returned resource, use World.LoadResource().
*/
func (p *Plugin) Related(tp *Node) *Nodes {
	return createNodes(true, C.lilv_plugin_get_related(p.plugin, tp.node))
}

/*
PortByIndex - Get a port on `plugin` by `index`.
*/
func (p *Plugin) PortByIndex(index uint32) *Port {
	port := new(Port)
	port.plugin = p.plugin
	port.port = C.lilv_plugin_get_port_by_index(p.plugin, (C.uint32_t)(index))
	return port
}

/*
PortBySymbol - Get a port on `plugin` by `symbol`.
Note this function is slower than lilv_plugin_get_port_by_index(),
especially on plugins with a very large number of ports.
*/
func (p *Plugin) PortBySymbol(symbol *Node) *Port {
	port := new(Port)
	port.plugin = p.plugin
	port.port = C.lilv_plugin_get_port_by_symbol(p.plugin, symbol.node)
	return port
}

/*
PortByDesignation - Get a port on `plugin` by its lv2:designation.
The designation of a port describes the meaning,
assignment, allocation or role of the port, e.g. "left channel" or "gain".
If found, the port with matching `port_class` and `designation` is be returned,
otherwise NULL is returned.  The `port_class` can be used to distinguish the
input and output ports for a particular designation.  If `port_class` is NULL,
any port with the given designation will be returned.
*/
func (p *Plugin) PortByDesignation(portClass *Node, designation *Node) *Port {
	port := new(Port)
	port.port = C.lilv_plugin_get_port_by_designation(
		p.plugin, portClass.node, designation.node)
	return port
}

/*
GetUIs - Get all UIs for `plugin`.
Returned value must be freed by caller using UIs.Free().
*/
func (p *Plugin) GetUIs() *UIs {
	if p == nil || p.plugin == nil {
		return nil
	}
	uis := new(UIs)
	uis.uis = C.lilv_plugin_get_uis(p.plugin)
	uis.managed = true
	return uis
}

/* PluginClasses */

// Free - destroy this collection
func (pcs *PluginClasses) Free() {
	if pcs != nil && pcs.pluginClasses != nil {
		if pcs.managed {
			C.lilv_plugin_classes_free(unsafe.Pointer(pcs.pluginClasses))
		}
		pcs.pluginClasses = nil
	}
}

// Size - start of collection
func (pcs *PluginClasses) Size() uint32 {
	if pcs == nil || pcs.pluginClasses == nil {
		return 0
	}
	return (uint32)(C.lilv_plugin_classes_size(unsafe.Pointer(pcs.pluginClasses)))
}

// Begin - start of collection
func (pcs *PluginClasses) Begin() *Iter {
	if pcs == nil || pcs.pluginClasses == nil {
		return nil
	}
	return (*Iter)(C.lilv_plugin_classes_begin(unsafe.Pointer(pcs.pluginClasses)))
}

// Get - get value of iter
func (pcs *PluginClasses) Get(iter *Iter) *PluginClass {
	if pcs == nil || pcs.pluginClasses == nil {
		return nil
	}

	pc := new(PluginClass)
	pc.pluginClass = C.lilv_plugin_classes_get(
		unsafe.Pointer(pcs.pluginClasses),
		unsafe.Pointer(iter))
	return pc
}

// Next - next iter
func (pcs *PluginClasses) Next(iter *Iter) *Iter {
	if pcs == nil || pcs.pluginClasses == nil || iter == nil {
		return nil
	}
	return (*Iter)(C.lilv_plugin_classes_next(
		unsafe.Pointer(pcs.pluginClasses),
		unsafe.Pointer(iter)))
}

// IsEnd - is end iter
func (pcs *PluginClasses) IsEnd(iter *Iter) bool {
	if pcs == nil || pcs.pluginClasses == nil || iter == nil {
		return true
	}
	return (bool)(C.lilv_plugin_classes_is_end(
		unsafe.Pointer(pcs.pluginClasses),
		unsafe.Pointer(iter)))
}

// GetByURI - get value of iter
func (pcs *PluginClasses) GetByURI(uri *Node) *PluginClass {
	if pcs == nil || pcs.pluginClasses == nil {
		return nil
	}

	pc := new(PluginClass)
	pc.pluginClass = C.lilv_plugin_classes_get_by_uri(
		unsafe.Pointer(pcs.pluginClasses), uri.node)
	return pc
}

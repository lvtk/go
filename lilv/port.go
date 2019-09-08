package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
static uint32_t invalid_port_index() {
	return (uint32_t)-1;
}
*/
import "C"

/*
Node - Get the RDF node of `port`.
Ports nodes may be may be URIs or blank nodes.
@return A shared node which must not be modified.
*/
func (p *Port) Node() *Node {
	if p == nil || p.port == nil || p.plugin == nil {
		return nil
	}
	return createSharedNode(C.lilv_port_get_node(p.plugin, p.port))
}

/*
Value - Port analog of Plugin.Value()
Returned value must be freed with Nodes.Free()
*/
func (p *Port) Value(predicate *Node) *Nodes {
	if p == nil || p.port == nil || p.plugin == nil || predicate == nil || predicate.node == nil {
		return nil
	}
	return createNodes(true, C.lilv_port_get_value(
		p.plugin, p.port, predicate.node))
}

/*
Get - Get a single property value of a port.

This is equivalent to lilv_nodes_get_first(lilv_port_get_value(...)) but is
simpler to use in the common case of only caring about one value.  The
caller is responsible for freeing the returned node.
*/
func (p *Port) Get(predicate *Node) *Node {
	if p == nil || p.plugin == nil || p.port == nil || predicate == nil || predicate.node == nil {
		return nil
	}
	return createManagedNode(C.lilv_port_get(p.plugin, p.port, predicate.node))
}

/*
Properties - Return the LV2 port properties of a port.
Retured value must be freed with Nodes.Free()
*/
func (p *Port) Properties() *Nodes {
	if p == nil || p.plugin == nil || p.port == nil {
		return nil
	}
	return createNodes(true, C.lilv_port_get_properties(p.plugin, p.port))
}

/*
HasProperty - Return whether a port has a certain property.
*/
func (p *Port) HasProperty(property *Node) bool {
	if p == nil || p.plugin == nil || p.port == nil || property == nil || property.node == nil {
		return false
	}
	return bool(C.lilv_port_has_property(
		p.plugin, p.port, property.node))
}

/*
SupportsEvent - Return whether a port supports a certain event type.
More precisely, this returns true iff the port has an atom:supports or an
ev:supportsEvent property with `event_type` as the value.
*/
func (p *Port) SupportsEvent(eventType *Node) bool {
	if p == nil || p.plugin == nil || p.port == nil || eventType == nil || eventType.node == nil {
		return false
	}
	return bool(C.lilv_port_supports_event(
		p.plugin, p.port, eventType.node))
}

/*
Index - Get the index of a port.
The index is only valid for the life of the plugin and may change between
versions.  For a stable identifier, use the symbol.
*/
func (p *Port) Index() uint32 {
	if p == nil || p.plugin == nil || p.port == nil {
		return uint32(C.invalid_port_index())
	}
	return uint32(C.lilv_port_get_index(p.plugin, p.port))
}

/*
Symbol - Get the symbol of a port.
The 'symbol' is a short string, a valid C identifier.
Returned value is owned by `port` and must not be modified.
*/
func (p *Port) Symbol() *Node {
	if p == nil || p.plugin == nil || p.port == nil {
		return nil
	}
	return createSharedNode(C.lilv_port_get_symbol(p.plugin, p.port))
}

/*
Name - Get the name of a port.
This is guaranteed to return the untranslated name (the doap:name in the
data file without a language tag).  Returned value must be freed by
the caller with Node.Free()
*/
func (p *Port) Name() *Node {
	if p == nil || p.plugin == nil || p.port == nil {
		return nil
	}
	return createManagedNode(C.lilv_port_get_name(p.plugin, p.port))
}

/*
Classes - Get all the classes of a port.
This can be used to determine if a port is an input, output, audio,
control, midi, etc, etc, though it's simpler to use lilv_port_is_a().
The returned list does not include lv2:Port, which is implied.
Returned value is shared and must not be modified by caller.
*/
func (p *Port) Classes() *Nodes {
	if p == nil || p.plugin == nil || p.port == nil {
		return nil
	}
	return createNodes(false, C.lilv_port_get_classes(p.plugin, p.port))
}

/*
IsA - Determine if a port is of a given class (input, output, audio, etc).
For convenience/performance/extensibility reasons, hosts are expected to
create a LilvNode for each port class they "care about".  Well-known type
URI strings are defined (e.g. LILV_URI_INPUT_PORT) for convenience, but
this function is designed so that Lilv is usable with any port types
without requiring explicit support in Lilv.
*/
func (p *Port) IsA(portClass *Node) bool {
	if p == nil || p.plugin == nil || p.port == nil || portClass == nil || portClass.node == nil {
		return false
	}
	return bool(C.lilv_port_is_a(p.plugin, p.port, portClass.node))
}

/*
ScalePoints - Get the scale points (enumeration values) of a port.
This returns a collection of 'interesting' named values of a port
(e.g. appropriate entries for a UI selector associated with this port).
Returned value may be NULL if `port` has no scale points, otherwise it
must be freed by caller with ScalePoints.Free().
*/
func (p *Port) ScalePoints() *ScalePoints {
	if p == nil || p.plugin == nil || p.port == nil {
		return nil
	}
	sps := new(ScalePoints)
	sps.managed = true
	sps.scalePoints = C.lilv_port_get_scale_points(p.plugin, p.port)
	return sps
}

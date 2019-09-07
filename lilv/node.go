package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

/* Node */

// Free - wraps lilv_node_free
// Does nothing if the node is owned by Lilv internals
func (n *Node) Free() {
	if n != nil && n.node != nil {
		if n.managed {
			C.lilv_node_free(n.node)
		}
		n.node = nil
	}
}

// Duplicate - wraps lilv_node_duplicate
// Returned value must be freed with Node.Free()
func (n *Node) Duplicate() *Node {
	return createManagedNode(C.lilv_node_duplicate(n.node))
}

// Equals - wraps lilv_node_equals
func (n *Node) Equals(o *Node) bool {
	if n == nil || n.node == nil || o == nil || o.node == nil {
		return false
	}
	return bool(C.lilv_node_equals(n.node, o.node))
}

// GetTurtleToken - get the turtle token
func (n *Node) GetTurtleToken() string {
	cstr := C.lilv_node_get_turtle_token(n.node)
	gstr := C.GoString(cstr)
	Free(unsafe.Pointer(cstr))
	return gstr
}

// IsURI - lilv_node_is_uri
func (n *Node) IsURI() bool {
	return bool(C.lilv_node_is_uri(n.node))
}

// URI - lilv_node_as_uri
func (n *Node) URI() string {
	return C.GoString(C.lilv_node_as_uri(n.node))
}

// IsBlank - lilv_node_is_blank
func (n *Node) IsBlank() bool {
	return bool(C.lilv_node_is_blank(n.node))
}

// Blank - lilv_node_as_blank
func (n *Node) Blank() string {
	return C.GoString(C.lilv_node_as_blank(n.node))
}

// IsLiteral - lilv_node_is_literal
func (n *Node) IsLiteral() bool {
	return bool(C.lilv_node_is_literal(n.node))
}

// IsString - lilv_node_is_string
func (n *Node) IsString() bool {
	return bool(C.lilv_node_is_string(n.node))
}

// String - lilv_node_as_string
func (n *Node) String() string {
	if n == nil || n.node == nil {
		return ""
	}
	return C.GoString(C.lilv_node_as_string(n.node))
}

// IsFloat - lilv_node_is_float
func (n *Node) IsFloat() bool {
	return bool(C.lilv_node_is_float(n.node))
}

// Float - lilv_node_is_float
func (n *Node) Float() float32 {
	return float32(C.lilv_node_as_float(n.node))
}

// IsInt - lilv_node_is_float
func (n *Node) IsInt() bool {
	return bool(C.lilv_node_is_int(n.node))
}

// Int - lilv_node_is_float
func (n *Node) Int() int32 {
	return int32(C.lilv_node_as_int(n.node))
}

// IsBool - lilv_node_is_bool
func (n *Node) IsBool() bool {
	return n != nil && n.node != nil && (bool)(C.lilv_node_is_bool(n.node))
}

// Bool - lilv_node_as_bool
func (n *Node) Bool() bool {
	if n == nil || n.node == nil {
		return false
	}
	return bool(C.lilv_node_as_bool(n.node))
}

/* Nodes */

// Free - Calls lilv_nodes_free
func (ns *Nodes) Free() {
	if ns != nil && ns.nodes != nil {
		if ns.managed {
			C.lilv_nodes_free(unsafe.Pointer(ns.nodes))
		}
		ns.nodes = nil
	}
}

// Size - num nodes in collection
func (ns *Nodes) Size() uint32 {
	if ns == nil || ns.nodes == nil {
		return 0
	}
	return (uint32)(C.lilv_nodes_size(unsafe.Pointer(ns.nodes)))
}

// Begin - start of collection
func (ns *Nodes) Begin() *Iter {
	return (*Iter)(C.lilv_nodes_begin(
		unsafe.Pointer(ns.nodes),
		unsafe.Pointer(iter)))
}

// Get - node value
func (ns *Nodes) Get(iter *Iter) *Node {
	return createSharedNode(C.lilv_nodes_get(
		unsafe.Pointer(ns.nodes),
		unsafe.Pointer(iter)))
}

// Next - next node iter
func (ns *Nodes) Next(iter *Iter) *Iter {
	return (*Iter)(C.lilv_nodes_next(
		unsafe.Pointer(ns.nodes),
		unsafe.Pointer(iter)))
}

// IsEnd - true if is end iterator
func (ns *Nodes) IsEnd(iter *Iter) bool {
	return bool(C.lilv_nodes_is_end(
		unsafe.Pointer(ns.nodes),
		unsafe.Pointer(iter)))
}

// GetFirst - first node in collection
func (ns *Nodes) GetFirst() *Node {
	return createSharedNode(C.lilv_nodes_get_first(
		unsafe.Pointer(ns.nodes)))
}

// Contains - true if contains given node
func (ns *Nodes) Contains(value *Node) bool {
	return bool(C.lilv_nodes_contains(
		unsafe.Pointer(ns.nodes), value.node))
}

// Merge - combine to nodes collections
// returned value MUSt be freed with Nodes.Free()
func (ns *Nodes) Merge(o *Nodes) *Nodes {
	return createNodes(true, C.lilv_nodes_merge(
		unsafe.Pointer(ns.nodes),
		unsafe.Pointer(o.nodes)))
}

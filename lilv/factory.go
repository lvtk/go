package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func createNode(managed bool, node *C.LilvNode) *Node {
	n := new(Node)
	n.managed = managed
	n.node = node
	return n
}

func createSharedNode(node *C.LilvNode) *Node {
	return createNode(false, node)
}

func createManagedNode(node *C.LilvNode) *Node {
	return createNode(true, node)
}

func createNodes(managed bool, nodes unsafe.Pointer) *Nodes {
	if nodes == nil {
		return nil
	}
	ns := new(Nodes)
	ns.managed = managed
	ns.nodes = nodes
	return ns
}

func createPluginClasses(managed bool, pluginClasses unsafe.Pointer) *PluginClasses {
	pcs := new(PluginClasses)
	pcs.pluginClasses = pluginClasses
	pcs.managed = managed
	return pcs
}

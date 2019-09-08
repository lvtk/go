package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Free this instance. Must be called when finished with the instance
// Does nothing if Instance is nil or has already been freed
func (i *Instance) Free() {
	if i == nil {
		return
	}
	if i.instance != nil {
		C.lilv_instance_free(i.instance)
		i.instance = nil
	}
}

// URI - get the plugin's URI identifier
func (i *Instance) URI() string {
	if i == nil || i.instance == nil {
		return ""
	}
	return C.GoString(C.lilv_instance_get_uri(i.instance))
}

// ConnectPort to data location
func (i *Instance) ConnectPort(index uint32, data unsafe.Pointer) {
	if i != nil && i.instance != nil {
		C.lilv_instance_connect_port(i.instance, C.uint32_t(index), data)
	}
}

// Run a processing cycle for `sampleCount` frames
func (i *Instance) Run(sampleCount uint32) {
	if i != nil && i.instance != nil {
		C.lilv_instance_run(i.instance, C.uint32_t(sampleCount))
	}
}

// Activate the plugin
func (i *Instance) Activate() {
	if i != nil && i.instance != nil {
		C.lilv_instance_activate(i.instance)
	}
}

// Deactivate the plugin
func (i *Instance) Deactivate() {
	if i != nil && i.instance != nil {
		C.lilv_instance_deactivate(i.instance)
	}
}

// ExtensionData - get plugin extension data
func (i *Instance) ExtensionData(uri string) unsafe.Pointer {
	if i == nil || i.instance == nil {
		return nil
	}
	curi := C.CString(uri)
	defer C.free(unsafe.Pointer(curi))
	return C.lilv_instance_get_extension_data(i.instance, curi)
}

// Handle - Get the native handle. Can't use this directly in Go, but
// could be passed to LV2_Features implemented in Go
func (i *Instance) Handle() unsafe.Pointer {
	if i == nil || i.instance == nil {
		return nil
	}
	return unsafe.Pointer(C.lilv_instance_get_handle(i.instance))
}

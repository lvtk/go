package pugl

/*
#cgo CFLAGS: -Isrc -Wno-deprecated-declarations
#cgo darwin LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include "pugl/pugl.h"
*/
import "C"
import "unsafe"

// World wrap PuglWorld
type World struct {
	world *C.PuglWorld
}

// NewWorld creates a new World
func NewWorld() *World {
	w := &World{world: C.puglNewWorld()}
	if w.world == nil {
		w = nil
	}
	return w
}

// Free the world
func (w *World) Free() {
	if w == nil {
		return
	}
	if w.world != nil {
		C.puglFreeWorld(w.world)
		w.world = nil
	}
}

// SetClassName must be unique per application
func (w *World) SetClassName(name string) Status {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Status(C.puglSetClassName(w.world, cname))
}

// GetTime get a timestamp
func (w *World) GetTime() float64 {
	return float64(C.puglGetTime(w.world))
}

// PollEvents - Poll for events that are ready to be processed.
//
// This polls for events that are ready for any view in the application,
// potentially blocking depending on `timeout`.
//
// @param timeout Maximum time to wait, in seconds.  If zero, the call returns
// immediately, if negative, the call blocks indefinitely.
// return Success if events are read, PUGL_FAILURE if not, or an error.
func (w *World) PollEvents(timeOut float64) Status {
	return Status(C.puglPollEvents(w.world, C.double(timeOut)))
}

// DispatchEvents to views.
//
// This processes all pending events, dispatching them to the appropriate
// views.  View event handlers will be called in the scope of this call.  This
// function does not block, if no events are pending it will return
// immediately.
func (w *World) DispatchEvents() Status {
	return Status(C.puglDispatchEvents(w.world))
}

func (w *World) createView() *View {
	v := new(View)
	v.world = w
	v.view = C.puglNewView(w.world)
	if v.view == nil {
		v.world = nil
		return nil
	}

	v.setInternalHandler()
	return v
}

// NewView - create a new view
func (w *World) NewView() *View {
	return w.createView()
}

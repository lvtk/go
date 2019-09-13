package pugl

/*
#cgo CFLAGS: -Isrc -Wno-deprecated-declarations
#cgo LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include "pugl/pugl.h"

#if defined(__APPLE__)
PuglStatus puglSetTransientFor (PuglView* view, PuglNativeWindow parent) {
	return PUGL_SUCCESS;
}
#endif

static void set_frame (PuglView* view, double x, double y, double w, double h) {
	PuglRect frame = { x, y, w, h };
	puglSetFrame(view, frame);
}
*/
import "C"
import "unsafe"

// Backend opaque backend pointer
type Backend C.PuglBackend

// View a pugl View wrapper
type View struct {
	world *World
	view  *C.PuglView
}

// SetParentWindow - Set the parent window before creating a window (for embedding).
func (x *View) SetParentWindow(parent uintptr) Status {
	return Status(C.puglSetParentWindow(x.view, C.PuglNativeWindow(parent)))
}

// SetFrame - Set the parent window before creating a window (for embedding).
func (x *View) SetFrame(frame Rect) {
	C.set_frame(x.view, C.double(frame.X), C.double(frame.Y),
		C.double(frame.Width), C.double(frame.Height))
}

// SetMinSize - Set the minimum window size before creating a window.
func (x *View) SetMinSize(width, height int) {
	C.puglSetMinSize(x.view, C.int(width), C.int(height))
}

// SetAspectRatio - Set the window aspect ratio range before creating a window.
//
// The x and y values here represent a ratio of width to height.  To set a
// fixed aspect ratio, set the minimum and maximum values to the same ratio.
//
// Note that setting different minimum and maximum constraints does not
// currenty work on MacOS (the minimum is used), so only setting a fixed aspect
// ratio works properly across all platforms.
func (x *View) SetAspectRatio(minX, minY, maxX, maxY int) {
	C.puglSetAspectRatio(x.view,
		C.int(minX), C.int(minY), C.int(maxX), C.int(maxY))
}

// SetTransientFor - Set transient parent before creating a window.
//
// On X11, parent must be a Window.
// On OSX, parent must be an NSView*.
func (x *View) SetTransientFor(parent NativeWindow) Status {
	return Status(C.puglSetTransientFor(x.view, C.PuglNativeWindow(parent)))
}

// SetViewHint - configure the view
func (x *View) SetViewHint(hint ViewHint, value int32) {
	C.puglSetViewHint(x.view, C.PuglViewHint(hint), C.int(value))
}

// SetBackend - init the backend
func (x *View) SetBackend(backend *Backend) {
	C.puglSetBackend(x.view, (*C.PuglBackend)(backend))
}

// GetProcAddress - Return the address of an OpenGL extension function.
func GetProcAddress(name string) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return unsafe.Pointer(C.puglGetProcAddress(cname))
}

// CreateWindow makes a window with the settings given by the various pugl.Init functions.
func (x *View) CreateWindow(title string) Status {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	return Status(C.puglCreateWindow(x.view, ctitle))
}

// ShowWindow shows the current window.
func (x *View) ShowWindow() Status {
	return Status(C.puglShowWindow(x.view))
}

// HideWindow - hide the current window
func (x *View) HideWindow() Status {
	return Status(C.puglHideWindow(x.view))
}

// SetWindowTitle sets the window title
func (x *View) SetWindowTitle(title string) Status {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	return Status(C.puglSetWindowTitle(x.view, ctitle))
}

// NativeWindow - A native window handle.
//
// On X11, this is a Window.
// On OSX, this is an NSView*.
// On Windows, this is a HWND.
type NativeWindow uintptr

// GetNativeWindow - Return the native window handle.
func (x *View) GetNativeWindow() NativeWindow {
	return NativeWindow(C.puglGetNativeWindow(x.view))
}

// GetVisible - true if visible
func (x *View) GetVisible() bool {
	return bool(C.puglGetVisible(x.view))
}

// GetFrame - Get the current size of the view
func (x *View) GetFrame() Rect {
	var r1 Rect
	r2 := C.puglGetFrame(x.view)
	r1.X = float64(r2.x)
	r1.Y = float64(r2.y)
	r1.Width = float64(r2.width)
	r1.Height = float64(r2.height)
	return r1
}

// EnterContext - enter the drawing context.
//
// Note that pugl automatically enters and leaves the drawing context during
// configure and expose events, so it is not normally necessary to call this.
// However, it can be used to enter the drawing context elsewhere, for example
// to call any GL functions during setup.
//
// @param drawing If true, prepare for drawing.
func (x *View) EnterContext(drawing bool) {
	C.puglEnterContext(x.view, C.bool(drawing))
}

// LeaveContext - Leave the drawing context.
//
// This must be called after puglEnterContext() with a matching `drawing`
// parameter.
//
// @param drawing If true, finish drawing, for example by swapping buffers.
func (x *View) LeaveContext(drawing bool) {
	C.puglLeaveContext(x.view, C.bool(drawing))
}

// WithContext runs a function between enter & exit context
func (x *View) WithContext(f func(), drawing bool) Status {
	if f == nil {
		return Failure
	}
	x.EnterContext(drawing)
	f()
	x.LeaveContext(drawing)
	return Success
}

// SetEventFunc - set callback handler
func (x *View) SetEventFunc(callback EventFunc) {
	entry := x.entry()
	entry.callback = callback
}

// PostRedisplay request a redisplay on event dispatch.
func (x *View) PostRedisplay() {
	C.puglPostRedisplay(x.view)
}

// Free the View
func (x *View) Free() {
	if x == nil || x.view == nil {
		return
	}
	C.puglFreeView(x.view)
	x.view = nil
}

// GetWorld - get the world
func (x *View) GetWorld() *World {
	if x == nil {
		return nil
	}
	return x.world
}

// HasFocus - Return true iff `view` has the input focus.
func (x *View) HasFocus() bool {
	return bool(C.puglHasFocus(x.view))
}

// GrabFocus - grab the input focus.
func (x *View) GrabFocus() Status {
	return Status(C.puglGrabFocus(x.view))
}

// GetClipboard - get the clipboard
func (x *View) GetClipboard() {

}

// SetClipboard - set the clipboard
func (x *View) SetClipboard() Status {
	return Success
}

// RequestAttention from the user.
//
// This hints to the system that the window or application requires attention
// from the user.  The exact effect depends on the platform, but is usually
// something like flashing a task bar entry.
func (x *View) RequestAttention() Status {
	return Status(C.puglRequestAttention(x.view))
}

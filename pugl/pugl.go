package pugl

/*
#cgo CFLAGS: -Isrc -Wno-deprecated-declarations
#cgo LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include "pugl/pugl.h"
*/
import "C"
import (
	"unsafe"
)

// Status returned status code.
type Status C.PuglStatus

// Status enum
const (
	Success          Status = C.PUGL_SUCCESS
	ErrCreateWindow  Status = C.PUGL_ERR_CREATE_WINDOW
	ErrSetFormat     Status = C.PUGL_ERR_SET_FORMAT
	ErrCreateContext Status = C.PUGL_ERR_CREATE_CONTEXT
)

// WindowHint - Window hint
type WindowHint C.PuglWindowHint

const (
	// UseCompatProfile - Use compatible (not core) OpenGL profile
	UseCompatProfile WindowHint = C.PUGL_USE_COMPAT_PROFILE
	// ContextVersionMajor - OpenGL context major version
	ContextVersionMajor WindowHint = C.PUGL_CONTEXT_VERSION_MAJOR
	// ContextVersionMinor - OpenGL context minor version
	ContextVersionMinor WindowHint = C.PUGL_CONTEXT_VERSION_MINOR
	// RedBits - Number of bits for red channel
	RedBits WindowHint = C.PUGL_RED_BITS
	// GreenBits - Number of bits for green channel
	GreenBits WindowHint = C.PUGL_GREEN_BITS
	// BlueBits - Number of bits for blue channel
	BlueBits WindowHint = C.PUGL_BLUE_BITS
	// AlphaBits - Number of bits for alpha channel
	AlphaBits WindowHint = C.PUGL_ALPHA_BITS
	// DepthBits - Number of bits for depth buffer
	DepthBits WindowHint = C.PUGL_DEPTH_BITS
	// StencilBits - Number of bits for stencil buffer
	StencilBits WindowHint = C.PUGL_STENCIL_BITS
	// Samples - Number of samples per pixel (AA)
	Samples WindowHint = C.PUGL_SAMPLES
	// DoubleBuffer - True if double buffering should be used
	DoubleBuffer WindowHint = C.PUGL_DOUBLE_BUFFER
	// Resizable - True if window should be resizable
	Resizable WindowHint = C.PUGL_RESIZABLE
	// IgnoreKeyRepeat - True if key repeat events are ignored
	IgnoreKeyRepeat WindowHint = C.PUGL_IGNORE_KEY_REPEAT
)

// WindowHintValue - Special window hint value.
type WindowHintValue C.PuglWindowHintValue

const (
	// DontCare - Use best available value */
	DontCare WindowHintValue = C.PUGL_DONT_CARE
	// False - Explicitly false */
	False WindowHintValue = C.PUGL_FALSE
	// True - Explicitly true */
	True WindowHintValue = C.PUGL_TRUE
)

// Mod - Keyboard modifier flags.
type Mod C.PuglMod

// Mod values
const (
	ModShift Mod = C.PUGL_MOD_SHIFT // Shift
	ModCtrl  Mod = C.PUGL_MOD_CTRL  // Control
	ModAlt   Mod = C.PUGL_MOD_ALT   // Alt
	ModSuper Mod = C.PUGL_MOD_SUPER // Super
)

// Key - Special keyboard keys.
//
// All keys, special or not, are expressed as a Unicode code point.  This
// enumeration defines constants for special keys that do not have a standard
// code point, and some convenience constants for control characters.
//
// Keys that do not have a standard code point use values in the Private Use
// Area in the Basic Multilingual Plane (U+E000 to U+F8FF).  Applications must
// take care to not interpret these values beyond key detection, the mapping
// used here is arbitrary and specific to Pugl.
type Key C.PuglKey

// Key values
const (
	// ASCII control codes
	KeyBackspace Key = C.PUGL_KEY_BACKSPACE
	KeyEscape    Key = C.PUGL_KEY_ESCAPE
	KeyDelete    Key = C.PUGL_KEY_DELETE

	// Unicode Private Use Area
	KeyF1          Key = C.PUGL_KEY_F1
	KeyF2          Key = C.PUGL_KEY_F2
	KeyF3          Key = C.PUGL_KEY_F3
	KeyF4          Key = C.PUGL_KEY_F4
	KeyF5          Key = C.PUGL_KEY_F5
	KeyF6          Key = C.PUGL_KEY_F6
	KeyF7          Key = C.PUGL_KEY_F7
	KeyF8          Key = C.PUGL_KEY_F8
	KeyF9          Key = C.PUGL_KEY_F9
	KeyF10         Key = C.PUGL_KEY_F10
	KeyF11         Key = C.PUGL_KEY_F11
	KeyF12         Key = C.PUGL_KEY_F12
	KeyLeft        Key = C.PUGL_KEY_LEFT
	KeyUp          Key = C.PUGL_KEY_UP
	KeyRight       Key = C.PUGL_KEY_RIGHT
	KeyDown        Key = C.PUGL_KEY_DOWN
	KeyPageUp      Key = C.PUGL_KEY_PAGE_UP
	KeyPageDown    Key = C.PUGL_KEY_PAGE_DOWN
	KeyHome        Key = C.PUGL_KEY_HOME
	KeyEnd         Key = C.PUGL_KEY_END
	KeyInsert      Key = C.PUGL_KEY_INSERT
	KeyShift       Key = C.PUGL_KEY_SHIFT
	KeyShiftL      Key = C.PUGL_KEY_SHIFT_L
	KeyShiftR      Key = C.PUGL_KEY_SHIFT_R
	KeyCtrl        Key = C.PUGL_KEY_CTRL
	KeyCtrlL       Key = C.PUGL_KEY_CTRL_L
	KeyCtrlR       Key = C.PUGL_KEY_CTRL_R
	KeyAlt         Key = C.PUGL_KEY_ALT
	KeyAltL        Key = C.PUGL_KEY_ALT_L
	KeyAltR        Key = C.PUGL_KEY_ALT_R
	KeySuper       Key = C.PUGL_KEY_SUPER
	KeySuperL      Key = C.PUGL_KEY_SUPER_L
	KeySuperR      Key = C.PUGL_KEY_SUPER_R
	KeyMenu        Key = C.PUGL_KEY_MENU
	KeyCapsLock    Key = C.PUGL_KEY_CAPS_LOCK
	KeyScrollLock  Key = C.PUGL_KEY_SCROLL_LOCK
	KeyNumLock     Key = C.PUGL_KEY_NUM_LOCK
	KeyPrintScreen Key = C.PUGL_KEY_PRINT_SCREEN
	KeyPause       Key = C.PUGL_KEY_PAUSE
)

// EventType - The type of an Event.
type EventType C.PuglEventType

// EventType values
const (
	Nothing       EventType = C.PUGL_NOTHING        /**< No event */
	ButtonPress   EventType = C.PUGL_BUTTON_PRESS   /**< Mouse button press */
	ButtonRelease EventType = C.PUGL_BUTTON_RELEASE /**< Mouse button release */
	Configure     EventType = C.PUGL_CONFIGURE      /**< View moved and/or resized */
	Expose        EventType = C.PUGL_EXPOSE         /**< View exposed redraw required */
	Close         EventType = C.PUGL_CLOSE          /**< Close view */
	KeyPress      EventType = C.PUGL_KEY_PRESS      /**< Key press */
	KeyRelease    EventType = C.PUGL_KEY_RELEASE    /**< Key release */
	Text          EventType = C.PUGL_TEXT           /**< Character entry */
	EnterNotify   EventType = C.PUGL_ENTER_NOTIFY   /**< Pointer entered view */
	LeaveNotify   EventType = C.PUGL_LEAVE_NOTIFY   /**< Pointer left view */
	MotionNotify  EventType = C.PUGL_MOTION_NOTIFY  /**< Pointer motion */
	Scroll        EventType = C.PUGL_SCROLL         /**< Scroll */
	FocusIn       EventType = C.PUGL_FOCUS_IN       /**< Keyboard focus entered view */
	FocusOut      EventType = C.PUGL_FOCUS_OUT      /**< Keyboard focus left view */
)

// EventFlag enum
type EventFlag C.PuglEventFlag

// EventFlag values
const (
	IsSendEvent EventFlag = C.PUGL_IS_SEND_EVENT
)

// CrossingMode - Reason for a PuglEventCrossing.
type CrossingMode C.PuglCrossingMode

// Crossing mode values
const (
	CrossingNormal CrossingMode = C.PUGL_CROSSING_NORMAL // Crossing due to pointer motion
	CrossingGrab   CrossingMode = C.PUGL_CROSSING_GRAB   // Crossing due to a grab
	CrossingUngrab CrossingMode = C.PUGL_CROSSING_UNGRAB // Crossing due to a grab release
)

// Backend interface
type Backend C.PuglBackend

// EventButton a button event
type EventButton struct {
	Time   float64
	X      float64
	Y      float64
	RootX  float64
	RootY  float64
	State  uint32
	Button uint32
}

// EventConfigure - configure event (resized)
type EventConfigure struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// EventExpose - draw event
type EventExpose struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
	Count  int
}

// EventKey - keypress event
type EventKey struct {
	Time    float64
	X       float64
	Y       float64
	RootX   float64
	RootY   float64
	State   uint32
	KeyCode uint32
	Key     uint32
}

// EventText - text event
type EventText struct {
	Time       float64
	X          float64
	Y          float64
	RootX      float64
	RootY      float64
	State      uint32
	KeyCode    uint32
	Character  uint32
	StringUTF8 [8]byte
}

// EventCrossing crossing event
type EventCrossing struct {
	Time  float64
	X     float64
	Y     float64
	RootX float64
	RootY float64
	State uint32
	Mode  CrossingMode
}

// EventMotion a motion event
type EventMotion struct {
	Time   float64
	X      float64
	Y      float64
	RootX  float64
	RootY  float64
	State  uint32
	IsHint bool
	Focus  bool
}

// EventScroll - scroll event
type EventScroll struct {
	Time  float64
	X     float64
	Y     float64
	RootX float64
	RootY float64
	State uint32
	DX    float64
	DY    float64
}

// EventFocus focus grab/ungrab event
type EventFocus struct {
	Grab bool
}

// Event abstraction
type Event struct {
	Type      EventType
	Flags     uint32
	Button    EventButton
	Configure EventConfigure
	Expose    EventExpose
	Key       EventKey
	Text      EventText
	Crossing  EventCrossing
	Motion    EventMotion
	Scroll    EventScroll
	Focus     EventFocus

	// Any EventAny
	// Close     EventClose

	null C.PuglEvent
	raw  [C.sizeof_PuglEvent]byte
}

// EventFunc - A function called when an event occurs.
type EventFunc func(view *View, event *Event)

// EventHandler interface
type EventHandler interface {
	Process(*Event)
}

// View a pugl View wrapper
type View struct {
	view *C.PuglView
}

// Init - Create a Pugl view.
//
// To create a window, call the various puglInit* functions as necessary, then
// call pugl.CreateWindow().
//
// @return A newly created view.
func Init() *View {
	view := new(View)
	view.view = C.puglInit(nil, nil)
	if view.view == nil {
		return nil
	}
	view.setInternalHandler()
	return view
}

// InitWindowHint - Set the window class name before creating a window.
func (x *View) InitWindowHint(hint WindowHint, value int) {
	C.puglInitWindowHint(x.view, (C.PuglWindowHint)(hint), C.int(value))
}

// InitWindowClass - Set the window class name before creating a window.
func (x *View) InitWindowClass(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.puglInitWindowClass(x.view, cname)
}

// InitWindowParent - Set the parent window before creating a window (for embedding).
func (x *View) InitWindowParent(parent uintptr) {
	C.puglInitWindowParent(x.view, C.intptr_t(parent))
}

// InitWindowSize - Set the parent window before creating a window (for embedding).
func (x *View) InitWindowSize(width, height int) {
	C.puglInitWindowSize(x.view, C.int(width), C.int(height))
}

// InitWindowMinSize - Set the minimum window size before creating a window.
func (x *View) InitWindowMinSize(width, height int) {
	C.puglInitWindowMinSize(x.view, C.int(width), C.int(height))
}

// InitWindowAspectRatio - Set the window aspect ratio range before creating a window.
//
// The x and y values here represent a ratio of width to height.  To set a
// fixed aspect ratio, set the minimum and maximum values to the same ratio.
//
// Note that setting different minimum and maximum constraints does not
// currenty work on MacOS (the minimum is used), so only setting a fixed aspect
// ratio works properly across all platforms.
func (x *View) InitWindowAspectRatio(minX, minY, maxX, maxY int) {
	C.puglInitWindowAspectRatio(x.view,
		C.int(minX), C.int(minY), C.int(maxX), C.int(maxY))
}

// InitTransientFor - Set transient parent before creating a window.
//
// On X11, parent must be a Window.
// On OSX, parent must be an NSView*.
func (x *View) InitTransientFor(parent uintptr) {
	C.puglInitTransientFor(x.view, C.uintptr_t(parent))
}

// InitBackend - init the backend
func (x *View) InitBackend(backend *Backend) {
	C.puglInitBackend(x.view, (*C.PuglBackend)(backend))
}

// CreateWindow makes a window with the settings given by the various pugl.Init functions.
// @return 1 (pugl does not currently support multiple windows).
func (x *View) CreateWindow(title string) int {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	return int(C.puglCreateWindow(x.view, ctitle))
}

// ShowWindow shows the current window.
func (x *View) ShowWindow() {
	C.puglShowWindow(x.view)
}

// HideWindow - hide the current window
func (x *View) HideWindow() {
	C.puglHideWindow(x.view)
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

/* ----- */

//Handle - opaque user data
type Handle unsafe.Pointer

// /**
//    Set the handle to be passed to all callbacks.

//    This is generally a pointer to a struct which contains all necessary state.
//    Everything needed in callbacks should be here, not in static variables.
// */
// PUGL_API void
// puglSetHandle(PuglView* view, PuglHandle handle);

// /**
//    Get the handle to be passed to all callbacks.
// */
// PUGL_API PuglHandle
// puglGetHandle(PuglView* view);

/* ----- */

// GetVisible - true if visible
func (x *View) GetVisible() bool {
	return bool(C.puglGetVisible(x.view))
}

// GetSize - Get the current size of the view
func (x *View) GetSize() (int, int) {
	var w, h C.int
	C.puglGetSize(x.view, &w, &h)
	return int(w), int(h)
}

// /**
//    Get the drawing context.

//    The context is only guaranteed to be available during an expose.

//    For OpenGL backends, this is unused and returns NULL.
//    For Cairo backends, this returns a pointer to a `cairo_t`.
// */
// PUGL_API void*
// puglGetContext(PuglView* view);

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

// SetEventFunc - set callback handler
func (x *View) SetEventFunc(callback EventFunc) {
	entry := x.entry()
	entry.callback = callback
}

// SetEventHandler - set callback handler
func (x *View) SetEventHandler(handler EventHandler) {
	entry := x.entry()
	entry.handler = handler
}

// WaitForEvent - wait for next event
func (x *View) WaitForEvent() Status {
	return Status(C.puglWaitForEvent(x.view))
}

// ProcessEvents - process events
func (x *View) ProcessEvents() Status {
	return Status(C.puglProcessEvents(x.view))
}

// PostRedisplay request a redisplay on the next call to puglProcessEvents().
func (x *View) PostRedisplay() {
	C.puglPostRedisplay(x.view)
}

// Destroy the view
func (x *View) Destroy() {
	if x.view == nil {
		return
	}
	C.puglDestroy(x.view)
	x.view = nil
}

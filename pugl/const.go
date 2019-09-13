// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 12 Sep 2019 01:42:04 EDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package pugl

/*
#include "pugl/pugl.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Status as declared in pugl/pugl.h:74
type Status int32

// Status enumeration from pugl/pugl.h:74
const (
	Success             Status = C.PUGL_SUCCESS
	Failure             Status = C.PUGL_FAILURE
	UnknownError        Status = C.PUGL_UNKNOWN_ERROR
	BadBackend          Status = C.PUGL_BAD_BACKEND
	BackendFailed       Status = C.PUGL_BACKEND_FAILED
	RegistrationFailed  Status = C.PUGL_REGISTRATION_FAILED
	CreateWindowFailed  Status = C.PUGL_CREATE_WINDOW_FAILED
	SetFormatFailed     Status = C.PUGL_SET_FORMAT_FAILED
	CreateContextFailed Status = C.PUGL_CREATE_CONTEXT_FAILED
	UnsupportedType     Status = C.PUGL_UNSUPPORTED_TYPE
)

// ViewHint as declared in pugl/pugl.h:96
type ViewHint int32

// ViewHint enumeration from pugl/pugl.h:96
const (
	UseCompatProfile    ViewHint = C.PUGL_USE_COMPAT_PROFILE
	ContextVersionMajor ViewHint = C.PUGL_CONTEXT_VERSION_MAJOR
	ContextVersionMinor ViewHint = C.PUGL_CONTEXT_VERSION_MINOR
	RedBits             ViewHint = C.PUGL_RED_BITS
	GreenBits           ViewHint = C.PUGL_GREEN_BITS
	BlueBits            ViewHint = C.PUGL_BLUE_BITS
	AlphaBits           ViewHint = C.PUGL_ALPHA_BITS
	DepthBits           ViewHint = C.PUGL_DEPTH_BITS
	StencilBits         ViewHint = C.PUGL_STENCIL_BITS
	Samples             ViewHint = C.PUGL_SAMPLES
	DoubleBuffer        ViewHint = C.PUGL_DOUBLE_BUFFER
	SwapInterval        ViewHint = C.PUGL_SWAP_INTERVAL
	Resizable           ViewHint = C.PUGL_RESIZABLE
	IgnoreKeyRepeat     ViewHint = C.PUGL_IGNORE_KEY_REPEAT
	NumWindowHints      ViewHint = C.PUGL_NUM_WINDOW_HINTS
)

// ViewHintValue as declared in pugl/pugl.h:105
type ViewHintValue int32

// ViewHintValue enumeration from pugl/pugl.h:105
const (
	DontCare ViewHintValue = C.PUGL_DONT_CARE
	False    ViewHintValue = C.PUGL_FALSE
	True     ViewHintValue = C.PUGL_TRUE
)

// Mod as declared in pugl/pugl.h:128
type Mod int32

// Mod enumeration from pugl/pugl.h:128
const (
	ModShift Mod = C.PUGL_MOD_SHIFT
	ModCtrl  Mod = C.PUGL_MOD_CTRL
	ModAlt   Mod = C.PUGL_MOD_ALT
	ModSuper Mod = C.PUGL_MOD_SUPER
)

// Key as declared in pugl/pugl.h:188
type Key int32

// Key enumeration from pugl/pugl.h:188
const (
	KeyBackspace   Key = C.PUGL_KEY_BACKSPACE
	KeyEscape      Key = C.PUGL_KEY_ESCAPE
	KeyDelete      Key = C.PUGL_KEY_DELETE
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

// EventType as declared in pugl/pugl.h:209
type EventType int32

// EventType enumeration from pugl/pugl.h:209
const (
	Nothing       EventType = C.PUGL_NOTHING
	ButtonPress   EventType = C.PUGL_BUTTON_PRESS
	ButtonRelease EventType = C.PUGL_BUTTON_RELEASE
	Configure     EventType = C.PUGL_CONFIGURE
	Expose        EventType = C.PUGL_EXPOSE
	Close         EventType = C.PUGL_CLOSE
	KeyPress      EventType = C.PUGL_KEY_PRESS
	KeyRelease    EventType = C.PUGL_KEY_RELEASE
	Text          EventType = C.PUGL_TEXT
	EnterNotify   EventType = C.PUGL_ENTER_NOTIFY
	LeaveNotify   EventType = C.PUGL_LEAVE_NOTIFY
	MotionNotify  EventType = C.PUGL_MOTION_NOTIFY
	Scroll        EventType = C.PUGL_SCROLL
	FocusIn       EventType = C.PUGL_FOCUS_IN
	FocusOut      EventType = C.PUGL_FOCUS_OUT
)

// EventFlag as declared in pugl/pugl.h:213
type EventFlag int32

// EventFlag enumeration from pugl/pugl.h:213
const (
	IsSendEvent EventFlag = C.PUGL_IS_SEND_EVENT
)

// CrossingMode as declared in pugl/pugl.h:222
type CrossingMode int32

// CrossingMode enumeration from pugl/pugl.h:222
const (
	CrossingNormal CrossingMode = C.PUGL_CROSSING_NORMAL
	CrossingGrab   CrossingMode = C.PUGL_CROSSING_GRAB
	CrossingUngrab CrossingMode = C.PUGL_CROSSING_UNGRAB
)
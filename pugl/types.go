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

// EventAny as declared in pugl/pugl.h:230
type EventAny struct {
	Type  EventType
	Flags uint32
}

// EventButton as declared in pugl/pugl.h:247
type EventButton struct {
	Type   EventType
	Flags  uint32
	Time   float64
	X      float64
	Y      float64
	XRoot  float64
	YRoot  float64
	State  uint32
	Button uint32
}

// EventConfigure as declared in pugl/pugl.h:259
type EventConfigure struct {
	Type   EventType
	Flags  uint32
	X      float64
	Y      float64
	Width  float64
	Height float64
	Scale  float64
}

// EventExpose as declared in pugl/pugl.h:272
type EventExpose struct {
	Type   EventType
	Flags  uint32
	X      float64
	Y      float64
	Width  float64
	Height float64
	Count  int32
}

// EventClose as declared in pugl/pugl.h:280
type EventClose struct {
	Type  EventType
	Flags uint32
}

// EventKey as declared in pugl/pugl.h:305
type EventKey struct {
	Type    EventType
	Flags   uint32
	Time    float64
	X       float64
	Y       float64
	XRoot   float64
	YRoot   float64
	State   uint32
	Keycode uint32
	Key     uint32
}

// EventText as declared in pugl/pugl.h:325
type EventText struct {
	Type      EventType
	Flags     uint32
	Time      float64
	X         float64
	Y         float64
	XRoot     float64
	YRoot     float64
	State     uint32
	Keycode   uint32
	Character uint32
	String    [8]byte
}

// EventCrossing as declared in pugl/pugl.h:340
type EventCrossing struct {
	Type  EventType
	Flags uint32
	Time  float64
	X     float64
	Y     float64
	XRoot float64
	YRoot float64
	State uint32
	Mode  CrossingMode
}

// Rect - frame of a view
type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// EventMotion as declared in pugl/pugl.h:356
type EventMotion struct {
	Type   EventType
	Flags  uint32
	Time   float64
	X      float64
	Y      float64
	XRoot  float64
	YRoot  float64
	State  uint32
	IsHint bool
	Focus  bool
}

// EventScroll as declared in pugl/pugl.h:378
type EventScroll struct {
	Type  EventType
	Flags uint32
	Time  float64
	X     float64
	Y     float64
	XRoot float64
	YRoot float64
	State uint32
	Dx    float64
	Dy    float64
}

// EventFocus as declared in pugl/pugl.h:387
type EventFocus struct {
	Type  EventType
	Flags uint32
	Grab  bool
}
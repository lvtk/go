package pugl

/*
#include "pugl/pugl.h"
#include "pugl/pugl_gl_backend.h"
#include <stdlib.h>
#include <stdio.h>
extern void event_handler_go(PuglView* view, PuglEvent* event);

static void event_handler(PuglView* view, const PuglEvent* event) {
	event_handler_go(view, (PuglEvent*)event);
}

static void set_event_handler(PuglView* view) {
	puglSetEventFunc(view, event_handler);
}

static PuglEventAny* any (const PuglEvent* event) {
	return (PuglEventAny*) event;
}

static PuglEventButton* button (const PuglEvent* event) {
	return (PuglEventButton*) event;
}

static PuglEventConfigure* configure (const PuglEvent* event) {
	return (PuglEventConfigure*) event;
}

static PuglEventExpose* expose (const PuglEvent* event) {
	return (PuglEventExpose*) event;
}

static PuglEventClose* close (const PuglEvent* event) {
	return (PuglEventClose*) event;
}

static PuglEventKey* key (const PuglEvent* event) {
	return (PuglEventKey*) event;
}

static PuglEventText* text (const PuglEvent* event) {
	return (PuglEventText*) event;
}

static PuglEventCrossing* crossing (const PuglEvent* event) {
	return (PuglEventCrossing*) event;
}

static PuglEventMotion* motion (const PuglEvent* event) {
	return (PuglEventMotion*) event;
}

static PuglEventScroll* scroll (const PuglEvent* event) {
	return (PuglEventScroll*) event;
}

static PuglEventFocus* focus (const PuglEvent* event) {
	return (PuglEventFocus*) event;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type evMapEntry struct {
	view     *View
	callback EventFunc
	handler  EventHandler
}

var evMap = make(map[*C.PuglView]*evMapEntry)
var ev *Event

//export event_handler_go
func event_handler_go(view *C.PuglView, event *C.PuglEvent) {
	entry := evMap[view]
	if entry == nil {
		fmt.Println("Event on nil entry")
		return
	}

	if entry.view == nil {
		fmt.Println("Event on nil pugl.View")
		return
	}

	if entry.view.view != view {
		fmt.Println("Event on mismatching C-Type view")
		return
	}

	if ev == nil {
		ev = new(Event)
	}
	ev.update(event)
	if entry.callback != nil {
		entry.callback(entry.view, ev)
	}
	if entry.handler != nil {
		entry.handler.Process(ev)
	}

	ev.reset()
}

func (x *Event) reset() {
	x.Type = Nothing
	x.raw = x.null
}

func (x *Event) update(cev *C.PuglEvent) {
	x.raw = *cev
	x.Type = EventType(x.raw[0])
	if x.Type == ButtonPress || x.Type == ButtonRelease {
		bev := C.button(cev)
		x.Button.Time = float64(bev.time)
		x.Button.X = float64(bev.x)
		x.Button.Y = float64(bev.y)
		x.Button.RootX = float64(bev.x_root)
		x.Button.RootY = float64(bev.y_root)
		x.Button.State = uint32(bev.state)
		x.Button.Button = uint32(bev.button)
	} else if x.Type == Configure {
		bev := C.configure(cev)
		x.Configure.X = float64(bev.x)
		x.Configure.Y = float64(bev.y)
		x.Configure.Width = float64(bev.width)
		x.Configure.Height = float64(bev.height)
	} else if x.Type == Expose {
		bev := C.expose(cev)
		x.Expose.X = float64(bev.x)
		x.Expose.Y = float64(bev.y)
		x.Expose.Width = float64(bev.width)
		x.Expose.Height = float64(bev.height)
		x.Expose.Count = int(bev.count)
	} else if x.Type == KeyPress || x.Type == KeyRelease {
		bev := C.key(cev)
		x.Key.Time = float64(bev.time)
		x.Key.X = float64(bev.x)
		x.Key.Y = float64(bev.y)
		x.Key.RootX = float64(bev.x_root)
		x.Key.RootY = float64(bev.y_root)
		x.Key.State = uint32(bev.state)
		x.Key.KeyCode = uint32(bev.keycode)
		x.Key.Key = uint32(bev.key)
	} else if x.Type == Text {
		bev := C.text(cev)
		x.Text.Time = float64(bev.time)
		x.Text.X = float64(bev.x)
		x.Text.Y = float64(bev.y)
		x.Text.RootX = float64(bev.x_root)
		x.Text.RootY = float64(bev.y_root)
		x.Text.State = uint32(bev.state)
		x.Text.KeyCode = uint32(bev.keycode)
		x.Text.Character = uint32(bev.character)
		// x.Text.StringUTF8 = C.GoString(bev.string)
	} else if x.Type == EnterNotify || x.Type == LeaveNotify {
		bev := C.crossing(cev)
		x.Crossing.Time = float64(bev.time)
		x.Crossing.X = float64(bev.x)
		x.Crossing.Y = float64(bev.y)
		x.Crossing.RootX = float64(bev.x_root)
		x.Crossing.RootY = float64(bev.y_root)
		x.Crossing.State = uint32(bev.state)
		x.Crossing.Mode = CrossingMode(bev.mode)
	} else if x.Type == MotionNotify {
		bev := C.motion(cev)
		x.Motion.Time = float64(bev.time)
		x.Motion.X = float64(bev.x)
		x.Motion.Y = float64(bev.y)
		x.Motion.RootX = float64(bev.x_root)
		x.Motion.RootY = float64(bev.y_root)
		x.Motion.State = uint32(bev.state)
		x.Motion.IsHint = bool(bev.is_hint)
		x.Motion.Focus = bool(bev.focus)
	} else if x.Type == Scroll {
		bev := C.scroll(cev)
		x.Scroll.Time = float64(bev.time)
		x.Scroll.X = float64(bev.x)
		x.Scroll.Y = float64(bev.y)
		x.Scroll.RootX = float64(bev.x_root)
		x.Scroll.RootY = float64(bev.y_root)
		x.Scroll.State = uint32(bev.state)
		x.Scroll.DX = float64(bev.dx)
		x.Scroll.DY = float64(bev.dy)
	} else if x.Type == FocusIn || x.Type == FocusOut {
		bev := C.focus(cev)
		x.Focus.Grab = bool(bev.grab)
	}
}

func (x *View) entry() *evMapEntry {
	e := evMap[x.view]
	if e == nil {
		e = new(evMapEntry)
		e.view = x
		evMap[x.view] = e
	}
	return e
}

func (x *View) setInternalHandler() {
	C.puglSetHandle(x.view, (C.PuglHandle)(unsafe.Pointer(x.view)))
	C.set_event_handler(x.view)
}

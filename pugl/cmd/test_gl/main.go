package main

/*
#cgo CFLAGS: -I../../src
#include "test_utils.h"
#include <stdlib.h>
static float valuef(const float* f, int i) { return f[i]; }
*/
import "C"
import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/lvtk/go/pugl"
	pgl "github.com/lvtk/go/pugl/gl"
)

type testApp struct {
	world      *pugl.World
	view       *pugl.View
	xAngle     float64
	yAngle     float64
	mouseX     float64
	mouseY     float64
	quit       bool
	projection *C.float
}

func (x *testApp) free() {
	x.view.Free()
	x.world.Free()
	if x.projection != nil {
		C.free(unsafe.Pointer(x.projection))
	}
}

func (x *testApp) configure(width, height int32) {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Viewport(0, 0, width, height)

	if x.projection == nil {
		x.projection = (*C.float)(C.malloc(C.sizeof_float * 16))
	}
	C.perspective(x.projection, 1.8,
		C.float(width)/C.float(height), 1.0, 100.0)

	var gproj [16]float32
	for i := 0; i < 16; i++ {
		gproj[i] = float32(C.valuef(x.projection, C.int(i)))
	}
	gl.LoadMatrixf(&gproj[0])
}

func (x *testApp) expose() {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	dist := 10.0
	gl.Translatef(0.0, 0.0, float32(dist*-1.0))
	gl.Rotatef(float32(x.xAngle), 0.0, 1.0, 0.0)
	gl.Rotatef(float32(x.yAngle), 1.0, 0.0, 0.0)

	// const float bg = app->mouseEntered ? 0.2f : 0.0f;
	bg := float32(0.1)
	gl.ClearColor(bg, bg, bg, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.EnableClientState(gl.COLOR_ARRAY)

	gl.VertexPointer(3, gl.FLOAT, 0, unsafe.Pointer(&gcubeVerticies[0]))
	gl.ColorPointer(3, gl.FLOAT, 0, unsafe.Pointer(&gcubeVerticies[0]))
	gl.DrawArrays(gl.TRIANGLES, 0, 12*3)

	gl.DisableClientState(gl.VERTEX_ARRAY)
	gl.DisableClientState(gl.COLOR_ARRAY)
}

func (x *testApp) onEvent(view *pugl.View, e *pugl.Event) {
	if e.Type == pugl.ButtonPress || e.Type == pugl.ButtonRelease {
		fmt.Println("button event:", e.Button)
	} else {
		switch e.Type {

		case pugl.Configure:
			c := &e.Configure
			x.configure(int32(c.Width*c.Scale),
				int32(c.Height*c.Scale))
		case pugl.Expose:
			x.expose()
		case pugl.Close:
			x.quit = true
		case pugl.MotionNotify:
			x.xAngle = math.Mod(x.xAngle-(e.Motion.X-x.mouseX), 360.0)
			x.yAngle = math.Mod(x.yAngle+(e.Motion.Y-x.mouseY), 360.0)
			x.mouseX = e.Motion.X
			x.mouseY = e.Motion.Y
			x.view.PostRedisplay()
		}
	}
}

func main() {
	app := new(testApp)
	defer app.free()
	app.quit = false
	app.world = pugl.NewWorld()
	app.view = app.world.NewView()

	if app.view == nil {
		fmt.Println("couldn't init pugl")
		return
	}

	app.world.SetClassName("PuglTest")
	app.view.SetFrame(pugl.Rect{0, 0, 512, 512})
	app.view.SetMinSize(256, 256)
	app.view.SetViewHint(pugl.ContextVersionMajor, 2)
	app.view.SetViewHint(pugl.ContextVersionMinor, 1)
	app.view.SetBackend(pgl.Backend())

	app.view.SetViewHint(pugl.Resizable, 1)
	app.view.SetViewHint(pugl.Samples, 4)

	app.view.SetEventFunc(app.onEvent)

	if 0 != app.view.CreateWindow("Test Window") {
		os.Exit(1)
	}

	app.view.WithContext(func() {
		gl.Init()
	}, false)

	app.view.ShowWindow()

	for {
		app.world.PollEvents(0)
		app.world.DispatchEvents()
		if app.quit {
			break
		}
	}
}

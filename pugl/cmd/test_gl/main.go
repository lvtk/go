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

func configure(width, height int32) {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Viewport(0, 0, width, height)

	projection := (*C.float)(C.malloc(C.sizeof_float * 16))
	C.perspective(projection, 1.8,
		C.float(width)/C.float(height), 1.0, 100.0)
	defer C.free(unsafe.Pointer(projection))
	var gproj [16]float32
	for i := 0; i < 16; i++ {
		gproj[i] = float32(C.valuef(projection, C.int(i)))
	}
	gl.LoadMatrixf(&gproj[0])
}

var gcubeVerticies = []float32{
	-1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,

	1.0, 1.0, -1.0,
	-1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,

	1.0, -1.0, 1.0,
	-1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,

	1.0, 1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, -1.0, -1.0,

	-1.0, -1.0, -1.0,
	-1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0,

	1.0, -1.0, 1.0,
	-1.0, -1.0, 1.0,
	-1.0, -1.0, -1.0,

	-1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0,
	1.0, -1.0, 1.0,

	1.0, 1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, 1.0, -1.0,

	1.0, -1.0, -1.0,
	1.0, 1.0, 1.0,
	1.0, -1.0, 1.0,

	1.0, 1.0, 1.0,
	1.0, 1.0, -1.0,
	-1.0, 1.0, -1.0,

	1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,

	1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0,
	1.0, -1.0, 1.0}

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

type testApp struct {
	view   *pugl.View
	xAngle float64
	yAngle float64
	mouseX float64
	mouseY float64
	quit   bool
}

func (x *testApp) Process(e *pugl.Event) {
	if e.Type == pugl.ButtonPress || e.Type == pugl.ButtonRelease {
		fmt.Println("button event:", e.Button)
	} else {
		switch e.Type {

		case pugl.Configure:
			configure(int32(e.Configure.Width),
				int32(e.Configure.Height))
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
	world := pugl.NewWorld()
	view := world.NewView()
	defer view.Free()
	defer world.Free()

	if view == nil {
		fmt.Println("couldn't init pugl")
		return
	}

	world.SetClassName("PuglTest")
	view.SetFrame(pugl.Rect{512, 512})
	view.SetMinSize(256, 256)
	view.SetViewHint(pugl.ContextVersionMajor, 2)
	view.SetViewHint(pugl.ContextVersionMinor, 1)
	view.SetBackend(pgl.Backend())

	view.SetViewHint(pugl.Resizable, 1)
	view.SetViewHint(pugl.Samples, 4)

	app := new(testApp)
	app.quit = false
	app.view = view

	view.SetEventHandler(app)
	if 0 != view.CreateWindow("Test Window") {
		os.Exit(1)
	}

	view.WithContext(func() {
		gl.Init()
	}, false)

	view.ShowWindow()

	for {
		world.PollEvents(0)
		world.DispatchEvents()
		if app.quit {
			break
		}
	}
}

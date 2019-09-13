package gl

/*
#cgo CFLAGS: -I../src
#cgo darwin LDFLAGS: -framework Cocoa -framework OpenGL
#include "pugl/pugl_gl_backend.h"
*/
import "C"
import "github.com/lvtk/go/pugl"

// Backend creates a GL backend
func Backend() *pugl.Backend {
	return (*pugl.Backend)(C.puglGlBackend())
}

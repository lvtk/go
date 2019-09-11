package pugl_test

import (
	"testing"

	"github.com/lvtk/go/pugl"
)

func TestPugl(t *testing.T) {
	view := pugl.Init()
	if view == nil {
		t.Fatalf("couldn't create view")
	}
}

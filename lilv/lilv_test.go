package lilv

import (
	"testing"
)

func TestNewWorldAndFree(t *testing.T) {
	world := NewWorld()
	world.Free()
	if world.world != nil {
		t.Errorf("world not freed")
	}
}

func TestNewWorldURI(t *testing.T) {
	world := NewWorld()
	defer world.Free()
	uri := world.NewURI("http://google.com")
	defer uri.Free()
	if !uri.IsURI() {
		t.Fatalf("not a URI")
	}
	if uri.String() != "http://google.com" {
		t.Fatalf("URI not consistent: %s != http://google.com", uri.String())
	}
}

func TestNodeGetTurtleToken(t *testing.T) {
	world := NewWorld()
	defer world.Free()
	uri := world.NewURI("http://google.com")
	defer uri.Free()
	uri.GetTurtleToken()
}

func TestNodeBool(t *testing.T) {
	world := NewWorld()
	defer world.Free()

	bval := world.NewBool(false)
	defer bval.Free()
	if !bval.IsBool() {
		t.Errorf("not a bool: %s", bval.String())
	}
	if bval.Bool() != false {
		t.Errorf("bool val invalid: %s", bval.String())
	}
}

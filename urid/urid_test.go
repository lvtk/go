package urid_test

import (
	"testing"

	"github.com/lvtk/go/urid"
)

func TestMapping(t *testing.T) {
	var d = new(urid.Directory)

	urid := d.Map("http://test.org")
	if urid != 1 {
		t.Errorf("zero not allowd for mapping")
	}
	if "http://test.org" != d.Unmap(urid) {
		t.Errorf("not mapped")
	}

	if d.MapFeature() == nil {
		t.Errorf("feature nil")
	}

	if d.MapRef() == nil {
		t.Errorf("ref nil")
	}

	if d.UnmapFeature() == nil {
		t.Errorf("feature nil")
	}

	if d.UnmapRef() == nil {
		t.Errorf("ref nil")
	}

	d.Free()
	d.Free() // double delete
	d = nil
	d.Free() // check no crash

	if d.Map("http://fake.uri") != 0 {
		t.Errorf("not freed")
	}

	if d.MapFeature() != nil {
		t.Errorf("not freed")
	}

	if d.MapRef() != nil {
		t.Errorf("not freed")
	}

	if d.Unmap(1) != "" {
		t.Errorf("not freed")
	}

	if d.UnmapFeature() != nil {
		t.Errorf("not freed")
	}

	if d.UnmapRef() != nil {
		t.Errorf("not freed")
	}
}

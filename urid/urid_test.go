package urid_test

import (
	"testing"

	"github.com/lvtk/go/urid"
)

func TestMapping(t *testing.T) {
	var d urid.Directory
	defer d.Free()
	urid := d.Map("http://test.org")
	if urid != 1 {
		t.Fatalf("zero not allowd for mapping")
	}
	if "http://test.org" != d.Unmap(urid) {
		t.Fatalf("not mapped")
	}
}

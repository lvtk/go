package lv2_test

import (
	"testing"

	"github.com/lvtk/go/lv2"
)

func TestFeatureNameEmpty(t *testing.T) {
	feature := new(lv2.Feature)
	if len(feature.URI) > 0 {
		t.Errorf("URI not empty")
	}
}

func TestFeatureDataNil(t *testing.T) {
	feature := new(lv2.Feature)
	if feature.Data != nil {
		t.Errorf("Data not nil")
	}
	feature.Free()
}

func TestFeatureSlice(t *testing.T) {
	fs := lv2.NewFeatures()
	defer fs.Clear()
	if fs.Size() != 0 {
		t.Fatalf("size incorrect")
	}

	nf := new(lv2.Feature)
	nf.URI = "http://hello.org"
	nf.Data = nil
	fs.Append(nf)

	if fs.Size() != 1 {
		t.Fatalf("size incorrect")
	}
}

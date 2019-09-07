package lv2

import (
	"testing"
)

func TestFeatureNameEmpty(t *testing.T) {
	feature := new(Feature)
	if len(feature.URI) > 0 {
		t.Errorf("URI not empty")
	}
}

func TestFeatureDataNil(t *testing.T) {
	feature := new(Feature)
	if feature.Data != nil {
		t.Errorf("Data not nil")
	}
}

func TestFeatureSlice(t *testing.T) {
	fs := NewFeatures()
	defer fs.Clear()
	if fs.Size() != 0 {
		t.Fatalf("size incorrect")
	}

	nf := new(Feature)
	nf.URI = "http://hello.org"
	nf.Data = nil
	fs.Append(nf)

	if fs.Size() != 1 {
		t.Fatalf("size incorrect")
	}
}

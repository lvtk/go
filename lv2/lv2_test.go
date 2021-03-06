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

func TestFeatureList(t *testing.T) {
	fs := lv2.NewFeatureList()
	defer fs.Free()

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

	if fs.Ref() == nil {
		t.Fatalf("C features are nil")
	}

	fs.Clear()
	if fs.Size() != 0 {
		t.Errorf("list not cleared")
	}

	fs = nil
	fs.Free()
	fs.Clear()
	fs.Append(nil)

	if fs.Ref() != nil {
		t.Errorf("ref not cleared")
	}

	if fs.Size() > 0 {
		t.Errorf("size incorrect of nil FeatureList")
	}
}

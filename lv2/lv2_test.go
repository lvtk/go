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

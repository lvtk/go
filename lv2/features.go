package lv2

/*
#include "cgo_helpers.h"
#include <assert.h>
#include <stdint.h>

// allocate new feature array with NULL-termination
static LV2_Feature** alloc_features() {
	LV2_Feature** fs = (LV2_Feature**) malloc (sizeof (LV2_Feature*));
	fs[0] = NULL;
	return fs;
}

// resize the feature array to count
static LV2_Feature** realloc_features(LV2_Feature** fs, uint32_t count) {
	assert(count > 0);
	fs = (LV2_Feature**) realloc (fs, count * (uint32_t)sizeof (LV2_Feature*));
	fs[count - 1] = NULL;
	return fs;
}

// assign a feature to the array by index
static void set_feature (LV2_Feature** fs, uint32_t index, LV2_Feature* f) {
	fs[index] = f;
}

*/
import "C"
import (
	"unsafe"
)

// Features a managed list of features
type Features struct {
	features []*Feature
	cfeats   **C.LV2_Feature
	count    uint32
}

// NewFeatures - create new feature list
func NewFeatures() *Features {
	f := new(Features)
	f.cfeats = C.alloc_features()
	f.count = 1
	return f
}

// Clear the list and free memory. This must be called
// when finished with the array or you will leak memory
func (f *Features) Clear() {
	C.free(unsafe.Pointer(f.cfeats))
	f.cfeats = C.alloc_features()
	f.count = 1

	for i := 0; i < len(f.features); i++ {
		ft := f.features[i]
		if ft != nil {
			ft.Free()
		}
	}

	f.features = make([]*Feature, 0)
}

// Size returns the number of features stored
func (f *Features) Size() uint32 {
	if f.count < 1 {
		return 0
	}
	return f.count - 1
}

// Get a raw feature array which is NULL-terminated. Can be
// passed to c functions that have const *LV2_Feature *const params
func (f *Features) Get() unsafe.Pointer {
	return unsafe.Pointer(f.cfeats)
}

// Append a new feature to the list
func (f *Features) Append(feature *Feature) {
	if f == nil || feature == nil {
		return
	}
	f.features = append(f.features, feature)
	f.count++
	f.cfeats = C.realloc_features(f.cfeats, C.uint32_t(f.count))
	C.set_feature(f.cfeats, C.uint32_t(f.count-2), feature.Ref())
}

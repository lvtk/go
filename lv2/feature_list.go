package lv2

/*
#include "cgo_helpers.h"
#include <assert.h>
#include <stdint.h>

// allocate new feature array with NULL-termination
static LV2_Feature** alloc_features(uint32_t size) {
	if (size < 1)
		size = 1;
	LV2_Feature** fs = (LV2_Feature**) malloc (size * sizeof (LV2_Feature*));
	fs[size - 1] = NULL;
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

static LV2_Feature* get_feature (LV2_Feature** fs, uint32_t index) {
	return fs[index];
}

static uint32_t feature_list_size (LV2_Feature** fs) {
	uint32_t i = 0;
	for (const LV2_Feature* f = fs[i]; f != NULL; ++i)
		;
	return i + 1;
}

*/
import "C"
import (
	"unsafe"
)

// FeatureList a managed list of features
// LV2 iteself doesn't provide a way to manage feature lists. This type exists
// so other go packages can easily pass const *LV2_Feature *const to other C
// libraries as a parameter (e.g. lilv, suil, and so on...)
type FeatureList struct {
	features []*Feature
	cfeats   **C.LV2_Feature
	count    uint32
}

// NewFeatureList - create new feature list
func NewFeatureList() *FeatureList {
	f := new(FeatureList)
	f.cfeats = C.alloc_features(1)
	f.count = 1
	return f
}

// NewFeatureListRef creates a new list from existing ref
// @param ref MUST be a pointer to a null terminated LV2_Feature** array
func NewFeatureListRef(ref unsafe.Pointer) *FeatureList {
	fs := NewFeatureList()
	cfs := (**C.LV2_Feature)(ref)
	size := uint32(C.feature_list_size(cfs))
	for i := uint32(0); i < size; i++ {
		f := NewFeatureRef(unsafe.Pointer(C.get_feature(cfs, C.uint32_t(i))))
		fs.Append(f)
	}

	return fs
}

// Free - This must be called when finished with the array or you will leak memory
func (f *FeatureList) Free() {
	if f == nil {
		return
	}

	if f.cfeats != nil {
		C.free(unsafe.Pointer(f.cfeats))
		f.cfeats = nil
	}

	// features slice should contain all referenced data. if not
	// this will make sure they get freed
	if f.features != nil {
		for i := 0; i < len(f.features); i++ {
			ft := f.features[i]
			if ft != nil {
				ft.Free()
			}
		}
		f.features = nil
	}
}

// Clear the list and free memory.
func (f *FeatureList) Clear() {
	if f == nil {
		return
	}

	f.Free()

	f.cfeats = C.alloc_features(1)
	f.count = 1
	f.features = make([]*Feature, 0)
}

// Size returns the number of features stored
func (f *FeatureList) Size() uint32 {
	if f == nil || f.count < 1 {
		return 0
	}
	return f.count - 1
}

// Ref a raw feature array which is NULL-terminated. Can be
// passed to c functions that have const *LV2_Feature *const params
func (f *FeatureList) Ref() unsafe.Pointer {
	if f == nil || f.cfeats == nil {
		return nil
	}
	return unsafe.Pointer(f.cfeats)
}

// Append a new feature to the list
func (f *FeatureList) Append(feature *Feature) {
	if f == nil || feature == nil {
		return
	}
	f.features = append(f.features, feature)
	f.count++
	f.cfeats = C.realloc_features(f.cfeats, C.uint32_t(f.count))
	C.set_feature(f.cfeats, C.uint32_t(f.count-2), feature.Ref())
}

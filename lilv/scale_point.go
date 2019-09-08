package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
*/
import "C"
import "unsafe"

/* ScalePoint */

/*
Label - Get the label of this scale point (enumeration value)
Returned value is owned by `point` and must not be freed.
*/
func (sp *ScalePoint) Label() *Node {
	if sp == nil || sp.scalePoint == nil {
		return nil
	}
	return createSharedNode(C.lilv_scale_point_get_label(sp.scalePoint))
}

/*
Value - Get the value of this scale point (enumeration value)
Returned value is owned by `point` and must not be freed.
*/
func (sp *ScalePoint) Value() *Node {
	if sp == nil || sp.scalePoint == nil {
		return nil
	}
	return createSharedNode(C.lilv_scale_point_get_value(sp.scalePoint))
}

/* ScalePoints */

// Free - free this ScalePoints collection
func (sps *ScalePoints) Free() {
	if sps == nil || sps.scalePoints == nil {
		return
	}
	if sps.managed {
		C.lilv_scale_points_free(sps.scalePoints)
	}
	sps.scalePoints = nil
}

// Size get the number of sps in the collection
func (sps *ScalePoints) Size() uint32 {
	if sps == nil || sps.scalePoints == nil {
		return 0
	}
	return uint32(C.lilv_scale_points_size(sps.scalePoints))
}

// Begin the start iter
func (sps *ScalePoints) Begin() *Iter {
	if sps == nil || sps.scalePoints == nil {
		return nil
	}
	return (*Iter)(C.lilv_scale_points_begin(sps.scalePoints))
}

// Get the ui for this iter
func (sps *ScalePoints) Get(iter *Iter) *ScalePoint {
	if sps == nil || sps.scalePoints == nil {
		return nil
	}
	sp := new(ScalePoint)
	sp.scalePoint = C.lilv_scale_points_get(sps.scalePoints, unsafe.Pointer(iter))
	return sp
}

// Next iterator
func (sps *ScalePoints) Next(iter *Iter) *Iter {
	if sps == nil || sps.scalePoints == nil {
		return nil
	}
	return (*Iter)(C.lilv_scale_points_next(sps.scalePoints,
		unsafe.Pointer(iter)))
}

// IsEnd true if end iter
func (sps *ScalePoints) IsEnd(iter *Iter) bool {
	if sps == nil || sps.scalePoints == nil {
		return false
	}
	return bool(C.lilv_scale_points_is_end(sps.scalePoints, unsafe.Pointer(iter)))
}

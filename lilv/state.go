package lilv

/*
#cgo pkg-config: lilv-0
#include <lilv/lilv.h>
*/
import "C"

// Free - destroy the state
func (s *State) Free() {
	if s == nil || s.state == nil {
		return
	}
	if s.state != nil {
		C.lilv_state_free(s.state)
	}
	s.state = nil
}

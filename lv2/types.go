// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 05 Sep 2019 15:48:59 EDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package lv2

/*
#cgo pkg-config: lv2
#include <lv2/lv2plug.in/ns/lv2core/lv2.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Handle type as declared in lv2core/lv2.h:133
type Handle unsafe.Pointer

// Feature as declared in lv2core/lv2.h:158
type Feature struct {
	URI            string
	Data           unsafe.Pointer
	refc2185671    *C.LV2_Feature
	allocsc2185671 interface{}
}

// LibHandle type as declared in lv2core/lv2.h:399
type LibHandle unsafe.Pointer

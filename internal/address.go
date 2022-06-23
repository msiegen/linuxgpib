// Copyright 2022 Google LLC
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// version 2 as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

package internal

/*
#cgo linux LDFLAGS: -lgpib
#include <stdlib.h>
#include <gpib/ib.h>
*/
import "C"

// Address is a GPIB device address, composed of a primary and a secondary
// address. If the secondary address is zero, as is the case with many devices,
// then the value of an Address is equal to its primary address.
type Address int

// NewAddress combines a primary and secondary address.
func NewAddress(pad, sad int) Address {
	return Address(C.MakeAddr(C.uint(pad), C.uint(sad)))
}

// Primary returns the primary address.
func (a Address) Primary() int {
	return int(C.GetPAD(C.Addr4882_t(a)))
}

// Secondary returns the secondary address.
func (a Address) Secondary() int {
	return int(C.GetSAD(C.Addr4882_t(a)))
}

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
#include "status.h"
*/
import "C"

// ThreadIbcnt, ThreadIberr, and ThreadIbsta cannot be used from Go because
// there is not guarantee that the caller goroutine will be multiplexed onto
// the same OS thread as was used for the last IO operation. Therefore we
// provide accessor functions for the ibcnt, iberr, and ibsta globals.

// Ibcnt returns the number of bytes sent or received by the last IO operation.
// It is also set to the value of errno after EDVR or EFSO errors.
func Ibcnt() int { return int(C.globalIbcnt()) }

// Iberr returns the last error. The meaning of each possible value is
// summarized in:
// https://linux-gpib.sourceforge.io/doc_html/reference-globals-iberr.html
func Iberr() int { return int(C.globalIberr()) }

// Ibsta returns the last status. The meaning of the bits is summarized in:
// https://linux-gpib.sourceforge.io/doc_html/reference-globals-ibsta.html
func Ibsta() int { return int(C.globalIbsta()) }

// Test helpers are defined here because test files cannot import C directly.
func testNOADDR() int  { return int(C.testNOADDR()) }
func testSTOPend() int { return int(C.testSTOPend()) }

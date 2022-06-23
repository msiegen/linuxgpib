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

#ifndef _LINUXGPIB_STATUS_H
#define _LINUXGPIB_STATUS_H

#include <gpib/ib.h>

// Define inline functions to return the value of C static variables, which
// cannot be accessed directly due to https://github.com/golang/go/issues/15980

static __inline__ int globalIbcnt(void) {
  return ibcnt;
}

static __inline__ int globalIberr(void) {
  return iberr;
}

static __inline__ int globalIbsta(void) {
  return ibsta;
}

static __inline__ Addr4882_t testNOADDR(void) {
  return NOADDR;
}

static __inline__ int testSTOPend(void) {
  return STOPend;
}

#endif

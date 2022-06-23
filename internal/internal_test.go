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

import (
	"os"
	"testing"
)

func TestTimeoutError(t *testing.T) {
	if !os.IsTimeout(TimeoutErr) {
		t.Fatal("os.IsTimeout(TimeoutErr) is false")
	}
}

func TestNOADDR(t *testing.T) {
	if g := testNOADDR(); g != NOADDR {
		t.Errorf("bad NOADDR: got 0x%x; want 0x%x", g, NOADDR)
	}
}

func TestSTOPend(t *testing.T) {
	if g := testSTOPend(); g != STOPend {
		t.Errorf("bad STOPend: got 0x%x; want 0x%x", g, STOPend)
	}
}

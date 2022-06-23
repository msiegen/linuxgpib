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

package linuxgpib

import (
	"testing"
)

func TestFormatLog(t *testing.T) {
	for i, c := range []struct {
		Input string
		Want  string
	}{
		{
			"0123456789",
			"\"0123456789\"",
		},
		{
			"0123456789----------0123456789----------0123456789----------",
			"\"0123456789----------0123456789----------0123456789----------\"",
		},
		{
			"0123456789----------0123456789----------0123456789----------0",
			"\"0123456789----------0123456789----------\"...(61 bytes total)",
		},
	} {
		g := formatLog([]byte(c.Input))
		if g != c.Want {
			t.Errorf("%d: got %v, want %v", i, g, c.Want)
		}
	}
}

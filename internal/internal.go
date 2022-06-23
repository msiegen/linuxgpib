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

// Package internal provides access to Linux-GPIB's C functions and constants.
//
// The functions here are small wrappers around the C functions in
// https://linux-gpib.sourceforge.io/doc_html/reference.html and are *not* safe
// for concurrent use.
package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type timeoutError struct{ error }

func (t *timeoutError) Timeout() bool { return true }

var (
	// TimeoutErr is returned by Err if an operation timed out.
	TimeoutErr = &timeoutError{errors.New("timed out")}

	ibstaCodes = []int{
		DCAS,
		DTAS,
		LACS,
		TACS,
		ATN,
		CIC,
		REM,
		LOK,
		CMPL,
		EVENT,
		SPOLL,
		RQS,
		SRQI,
		END,
		TIMO,
		ERR,
	}
	ibstaStrings = map[int]string{
		DCAS:  "DCAS",
		DTAS:  "DTAS",
		LACS:  "LACS",
		TACS:  "TACS",
		ATN:   "ATN",
		CIC:   "CIC",
		REM:   "REM",
		LOK:   "LOK",
		CMPL:  "CMPL",
		EVENT: "EVENT",
		SPOLL: "SPOLL",
		RQS:   "RQS",
		SRQI:  "SRQI",
		END:   "END",
		TIMO:  "TIMO",
		ERR:   "ERR",
	}
	iberrStrings = map[int]string{
		EDVR: "EDVR",
		ECIC: "ECIC",
		ENOL: "ENOL",
		EADR: "EADR",
		EARG: "EARG",
		ESAC: "ESAC",
		EABO: "EABO",
		ENEB: "ENEB",
		EDMA: "EDMA",
		EOIP: "EOIP",
		ECAP: "ECAP",
		EFSO: "EFSO",
		EBUS: "EBUS",
		ESTB: "ESTB",
		ESRQ: "ESRQ",
		ETAB: "ETAB",
	}
)

// formatIberr returns the name of the error enum constant.
func formatIberr(iberr int) string {
	if s, ok := iberrStrings[iberr]; ok {
		return s
	}
	return "UNKNOWN"
}

// Err returns an error if ibsta has the ERR bit set, and nil otherwise. If the
// TIMO bit is set, TimeoutErr is returned.
//
// For non-timeout errors, Err accesses the iberr and ibcnt globals. It must
// therefore be called prior to any subsequent operations which might overwrite
// those globals.
func Err(ibsta int) error {
	if ibsta&TIMO != 0 {
		return TimeoutErr
	}
	if ibsta&ERR != 0 {
		v := Iberr()
		switch v {
		case EDVR:
			errno := syscall.Errno(Ibcnt())
			return fmt.Errorf("EDVR: %v", errno)
		case EFSO:
			errno := syscall.Errno(Ibcnt())
			return fmt.Errorf("EFSO: %v", errno)
		default:
			return errors.New(formatIberr(int(v)))
		}
	}
	return nil
}

// FormatIbsta returns a string representation of ibsta.
func FormatIbsta(ibsta int) string {
	bits := make([]string, 0, len(ibstaCodes))
	for i := len(ibstaCodes) - 1; i >= 0; i-- {
		code := ibstaCodes[i]
		if ibsta&code != 0 {
			bits = append(bits, ibstaStrings[code])
		}
	}
	return strings.Join(bits, " ")
}

// FormatIblines returns the bus line states in human-readable form.
func FormatIblines(iblines int) string {
	s := make([]string, 0, 9)
	s = append(s, strconv.FormatInt(int64(iblines), 16))
	for _, l := range []struct {
		valid, bit int
		f, t       string
	}{
		{0x80, 0x8000, "eoi", "EOI"},
		{0x40, 0x4000, "atn", "ATN"},
		{0x20, 0x2000, "srq", "SRQ"},
		{0x10, 0x1000, "ren", "REN"},
		{0x8, 0x800, "ifc", "IFC"},
		{0x4, 0x400, "nrfd", "NRFD"},
		{0x2, 0x200, "ndac", "NDAC"},
		{0x1, 0x100, "dav", "DAV"},
	} {
		if l.valid&iblines != 0 {
			if l.bit&iblines == 0 {
				s = append(s, l.f)
			} else {
				s = append(s, l.t)
			}
		}
	}
	if len(s) == 0 {
		return "unavailable"
	}
	return strings.Join(s, " ")
}

// IblinesString returns the bus line states in human-readable form.
func IblinesString(board int) string {
	ibsta, iblines := Iblines(board)
	if err := Err(ibsta); err != nil {
		return err.Error()
	}
	return FormatIblines(iblines)
}

// Timeout returns a timeout constant not shorter than the specified
// duration. The returned timeout may be longer, up to the max supported by
// GPIB.
func Timeout(d time.Duration) int {
	for _, c := range []struct {
		D time.Duration
		C int
	}{
		{0, TNONE},
		{10 * time.Microsecond, T10us},
		{30 * time.Microsecond, T30us},
		{100 * time.Microsecond, T100us},
		{300 * time.Microsecond, T300us},
		{1 * time.Millisecond, T1ms},
		{3 * time.Millisecond, T3ms},
		{10 * time.Millisecond, T10ms},
		{30 * time.Millisecond, T30ms},
		{100 * time.Millisecond, T100ms},
		{300 * time.Millisecond, T300ms},
		{1 * time.Second, T1s},
		{3 * time.Second, T3s},
		{10 * time.Second, T10s},
		{30 * time.Second, T30s},
		{100 * time.Second, T100s},
		{300 * time.Second, T300s},
	} {
		if d <= c.D {
			return c.C
		}
	}
	return T1000s
}

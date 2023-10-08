// Copyright 2023 Google LLC
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// version 2 as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

/*
Identify queries the identification string of a GPIB device.

It works for [SCPI] devices that support the *IDN? query.

Usage:

	identify [-verbose] [-board=BOARD] -address=ADDRESS

The flags are:

	-verbose
		Turn on logging. Without this, only the result or first error is printed.

	-board
		The board number. Defaults to zero, which corresponds to /dev/gpib0.

	-address
		The primary address of the GPIB device to query.

Examples:

	$ identify -address 22
	HEWLETT-PACKARD,34401A,0,10-5-2
	$

	$ identify -verbose -address 22
	2023/10/08 14:26:53 linuxgpib.go:132 Opened board 0 with version 4.3.4
	2023/10/08 14:26:53 linuxgpib.go:210 Opened address 22 (22/0) on board 0 as device 16
	2023/10/08 14:26:53 linuxgpib.go:434 Wrote "*IDN?\n" in 2ms to address 22
	2023/10/08 14:26:53 linuxgpib.go:333 Read "HEWLETT-PACKARD,34401A,0,10-5-2\n" in 35ms from address 22
	HEWLETT-PACKARD,34401A,0,10-5-2
	2023/10/08 14:26:53 linuxgpib.go:286 Closing address 22
	$

[SCPI]: https://en.wikipedia.org/wiki/Standard_Commands_for_Programmable_Instruments
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/msiegen/linuxgpib"
)

func main() {
	verbose := flag.Bool(
		"verbose", false,
		"Turn on logging. Without this, only the result or first error is printed.",
	)
	board := flag.Int(
		"board", 0,
		"The board number. Defaults to zero, which corresponds to /dev/gpib0.",
	)
	address := flag.Int(
		"address", -1,
		"The primary address of the GPIB device to query.",
	)

	flag.Parse()

	if *address < 0 {
		fmt.Fprintln(os.Stderr, "Please specify an -address!")
		os.Exit(1)
	}

	var opts []linuxgpib.Option
	if *verbose {
		opts = append(opts, linuxgpib.Log(log.Default()))
	}

	d, err := linuxgpib.NewDevice(*board, *address, opts...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open device:", err)
		os.Exit(1)
	}
	defer d.Close()

	if _, err := fmt.Fprintln(d, "*IDN?"); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write to device:", err)
		os.Exit(1)
	}

	r := bufio.NewReader(d)
	s, err := r.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read from device:", err)
		os.Exit(1)
	}

	fmt.Println(strings.Trim(s, " \r\n"))
}

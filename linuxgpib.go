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

// Package linuxgpib enables communication over the IEEE-488 bus.
//
// To use it you must install https://linux-gpib.sourceforge.io/.
package linuxgpib

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/msiegen/linuxgpib/internal"
)

const (
	defaultTimeout = 10 * time.Second
	maxLogData     = 60
	minLogHide     = 20 // must be smaller tha maxLogData
)

var (
	// Internal functions are not safe for concurrent access. mu allows clients to
	// coordinate their access so that multiple goroutines multiplexed to a single
	// OS thread do not step on each others' ibsta, iberr, and ibcnt values.
	mu sync.Mutex
	// Keep a map of boards that are in use to prevent duplicate instances.
	activeBoards = map[int]bool{}
)

// Logger writes lines of output for debug purposes.
type Logger interface {
	Printf(string, ...interface{})
}

type options struct {
	timeout  time.Duration
	readEOS  string
	logger   Logger
	activity func(bool)
}

func newOptions() *options {
	return &options{
		timeout: defaultTimeout,
	}
}

func cloneOptions(o *options) *options {
	n := *o
	return &n
}

func (o *options) logf(format string, v ...interface{}) {
	if o.logger != nil {
		if _, f, l, ok := runtime.Caller(1); ok {
			format = filepath.Base(f) + ":" + strconv.Itoa(l) + " " + format
		}
		o.logger.Printf(format, v...)
	}
}

// An Option configures GPIB communication with the device.
type Option func(*options)

// Log enables logging GPIB traffic in human readable form.
func Log(l Logger) Option {
	return func(o *options) {
		o.logger = l
	}
}

// Timeout sets the timeout for GPIB operations. It defaults to 10s and may be
// changed at runtime by calling SetTimeout.
//
// The duration will be rounded up to one of the discrete values in
// https://linux-gpib.sourceforge.io/doc_html/reference-function-ibtmo.html
func Timeout(d time.Duration) Option {
	return func(o *options) {
		o.timeout = d
	}
}

// ReadEOS enables the termination of reads when the specified character is
// received. If set to the empty string, the default, reads are terminated when
// the remote device asserts EOI.
func ReadEOS(char string) Option {
	return func(o *options) {
		o.readEOS = char
	}
}

// Activity registers a callback whereby a callee can be informed of activity on
// the bus. This can be used to control an indicator lamp, for example.
func Activity(f func(bool)) Option {
	return func(o *options) {
		o.activity = f
	}
}

// Board is a GPIB interface board.
type Board struct {
	index         int
	options       *options
	activeDevices map[int]bool
}

func NewBoard(index int, opts ...Option) (*Board, error) {
	o := newOptions()
	for _, opt := range opts {
		opt(o)
	}
	mu.Lock()
	defer mu.Unlock()
	if activeBoards[index] {
		return nil, fmt.Errorf("board in use: %d", index)
	}
	activeBoards[index] = true
	o.logf("Opened board %d with version %v", index, internal.Ibvers())
	return &Board{
		index:         index,
		options:       o,
		activeDevices: map[int]bool{},
	}, nil
}

// Device is a connection to a single GPIB device.
//
// All methods acquire the global GPIB lock for the duration of their
// execution, making it safe to use multiple devices each from a different
// goroutine.
type Device struct {
	addr     int
	board    *Board
	ud       int
	options  *options
	isClosed bool
}

// NewDevice returns a GPIB device.
//
// Board is the board index, 0 for the first board. Address is normally the
// primary address of the device. Secondary addresses are supported by casting
// a internal.Address to an int.
func (b *Board) NewDevice(addr int, opts ...Option) (*Device, error) {
	a := internal.Address(addr)
	o := cloneOptions(b.options)
	for _, opt := range opts {
		opt(o)
	}

	var eos int
	switch len(o.readEOS) {
	case 0:
		break
	case 1:
		eos = internal.BIN | internal.REOS | int(o.readEOS[0])
	default:
		return nil, errors.New("invalid read eos: must be a single character")
	}

	mu.Lock()
	defer mu.Unlock()

	if b.activeDevices[addr] {
		return nil, fmt.Errorf("device already in use: %d", addr)
	}

	if b.options.activity != nil {
		b.options.activity(true)
		defer b.options.activity(false)
	}

	if len(b.activeDevices) == 0 {
		if err := internal.Err(internal.Ibsre(b.index, 1)); err != nil {
			o.logf("Failed to enable remote mode on board %d", b.index)
			return nil, errors.New("ibsre failed")
		}
	}

	// Open the device.
	pad := a.Primary()
	sad := a.Secondary()
	tmo := internal.Timeout(o.timeout)
	ud := internal.Ibdev(b.index, pad, sad, tmo, 1 /*eoi*/, eos)
	if ud == -1 {
		if err := internal.Err(internal.Ibsta()); err != nil {
			o.logf("Failed to open address %d (%d/%d) on board %d: %v", addr, pad, sad, b.index, err)
			return nil, err
		}
		o.logf("Failed to open address %d (%d/%d) on board %d: unknown error", addr, pad, sad, b.index)
		return nil, errors.New("ibdev failed without setting an error")
	}

	b.activeDevices[addr] = true

	o.logf("Opened address %d (%d/%d) on board %d as device %d", addr, pad, sad, b.index, ud)
	return &Device{
		addr:    addr,
		board:   b,
		ud:      ud,
		options: o,
	}, nil
}

// Clear issues a GPIB device clear command.
func (d *Device) Clear() error {
	mu.Lock()
	defer mu.Unlock()
	if d.isClosed {
		return errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	// Clear the device.
	d.options.logf("Clearing device at address %d", d.addr)
	if err := internal.Err(internal.Ibclr(d.ud)); err != nil {
		d.options.logf("Failed to clear device %d: %v", d.ud, err)
		return err
	}

	// Wait for the device to unassert "not ready for data". Some devices will
	// cause a timeout if a write is attempted immediately after the device is
	// cleared.
	cleared := time.Now()
	for {
		time.Sleep(50 * time.Millisecond)
		ibsta, lines := internal.Iblines(d.board.index)
		if err := internal.Err(ibsta); err != nil {
			d.options.logf("Failed to monitor iblines after clearing device %d: %v", d.ud, err)
			return err
		}
		if lines&internal.ValidNRFD == 0 {
			// The BusNRFD bit is invalid. We won't be able to tell when the device is
			// ready, so just use a generous delay.
			time.Sleep(1 * time.Second)
			break
		}
		if lines&internal.BusNRFD == 0 {
			// The device is ready!
			break
		}
		if d.options.timeout != 0 && time.Now().Sub(cleared) > d.options.timeout {
			d.options.logf("Timed out after clearing device %d", d.ud)
			return internal.TimeoutErr
		}
	}

	return nil
}

// Close releases resources associated with the GPIB device.
func (d *Device) Close() error {
	mu.Lock()
	defer mu.Unlock()

	if d.isClosed {
		return errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	d.isClosed = true
	delete(d.board.activeDevices, d.addr)

	d.options.logf("Closing address %d", d.addr)
	if err := internal.Err(internal.Ibonl(d.ud, 0)); err != nil {
		d.options.logf("Failed to close address %d device %d: %v", d.addr, d.ud, err)
		return err
	}

	if len(d.board.activeDevices) == 0 {
		if err := internal.Err(internal.Ibsre(d.board.index, 0)); err != nil {
			d.options.logf("Failed to disable remote mode on board %d", d.board.index)
			return errors.New("ibsre failed")
		}
	}

	return nil
}

// formatLog returns a possibly shortened representation of the input data for
// logging.
func formatLog(b []byte) string {
	if len(b) > maxLogData {
		return fmt.Sprintf("%q...(%d bytes total)", string(b[:maxLogData-minLogHide]), len(b))
	}
	return fmt.Sprintf("%q", string(b))
}

// Read gets data from the GPIB device.
func (d *Device) Read(b []byte) (n int, err error) {
	mu.Lock()
	defer mu.Unlock()
	if d.isClosed {
		return 0, errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	started := time.Now()
	ibsta := internal.Ibrd(d.ud, b)
	took := time.Since(started)
	err = internal.Err(ibsta)
	n = internal.Ibcnt()

	if err != nil {
		d.options.logf("Failed to read from address %d device %d: %v", d.addr, d.ud, err)
	} else {
		d.options.logf("Read %s in %v from address %d", formatLog(b[:n]), took.Truncate(time.Millisecond), d.addr)
	}
	return
}

// SetTimeout changes the timeout for future GPIB operations.
//
// The duration will be rounded up to one of the discrete values in
// https://linux-gpib.sourceforge.io/doc_html/reference-function-ibtmo.html
func (d *Device) SetTimeout(t time.Duration) error {
	mu.Lock()
	defer mu.Unlock()
	if d.isClosed {
		return errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	d.options.logf("Setting timeout to %v on address %d", t, d.addr)
	if err := internal.Err(internal.Ibtmo(d.ud, internal.Timeout(t))); err != nil {
		d.options.logf("Failed to set timeout on address %d device %d: %v", d.addr, d.ud, err)
		return err
	}
	return nil
}

// Spoll gets the status byte from a device via serial poll.
func (d *Device) Spoll() (byte, error) {
	mu.Lock()
	defer mu.Unlock()
	if d.isClosed {
		return 0, errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	started := time.Now()
	ibsta, spr := internal.Ibrsp(d.ud)
	took := time.Since(started)
	if err := internal.Err(ibsta); err != nil {
		d.options.logf("Failed to poll address %d device %d: %v", d.addr, d.ud, err)
		return 0, err
	}
	d.options.logf("Polled status %02X in %v from address %d", spr, took.Truncate(time.Millisecond), d.addr)

	return spr, nil
}

// Trigger sends a GET (group execute trigger) command to the device.
func (d *Device) Trigger() error {
	mu.Lock()
	defer mu.Unlock()
	if d.isClosed {
		return errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	d.options.logf("Triggering device at address %d", d.addr)
	ibsta := internal.Ibtrg(d.ud)
	if err := internal.Err(ibsta); err != nil {
		d.options.logf("Failed to trigger address %d device %d: %v", d.addr, d.ud, err)
		return err
	}

	return nil
}

// Write sends data to the GPIB device.
func (d *Device) Write(b []byte) (n int, err error) {
	mu.Lock()
	defer mu.Unlock()
	if d.isClosed {
		return 0, errors.New("already closed")
	}

	if d.options.activity != nil {
		d.options.activity(true)
		defer d.options.activity(false)
	}

	started := time.Now()
	ibsta := internal.Ibwrt(d.ud, b)
	took := time.Since(started)
	err = internal.Err(ibsta)
	n = internal.Ibcnt()

	if err != nil {
		d.options.logf("Failed to write to address %d device %d: %v", d.addr, d.ud, err)
		return
	}

	d.options.logf("Wrote %s in %v to address %d", formatLog(b), took.Truncate(time.Millisecond), d.addr)

	return
}

// Enumerate returns the primary addresses of all devices on the bus.
func (b *Board) Enumerate() ([]int, error) {
	mu.Lock()
	defer mu.Unlock()

	if b.options.activity != nil {
		b.options.activity(true)
		defer b.options.activity(false)
	}

	started := time.Now()

	// Perform an interface clear so that all devices untalk and unlisten. This is
	// necessary because some older devices like the HP 3478A, if previously
	// addressed as talker, will write data to the bus as soon as another device
	// is addressed as a listener by ibln.
	ibsta := internal.Ibsic(b.index)
	if err := internal.Err(ibsta); err != nil {
		b.options.logf("Board %d returned ibsic error: %v", b.index, err)
		return nil, err
	}

	// Verify that the board has the capabilities needed for enumeration.
	ibsta, iblines := internal.Iblines(b.index)
	if err := internal.Err(ibsta); err != nil {
		b.options.logf("Board %d returned iblines error: %v", b.index, err)
		return nil, err
	}
	if iblines&internal.ValidNDAC == 0 {
		b.options.logf("Board %d does not support monitoring NDAC", b.index)
		return nil, errors.New("board does not support monitoring NDAC")
	}

	// Check all the addresses to see if a listener is present, except address 0
	// which is the controller.
	var ds []int
	for i := 1; i <= 30; i++ {
		ibsta, found := internal.Ibln(b.index, i, 0)
		if err := internal.Err(ibsta); err != nil {
			b.options.logf("Failed to enumerate board %d address %d: %v", b.index, i, err)
			return nil, err
		}
		if found != 0 {
			b.options.logf("Found device at address %d on board %d", i, b.index)
			ds = append(ds, i)
		}
	}

	b.options.logf("Found %d devices in %v on board %d", len(ds), time.Since(started).Truncate(time.Millisecond), b.index)

	return ds, nil
}

// NewDevice returns a GPIB device. See Board's NewDevice method for more
// details.
func NewDevice(board, addr int, opts ...Option) (*Device, error) {
	b, err := NewBoard(board, opts...)
	if err != nil {
		return nil, err
	}
	return b.NewDevice(addr)
}

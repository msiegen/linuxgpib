[![Go Reference](https://pkg.go.dev/badge/github.com/msiegen/linuxgpib.svg)](https://pkg.go.dev/github.com/msiegen/linuxgpib)

# Go Library for Linux GPIB

The `linuxgpib` package provides an idiomatic Go interface to the
[Linux GPIB](https://linux-gpib.sourceforge.io/) C library. It's handy if you
have vintage electronics test equipment on an
[IEEE-488](https://en.wikipedia.org/wiki/IEEE-488)
bus and want to control it or perform data acquisition using the
[Go language](https://go.dev/).

## Usage

You can use Go's standard Reader and Writer interfaces to communicate with
instruments. A simple transaction to read the identity of a
[SCPI](https://en.wikipedia.org/wiki/Standard_Commands_for_Programmable_Instruments)
device at address 22 looks like:

```go
d, err := linuxgpib.NewDevice(0, 22)
_, err = fmt.Fprintln(d, "*IDN?")
result, err := bufio.NewReader(d).ReadString('\n')
d.Close()
```

For a more complete version (with logging and error handling!) see the
[identify command](https://github.com/msiegen/linuxgpib/blob/main/cmd/identify/identify.go)
in this repository.

## Building

To build your application that imports linuxgpib you need to install the
[Linux GPIB](https://linux-gpib.sourceforge.io/) userspace C library per its
[instructions](https://sourceforge.net/p/linux-gpib/code/HEAD/tree/trunk/linux-gpib-user/INSTALL).

It's generally easiest to install the C library systemwide into its default
prefix, which is `/usr/local`. No special flags to the Go compiler should be
needed in that case.

If a systemwide install is not possible (such as when you don't have root), you
can `./configure` it with an alternative `--prefix=/some/path` and then set
variables in the environment before invoking `go build`:

- `CGO_CFLAGS=-I/some/path/include`
- `CGO_LDFLAGS=-L/some/path/lib`

## Running

To run your application you should install the userspace C library described
above _and_ the
[kernel driver](https://sourceforge.net/p/linux-gpib/code/HEAD/tree/trunk/linux-gpib-kernel/INSTALL).
You'll also need to
[configure a board](https://linux-gpib.sourceforge.io/doc_html/configuration.html)
and pass its minor number (aka index) to the linuxgpib Go code.

In certain scenarios the dynamic link loader may fail to find libgpib.so.0. If
that happens to you, give it an extra hint with an environment variable to the
path where you installed the userspace C library:

- `LD_LIBRARY_PATH=/usr/local/lib`

## Notes
This is not an officially supported Google product.

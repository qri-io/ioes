// Package ioes or "in, out, error streams" provides standard names and utilities
// for working with traditional stdin/stdout/stderr streams.
package ioes

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-isatty"
)

// IOStreams provides standard names for iostreams and common methods for managing
// progress indicators & other contextual output. IOStreams is broken out as a
// sepearate package to facilitate plumbing contextual feedback into the depths
// of an application architecture, while also providing clear state management
// for when contextual feedback doesn't work. As an example, IOStreams may be
// routed over websockets to provide HTTP output, and in this context it's far
// easier to disable things like spinners outright
// IOStreams must be created with a "New" method
type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer

	sp *spinner.Spinner
}

// IsTerminal returns true when IOStreams Out is a terminal file descriptor
func (s IOStreams) IsTerminal() bool {
	if osOutFile, ok := s.Out.(*os.File); ok {
		if osOutFile == os.Stdout {
			return isatty.IsTerminal(osOutFile.Fd())
		}
	}
	return false
}

// IsCygwinTerminal returns true when IOStreams Out is a Cygwin file descriptor
func (s IOStreams) IsCygwinTerminal() bool {
	if osOutFile, ok := s.Out.(*os.File); ok {
		if osOutFile == os.Stdout {
			return isatty.IsCygwinTerminal(osOutFile.Fd())
		}
	}
	return false
}

// StartSpinner begins the progress spinner
func (s IOStreams) StartSpinner() {
	s.sp.Start()
}

// StopSpinner halts the progress spinner
func (s IOStreams) StopSpinner() {
	s.sp.Stop()
}

// SpinnerActive returns the active state of the progres spinner
func (s IOStreams) SpinnerActive() bool {
	return s.sp.Active()
}

// SpinnerMsg sets the spinner suffix message
func (s IOStreams) SpinnerMsg(msg string) {
	s.sp.Suffix = msg
}

// Close checks to see if any/all of in/out/errOut are closable,
// and if so closes them
func (s IOStreams) Close() error {
	// TODO
	return nil
}

// Print writes a msg to the out stream, printing
func (s IOStreams) Print(msg string) {
	if s.SpinnerActive() {
		s.StopSpinner()
		defer s.StartSpinner()
	}
	s.Out.Write([]byte(msg))
}

// PrintErr writes a msg to the Err stream, printing
func (s IOStreams) PrintErr(msg string) {
	if s.SpinnerActive() {
		s.StopSpinner()
		defer s.StartSpinner()
	}
	s.ErrOut.Write([]byte(msg))
}

// NewIOStreams creates streams
func NewIOStreams(in io.Reader, out, errOut io.Writer) IOStreams {
	sp := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	sp.Writer = errOut

	return IOStreams{
		In:     in,
		Out:    out,
		ErrOut: errOut,

		sp: sp,
	}
}

// NewStdIOStreams creates a standard set of streams, with in, out, and error mapped
// to os.Stdin, os.Stdout, and os.Stderr respectively
func NewStdIOStreams() IOStreams {
	sp := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	sp.Writer = os.Stderr

	return IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,

		sp: sp,
	}
}

// NewTestIOStreams returns a valid IOStreams and in, out, errout buffers for unit tests
func NewTestIOStreams() (IOStreams, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	errOut := &bytes.Buffer{}
	sp := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	sp.Writer = ioutil.Discard

	return IOStreams{
		In:     in,
		Out:    out,
		ErrOut: errOut,

		sp: sp,
	}, in, out, errOut
}

// NewDiscardIOStreams returns a valid IOStreams that just discards
func NewDiscardIOStreams() IOStreams {
	in := &bytes.Buffer{}
	sp := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	sp.Writer = ioutil.Discard
	return IOStreams{
		In:     in,
		Out:    ioutil.Discard,
		ErrOut: ioutil.Discard,

		sp: sp,
	}
}

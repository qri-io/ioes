// Package ioes or "in, out, error streams" provides standard names and utilities
// for working with traditional "stdout" streams
package ioes

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/briandowns/spinner"
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

// Print writes a msg to the Out stream
func (s IOStreams) Print(msg string) {
	if s.SpinnerActive() {
		s.StopSpinner()
		defer s.StartSpinner()
	}
	s.Out.Write([]byte(msg))
}

// NewIOStreams creates streams
func NewIOStreams(in io.Reader, out, errOut io.Writer) IOStreams {
	return IOStreams{
		In:     in,
		Out:    out,
		ErrOut: errOut,

		sp: spinner.New(spinner.CharSets[24], 100*time.Millisecond),
	}
}

// NewStdIOStreams creates a standard set of streams, with in, out, and error mapped
// to os.Stdin, os.Stdout, and os.Stderr respectively
func NewStdIOStreams() IOStreams {
	return IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,

		sp: spinner.New(spinner.CharSets[24], 100*time.Millisecond),
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

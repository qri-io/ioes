# ioes
--
    import "github.com/qri-io/ioes"

Package ioes or "in, out, error streams" provides standard names and utilities
for working with traditional "stdout" streams

## Usage

#### type IOStreams

```go
type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}
```

IOStreams provides standard names for iostreams and common methods for managing
progress indicators & other contextual output. IOStreams is broken out as a
sepearate package to facilitate plumbing contextual feedback into the depths of
an application architecture, while also providing clear state management for
when contextual feedback doesn't work. As an example, IOStreams may be routed
over websockets to provide HTTP output, and in this context it's far easier to
disable things like spinners outright IOStreams must be created with a "New"
method

#### func  NewDiscardIOStreams

```go
func NewDiscardIOStreams() IOStreams
```
NewDiscardIOStreams returns a valid IOStreams that just discards

#### func  NewIOStreams

```go
func NewIOStreams(in io.Reader, out, errOut io.Writer) IOStreams
```
NewIOStreams creates streams

#### func  NewStdIOStreams

```go
func NewStdIOStreams() IOStreams
```
NewStdIOStreams creates a standard set of streams, with in, out, and error
mapped to os.Stdin, os.Stdout, and os.Stderr respectively

#### func  NewTestIOStreams

```go
func NewTestIOStreams() (IOStreams, *bytes.Buffer, *bytes.Buffer, *bytes.Buffer)
```
NewTestIOStreams returns a valid IOStreams and in, out, errout buffers for unit
tests

#### func (IOStreams) Print

```go
func (s IOStreams) Print(msg string)
```
Print writes a msg to the Out stream

#### func (IOStreams) SpinnerActive

```go
func (s IOStreams) SpinnerActive() bool
```
SpinnerActive returns the active state of the progres spinner

#### func (IOStreams) SpinnerMsg

```go
func (s IOStreams) SpinnerMsg(msg string)
```
SpinnerMsg sets the spinner suffix message

#### func (IOStreams) StartSpinner

```go
func (s IOStreams) StartSpinner()
```
StartSpinner begins the progress spinner

#### func (IOStreams) StopSpinner

```go
func (s IOStreams) StopSpinner()
```
StopSpinner halts the progress spinner

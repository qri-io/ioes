package ioes

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestNewIOStreams(t *testing.T) {
	NewIOStreams(&bytes.Buffer{}, ioutil.Discard, ioutil.Discard)
	NewStdIOStreams()
	NewTestIOStreams()
	NewDiscardIOStreams()
}

func TestSpinner(t *testing.T) {
	s := NewDiscardIOStreams()
	if s.SpinnerActive() != false {
		t.Error("expected spinner to start in a non-active state")
	}
	s.SpinnerMsg("foo")
	s.StartSpinner()
	s.StopSpinner()
}
func TestSpinnerPrint(t *testing.T) {
	s := NewDiscardIOStreams()
	s.StartSpinner()
	s.Print("hallo!")
	s.StopSpinner()
}

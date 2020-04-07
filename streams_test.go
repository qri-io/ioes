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

func TestIsNotATerminal(t *testing.T) {
	str, _, _, _ := NewTestIOStreams()
	if str.IsTerminal() {
		t.Fatal("expected buffer streams to not be a terminal")
	} else if str.IsCygwinTerminal() {
		t.Fatal("expected buffer streams to not be a cygwin terminal")
	}
}

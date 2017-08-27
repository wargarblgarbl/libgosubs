package webvtt

import (
	"testing"
)

func TestLoad(t *testing.T) {
	v, err := ParseWebVtt("../testfiles/sample.vtt")
	if err != nil {
		t.Errorf("Unexpected error Parsing .vtt")
	}
	err2 := WriteWebVtt(v, "../testfiles/sample2.vtt")
	if err2 != nil {
		t.Errorf("Unexpected error writing .vtt")
	}
}

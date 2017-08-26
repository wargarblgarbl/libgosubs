package webvtt

import (
	"testing"
)

func TestLoad(t *testing.T) {
	v := &WebVtt{}
	err := LoadWebVtt(v, "../testfiles/sample.vtt")
	if err != nil {
		t.Errorf("Unexpected error Parsing .vtt")
	}

}

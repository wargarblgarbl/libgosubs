package srt

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSRTIDs(t *testing.T) {
	test, err := ParseSrt("../testfiles/sample.srt")
	if err != nil {
		t.Errorf("Unexpected error")
	}
	for z, i := range test.Subtitle.Content {
		if i.Id != z+1 {
			t.Errorf("IDs do not match")
		}
	}

}

func TestLoadAndWrite(t *testing.T) {
	test, err := ParseSrt("../testfiles/sample.srt")
	if err != nil {
		t.Errorf("Unexpected error")
	}

	err2 := WriteSrt(test, "../testfiles/sample2.srt")
	if err2 != nil {
		t.Errorf("Unexpected error")
	}

	test2, err3 := ParseSrt("../testfiles/sample2.srt")
	if err3 != nil {
		t.Errorf("Unexpected error")
	}

	if cmp.Equal(test, test2) != true {
		t.Errorf("Read structs of input and output do not match", cmp.Diff(test, test2))
	}
}

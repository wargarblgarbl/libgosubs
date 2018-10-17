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
	test, err := ParseSrt("../testfiles/test2.srt")
	if err != nil {
		t.Errorf("Unexpected error parsing .srt")
	}

	err2 := WriteSrt(test, "../testfiles/sample2.srt")
	if err2 != nil {
		t.Errorf("Unexpected error writing .srt")
	}

	test2, err3 := ParseSrt("../testfiles/sample2.srt")
	if err3 != nil {
		t.Errorf("Unexpected error reading sample2.srt")
	}

	if cmp.Equal(test, test2) != true {
		diff := cmp.Diff(test, test2)
		t.Errorf("Read structs of input and output do not match" + diff)
	}
}

func TestBadFileInput(t *testing.T) {
	_, err := ParseSrt("")
	if err == nil {
		t.Errorf("ParseAss did not error correctly on nul string input")
	}
	obj := &SubRip{}
	err2 := WriteSrt(obj, "")
	if err2 == nil {
		t.Errorf("WriteAss did not error correctly on nul string input")
	}
}

func TestCreateSubtile(t *testing.T) {
	text := []string{"what", "subtitle"}
	a := CreateSubtitle(1, "00:02:17,440", "00:02:17,440", text)
	b := &Subtitle{Id: 1, Start: "00:02:17,440", End: "00:02:17,440", Line: text}
	if cmp.Equal(a, b) != true {
		t.Errorf("CreateSubtitle not functioning as expected" + cmp.Diff(a, b))
	}
}

package ass

import (
	"testing"
	//  "strings"
	//  "reflect"
	"github.com/google/go-cmp/cmp"
	//  "github.com/wargarblgarbl/libgosubs/ass"
)

func TestLoadAndWrite(t *testing.T) {
	test, err := ParseAss("../testfiles/sample.ass")
	if err != nil {
		t.Errorf("Unexpected error")
	}
	err2 := WriteAss(test, "../testfiles/sample2.ass")
	if err2 != nil {
		t.Errorf("Unexpected error")
	}

	test2, err3 := ParseAss("../testfiles/sample2.ass")
	if err3 != nil {
		t.Errorf("Unexpected error")
	}

	if cmp.Equal(test, test2) == false {
		t.Errorf("Read structs of input and output do not match", cmp.Diff(test, test2))
	}
}

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
		t.Errorf("Unexpected error Parsing .ass")
	}
	err2 := WriteAss(test, "../testfiles/sample2.ass")
	if err2 != nil {
		t.Errorf("Unexpected error writing .ass")
	}

	test2, err3 := ParseAss("../testfiles/sample2.ass")
	if err3 != nil {
		t.Errorf("Unexpected error parsing sample2.ass")
	}

	if cmp.Equal(test, test2) == false {
		diff := cmp.Diff(test, test2)
		t.Errorf("Read structs of input and output do not match" + diff)
	}
}

func TestBadFileInput(t *testing.T) {
	_, err := ParseAss("")
	if err == nil {
		t.Errorf("ParseAss did not error correctly on nul string input")
	}
	obj := &Ass{}
	err2 := WriteAss(obj, "")
	if err2 == nil {
		t.Errorf("WriteAss did not error correctly on nul string input")
	}
}

func TestHelperFunctions(t *testing.T) {
	floated := floatit("")
	if floated != 0.0 {
		t.Errorf("Floatit did not return 0.0 on emptystring")
	}
	inted := intit("")
	if inted != 0 {
		t.Errorf("Intit did not return 0.0 on emptystring")
	}

}

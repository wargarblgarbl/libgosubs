package mdvd

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadAndWrite(t *testing.T) {
	test, err := ParseMdvd("../testfiles/sample.sub")
	if err != nil {
		fmt.Println(err)
		t.Errorf("Unexpected error parsing .sub")
	}
	fmt.Println(test)
	err2 := WriteMdvd(test, "../testfiles/sample2.sub")
	if err2 != nil {
		t.Errorf("Unexpected error writing .sub")
	}
	test2, err3 := ParseMdvd("../testfiles/sample2.sub")
	if err3 != nil {
		t.Errorf("Unexpected error parsing sample2.sub")
	}
	if cmp.Equal(test, test2) == false {
		diff := cmp.Diff(test, test2)
		t.Errorf("Read structs of input and output do not match" + diff)
	}

}

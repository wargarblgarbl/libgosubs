package ass
import (
  "testing"
//  "strings"
//  "reflect"
  "github.com/google/go-cmp/cmp"
//  "github.com/wargarblgarbl/libgosubs/ass"
)

func TestLoadAndWrite(t *testing.T) {
  test := ParseAss("../testfiles/sample.ass")
  WriteAss(test, "../testfiles/sample2.ass")
  test2 := ParseAss("../testfiles/sample2.ass")
  if cmp.Equal(test, test2) == false {
    t.Errorf("Read structs of input and output do not match", cmp.Diff(test, test2))
  }
}

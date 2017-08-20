package ttml
import (
  "testing"
  "github.com/google/go-cmp/cmp"
)

func TestLoadAndWrite(t *testing.T) {
  test := ParseTtml("../testfiles/sample.ttml")
  a := WTt(*test)

  WriteTtml(&a, "../testfiles/sample2.ttml")
  test2 := ParseTtml("../testfiles/sample2.ttml")
  if cmp.Equal(test, test2) == false {
    t.Errorf("Read structs of input and output do not match", cmp.Diff(test, test2))
  }
}

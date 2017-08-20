package srt
import (
  "testing"
  "github.com/google/go-cmp/cmp"

)

func TestSRTIDs(t *testing.T) {
  test := ParseSrt("../testfiles/sample.srt")
  for z, i := range test.Subtitle.Content {
    if i.Id != z+1 {
      t.Errorf("IDs do not match")
    }
  }

}

func TestLoadAndWrite(t *testing.T) {
  test := ParseSrt("../testfiles/sample.srt")
  WriteSrt(test, "../testfiles/sample2.srt")
  test2 := ParseSrt("../testfiles/sample2.srt")
  if cmp.Equal(test, test2) != true {
    t.Errorf("Read structs of input and output do not match", cmp.Diff(test, test2))
  }
}

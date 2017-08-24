package srt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//WriteSrt takes a SubRip object and the path to which to write the file as a string
func WriteSrt(v *SubRip, outpath string) error {
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	var outout []string
	for _, z := range v.Subtitle.Content {
		lines := strings.Join(z.Line, "\n")
		a := strconv.Itoa(z.Id) + "\n" + z.Start + " --> " + z.End + "\n" + lines
		outout = append(outout, a)
	}
	fmt.Fprint(f, strings.Join(outout, "\n\n"))
	return nil
}

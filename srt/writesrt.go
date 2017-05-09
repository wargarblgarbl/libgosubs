package srt

import (
	"fmt"
	"os"
	"strings"
)

func WriteSrt(v *SubRip, outpath string) {
	f, err := os.Create(outpath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	var outout []string
	for _, z := range v.subtitle.content {
		lines := strings.Join(z.line, "\n")
		a := z.id + "\n" + z.start + "-->" + z.end + "\n" + lines + "\n"
		outout = append(outout, a)
	}
	fmt.Fprintf(f, "%", string(outout))
}


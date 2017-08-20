package srt

import (
	"fmt"
	"os"
	"strconv"
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
		lines := strings.Join(z.Line, "\n")
		a := strconv.Itoa(z.Id) + "\n" + z.Start + "-->" + z.End + "\n" + lines + "\n"
		outout = append(outout, a)
	}
	fmt.Fprintf(f, "%", strings.Join(outout, " "))
}

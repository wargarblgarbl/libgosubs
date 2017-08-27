package webvtt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func wparsepos(v *Position) (posstring string) {
	var posinfo []string
	if v.Vertical != "" {
		posinfo = append(posinfo, "vertical:"+v.Vertical)
	}
	if v.Line >= -1 {
		i := strconv.Itoa(v.Line)
		if v.Linepercent {
			i = i + "%"
		}
		posinfo = append(posinfo, "line:"+i)
	}

	if v.Position >= -1 {
		i := strconv.Itoa(v.Position)
		i = "position:" + i + "%"
		if v.Posstring != "" {
			i = i + "," + v.Posstring
		}
		posinfo = append(posinfo, i)
	}
	if v.Align != "" {
		posinfo = append(posinfo, "align:"+v.Align)
	}
	if v.Size >= -1 {
		i := strconv.Itoa(v.Size)
		posinfo = append(posinfo, "size"+i+"%")
	}

	posstring = strings.Join(posinfo, " ")
	return
}

//WriteWebVtt takes a WebVtt object and the path to which to write the file as a string
func WriteWebVtt(v *WebVtt, outpath string) error {
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	var outout []string
	var outstyles []string
	outout = append(outout, v.Header)
	if v.Styles != nil {
		for z, s := range v.Styles {
			outstyles = append(outstyles, "STYLE")
			outstyles = append(outstyles, s.Header+"{")
			if s.Value != nil {
				for z, p := range s.Value {
					line := fmt.Sprint(z, ":", p, ";")
					outstyles = append(outstyles, line)
				}
			}
			if z == len(v.Styles)-1 {
				outstyles = append(outstyles, "}")

			} else {
				outstyles = append(outstyles, "}\n")

			}
		}
		outout = append(outout, strings.Join(outstyles, "\n"))

	}
	for _, z := range v.Subtitle.Content {
		lines := strings.Join(z.Line, "\n")
		if z.Note {
			outout = append(outout, lines)
		}
		if z.Haspos {
			posstring := wparsepos(&z.Position)
			z.End = z.End + " " + posstring
		}
		if z.Note == false && z.Cue != "" {
			a := z.Cue + "\n" + z.Start + " --> " + z.End + "\n" + lines
			outout = append(outout, a)
		}
		if z.Note == false && z.Cue == "" {
			a := z.Start + " --> " + z.End + "\n" + lines
			outout = append(outout, a)
		}
	}
	fmt.Fprint(f, strings.Join(outout, "\n\n"))
	return nil
}

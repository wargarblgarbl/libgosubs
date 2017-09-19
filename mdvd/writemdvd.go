package mdvd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//helper function

func stringintsf(in int64) (out string) {
	out = strconv.FormatInt(in, 10)
	return
}

func handletags(in []Tag) (out string) {
	var outout []string
	for _, i := range in {
		outout = append(outout, "{"+i.Type+":"+i.Value+"}")
	}
	out = strings.Join(outout, "")
	return
}

//WriteMdvd writes a .sub file from an Mdvd object
func WriteMdvd(v *Mdvd, outpath string) error {

	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	var outout []string
	for _, z := range v.Body {
		var anevent []string
		if z.IsDefault {
			anevent = append(anevent, "{DEFAULT}")
			anevent = append(anevent, handletags(z.Tags))
			anevent = append(anevent, z.Text)

		} else {
			anevent = append(anevent, "{"+stringintsf(z.Start)+"}")
			anevent = append(anevent, "{"+stringintsf(z.End)+"}")
			anevent = append(anevent, handletags(z.Tags))
			anevent = append(anevent, z.Text)
		}
		outout = append(outout, strings.Join(anevent, ""))

	}
	fmt.Fprintf(f, strings.Join(outout, "\n"))
	return nil
}

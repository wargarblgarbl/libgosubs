package ttml

import (
	"os"
	"fmt"
	"encoding/xml"
)

func WriteTtml (v *Tt, outpath string) {
	f, err := os.Create(outpath)
	if err != nil {
		fmt.Println(err)

	}
	//marshall XML to file, see the
//	xml.Marshal(v, )
}
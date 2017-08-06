package ttml

import (
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
)




func WriteTtml (v *WTt, outpath string) {
	f, err := os.Create(outpath)
	if err != nil {
		fmt.Println(err)

	}
	out, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	bytehead := []byte(xml.Header)
	out = append(bytehead, out ...)
	f, arr := os.Create(outpath)
	if arr != nil {
		panic(arr)
	}
	defer f.Close()
	trr := ioutil.WriteFile(outpath, []byte(out), 0666)
	if trr != nil {
		panic(trr)
	}
}

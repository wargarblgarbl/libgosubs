package ttml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

//WriteTtml writes the TTML file
func WriteTtml(v *WTt, outpath string) error {
	f, err := os.Create(outpath)
	if err != nil {
		return (err)

	}
	out, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	bytehead := []byte(xml.Header)
	out = append(bytehead, out...)
	f, arr := os.Create(outpath)
	if arr != nil {
		return arr
	}
	defer f.Close()
	trr := ioutil.WriteFile(outpath, []byte(out), 0666)
	if trr != nil {
		return trr
	}
	//return nil error
	return nil
}

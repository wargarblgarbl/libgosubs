package ttml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//LoadTtml loads the TTML file from a given filepath
func LoadTtml(v *Tt, filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Cannot read file", filepath)
		os.Exit(1)
	}
	bytef, berr := ioutil.ReadAll(f)
	if berr != nil {
		fmt.Println("error decoding")
	}
	xml.Unmarshal(bytef, &v)
}

//Generic loader for TTML files
func ParseTtml(filename string) *Tt {
	v := &Tt{}
	LoadTtml(v, filename)
	return v
}

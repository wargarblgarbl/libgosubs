package ttml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"errors"
)

//LoadTtml loads the TTML file from a given filepath
func LoadTtml(v *Tt, filepath string)(error) {
	f, err := os.Open(filepath)
	if err != nil {
		return errors.New("Cannot read file :"+filepath)
}
	bytef, berr := ioutil.ReadAll(f)
	if berr != nil {
		return errors.New("error decoding")
	}
	xml.Unmarshal(bytef, &v)
	return nil
}

//ParseTtml is a generic loader for TTML files
func ParseTtml(filename string) *Tt {
	v := &Tt{}
	err := LoadTtml(v, filename)
	if err != nil {
		panic(err)
	}
	return v
}

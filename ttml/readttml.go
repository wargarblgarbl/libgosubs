package ttml
import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)

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

func ParseTtml(filename string) *Tt{
	v := &Tt{}
	LoadTtml(v, filename)
	return v
}



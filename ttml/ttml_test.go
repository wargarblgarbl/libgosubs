package ttml

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

//Convert region struct
func convertregion(i *Region) Wregion {
	a := Wregion(*i)
	return a
}

func convertstyle(i *Style) Wstyle {
	a := Wstyle(*i)
	return a

}

func convertsub(i *Subtitle) Wsubtitle {
	a := Wsubtitle(*i)
	return a
}

func convertstruct(i *Tt) *WTt {
	a := &WTt{}
	a.Xmlns = i.Xmlns
	a.XmlnsTtp = i.XmlnsTtp
	a.XmlnsTts = i.XmlnsTts
	a.XmlnsTtm = i.XmlnsTtm
	a.XmlnsXML = i.XmlnsXML
	a.TtpTimeBase = i.TtpTimeBase
	a.TtpFrameRate = i.TtpFrameRate
	a.XMLLang = i.XMLLang
	a.Head.Metadata.TtmTitle = i.Head.Metadata.TtmTitle
	style := []Wstyle{}
	for _, e := range i.Head.Styling.Style {
		style = append(style, convertstyle(&e))
	}
	a.Head.Styling.Style = style
	region := []Wregion{}
	for _, e := range i.Head.Layout.Region {
		region = append(region, convertregion(&e))
	}
	a.Head.Layout.Region = region
	a.Body.Region = i.Body.Region
	a.Body.Style = i.Body.Style
	wsubtitle := []Wsubtitle{}
	for _, e := range i.Body.Div.P {
		wsubtitle = append(wsubtitle, convertsub(&e))
	}
	a.Body.Div.P = wsubtitle

	return a

}

//Test our testing functions
func TestAlllocal(t *testing.T) {
	a := &Tt{}
	b := &WTt{}
	c := &Region{}
	d := &Style{}
	e := &Subtitle{}
	f := &Wregion{}
	g := &Wsubtitle{}
	h := &Wstyle{}

	if cmp.Equal(convertstruct(a), b) {
		//t.Errorf("convertstruct test function failed")
	}
	if convertregion(c) != *f {
		t.Errorf("convertregion test function failed")
	}
	if convertsub(e) != *g {
		t.Errorf("convertsub test function failed")
	}
	if convertstyle(d) != *h {
		t.Errorf("convertstyle test function failed")
	}
	// Output: matches
}

//Test loading and writing
func TestLoadAndWrite(t *testing.T) {
	test, err2 := ParseTtml("../testfiles/sample.ttml")
	if err2 != nil {
		t.Errorf("WriteTtml returned an unexpected error")
	}

	a := convertstruct(test)
	WriteTtml(a, "../testfiles/sample2.ttml")
	test2, err := ParseTtml("../testfiles/sample2.ttml")
	if err != nil {
		t.Errorf("WriteTtml returned an unexpected error")
	}
	if cmp.Equal(test, test2) == false {
		diff := cmp.Diff(test, test2)
		t.Errorf("Read structs of input and output do not match" + diff)
	}
}

//Test bad input
func TestFilepath(t *testing.T) {
	v := &Tt{}
	err := LoadTtml(v, "bunkpath")
	if err == nil {
		t.Errorf("LoadTTml does not return error")
	}
}

//Test writing errors
func TestWrite(t *testing.T) {
	v := &WTt{}
	err := WriteTtml(v, "")
	// t.Errorf(err)
	if err == nil {
		t.Errorf("WriteTtml does not return error")
	}
}

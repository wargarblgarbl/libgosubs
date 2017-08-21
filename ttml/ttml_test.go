package ttml

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

//Convert region struct
func convertregion(i *Region)(Wregion) {
	a := Wregion(*i)
	return a
}

func convertstyle(i *Style)(Wstyle) {
	a := Wstyle(*i)
	return a

}

func convertsub(i *Subtitle)(Wsubtitle){
	a := Wsubtitle(*i)
	return a
}

func convertstruct(i *Tt)(*WTt) {
	a := &WTt{}
	a.Xmlns = i.Xmlns
	a.XmlnsTtp =  i.XmlnsTtp   
	a.XmlnsTts =    i.XmlnsTts
	a.XmlnsTtm =    i.XmlnsTtm
	a.XmlnsXML   =  i.XmlnsXML
	a.TtpTimeBase  = i.TtpTimeBase
	a.TtpFrameRate = i.TtpFrameRate
	a.XMLLang      = i.XMLLang
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

func TestLoadAndWrite(t *testing.T) {
	test := ParseTtml("../testfiles/sample.ttml")
	a := convertstruct(test)
	WriteTtml(a, "../testfiles/sample2.ttml")
	test2 := ParseTtml("../testfiles/sample2.ttml")
	if cmp.Equal(test, test2) == false {
		t.Errorf("Read structs of input and output do not match", cmp.Diff(test, test2))
	}
}

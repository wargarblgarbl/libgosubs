package ttml

import (
	"encoding/xml"
)

//WTt is the Write TTML struct.
//Primary reason for using a separate struct is Go's strange handling of complex XML parameters
type WTt struct {
	XMLName      xml.Name `xml:"tt"`
	Xmlns        string   `xml:"xmlns,attr"`
	XmlnsTtp     string   `xml:"xmlns:ttp,attr"`
	XmlnsTts     string   `xml:"xmlns:tts,attr"`
	XmlnsTtm     string   `xml:"xmlns:ttm,attr"`
	XmlnsXML     string   `xml:"xmlns:xml,attr"`
	TtpTimeBase  string   `xml:"ttp:timeBase,attr"`
	TtpFrameRate string   `xml:"ttp:frameRate,attr"`
	XMLLang      string   `xml:"xml:lang,attr"`
	Head         struct {
		Metadata struct {
			TtmTitle string `xml:"ttm:title"`
		} `xml:"metadata"`
		Styling struct {
			Style []Wstyle `xml:"style"`
		} `xml:"styling"`
		Layout struct {
			Region []Wregion `xml:"region"`
		} `xml:"layout"`
	} `xml:"head"`
	Body struct {
		Region string `xml:"region,attr"`
		Style  string `xml:"style,attr"`
		Div    struct {
			P []Wsubtitle `xml:"p"`
		} `xml:"div"`
	} `xml:"body"`
}

//Wregion - Region write struct
type Wregion struct {
	XMLID           string `xml:"xml:id,attr"`
	TtsDisplayAlign string `xml:"tts:displayAlign,attr"`
	TtsExtent       string `xml:"tts:extent,attr"`
	TtsOrigin       string `xml:"tts:origin,attr"`
}

//Wstyle - Style write struct
type Wstyle struct {
	XMLID         string `xml:"xml:id,attr"`
	TtsTextAlign  string `xml:"tts:textAlign,attr"`
	TtsFontFamily string `xml:"tts:fontFamily,attr"`
	TtsFontSize   string `xml:"tts:fontSize,attr"`
}

//Wsubtitle - Subtitle write struct
type Wsubtitle struct {
	Id     string `xml:"xml:id,attr"`
	Begin  string `xml:"begin,attr"`
	End    string `xml:"end,attr"`
	Style  string `xml:"style,attr,omitempty"`
	Region string `xml:"region,attr,omitempty"`
	Text   string `xml:",innerxml"`
}
